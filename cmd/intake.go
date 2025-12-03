package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/scooter-indie/gh-pmu/internal/api"
	"github.com/scooter-indie/gh-pmu/internal/config"
	"github.com/spf13/cobra"
)

type intakeOptions struct {
	apply  bool
	dryRun bool
	json   bool
}

func newIntakeCommand() *cobra.Command {
	opts := &intakeOptions{}

	cmd := &cobra.Command{
		Use:   "intake",
		Short: "Find issues not yet added to the project",
		Long: `Find open issues in configured repositories that are not yet tracked in the project.

This helps ensure all work is captured on your project board.
Use --apply to automatically add discovered issues to the project.`,
		Aliases: []string{"in"},
		Example: `  # List untracked issues
  gh pmu intake

  # Preview what would be added
  gh pmu intake --dry-run

  # Add untracked issues to project
  gh pmu intake --apply

  # Output as JSON
  gh pmu intake --json`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runIntake(cmd, opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.apply, "apply", "a", false, "Add untracked issues to the project with default fields")
	cmd.Flags().BoolVar(&opts.dryRun, "dry-run", false, "Show what would be added without making changes")
	cmd.Flags().BoolVar(&opts.json, "json", false, "Output in JSON format")

	return cmd
}

func runIntake(cmd *cobra.Command, opts *intakeOptions) error {
	// Load configuration
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	cfg, err := config.LoadFromDirectory(cwd)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w\nRun 'gh pmu init' to create a configuration file", err)
	}

	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	if len(cfg.Repositories) == 0 {
		return fmt.Errorf("no repositories configured in .gh-pmu.yml")
	}

	// Create API client
	client := api.NewClient()

	// Get project
	project, err := client.GetProject(cfg.Project.Owner, cfg.Project.Number)
	if err != nil {
		return fmt.Errorf("failed to get project: %w", err)
	}

	// Get all issues currently in the project
	projectItems, err := client.GetProjectItems(project.ID, nil)
	if err != nil {
		return fmt.Errorf("failed to get project items: %w", err)
	}

	// Build set of issue IDs already in project
	trackedIssues := make(map[string]bool)
	for _, item := range projectItems {
		if item.Issue != nil {
			trackedIssues[item.Issue.ID] = true
		}
	}

	// Find untracked issues from each repository
	var untrackedIssues []api.Issue
	for _, repoFullName := range cfg.Repositories {
		parts := strings.SplitN(repoFullName, "/", 2)
		if len(parts) != 2 {
			cmd.PrintErrf("Warning: invalid repository format %q, expected owner/repo\n", repoFullName)
			continue
		}
		owner, repo := parts[0], parts[1]

		// Get open issues from repository
		issues, err := client.GetRepositoryIssues(owner, repo, "open")
		if err != nil {
			cmd.PrintErrf("Warning: failed to get issues from %s: %v\n", repoFullName, err)
			continue
		}

		// Filter to untracked issues
		for _, issue := range issues {
			if !trackedIssues[issue.ID] {
				issue.Repository = api.Repository{Owner: owner, Name: repo}
				untrackedIssues = append(untrackedIssues, issue)
			}
		}
	}

	// Handle output
	if len(untrackedIssues) == 0 {
		if !opts.json {
			cmd.Println("All issues are already tracked in the project")
		} else {
			encoder := json.NewEncoder(os.Stdout)
			encoder.SetIndent("", "  ")
			encoder.Encode(map[string]interface{}{"issues": []interface{}{}, "count": 0})
		}
		return nil
	}

	// Dry run - just show what would be added
	if opts.dryRun {
		if opts.json {
			return outputIntakeJSON(cmd, untrackedIssues, "dry-run")
		}
		cmd.Printf("Would add %d issue(s) to project:\n\n", len(untrackedIssues))
		return outputIntakeTable(cmd, untrackedIssues)
	}

	// Apply - add issues to project
	if opts.apply {
		var added []api.Issue
		var failed []api.Issue

		for _, issue := range untrackedIssues {
			itemID, err := client.AddIssueToProject(project.ID, issue.ID)
			if err != nil {
				cmd.PrintErrf("Failed to add #%d: %v\n", issue.Number, err)
				failed = append(failed, issue)
				continue
			}

			// Apply default fields from config
			if cfg.Defaults.Status != "" {
				statusValue := cfg.ResolveFieldValue("status", cfg.Defaults.Status)
				if err := client.SetProjectItemField(project.ID, itemID, "Status", statusValue); err != nil {
					cmd.PrintErrf("Warning: failed to set status on #%d: %v\n", issue.Number, err)
				}
			}
			if cfg.Defaults.Priority != "" {
				priorityValue := cfg.ResolveFieldValue("priority", cfg.Defaults.Priority)
				if err := client.SetProjectItemField(project.ID, itemID, "Priority", priorityValue); err != nil {
					cmd.PrintErrf("Warning: failed to set priority on #%d: %v\n", issue.Number, err)
				}
			}

			added = append(added, issue)
		}

		if opts.json {
			return outputIntakeJSON(cmd, added, "applied")
		}

		cmd.Printf("Added %d issue(s) to project", len(added))
		if len(failed) > 0 {
			cmd.Printf(" (%d failed)", len(failed))
		}
		cmd.Println()
		return nil
	}

	// Default - just list untracked issues
	if opts.json {
		return outputIntakeJSON(cmd, untrackedIssues, "untracked")
	}

	cmd.Printf("Found %d untracked issue(s):\n\n", len(untrackedIssues))
	if err := outputIntakeTable(cmd, untrackedIssues); err != nil {
		return err
	}
	cmd.Println("\nUse --apply to add these issues to the project")
	return nil
}

func outputIntakeTable(cmd *cobra.Command, issues []api.Issue) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "NUMBER\tTITLE\tREPOSITORY\tSTATE")

	for _, issue := range issues {
		title := issue.Title
		if len(title) > 50 {
			title = title[:47] + "..."
		}
		repoName := fmt.Sprintf("%s/%s", issue.Repository.Owner, issue.Repository.Name)
		fmt.Fprintf(w, "#%d\t%s\t%s\t%s\n", issue.Number, title, repoName, issue.State)
	}

	return w.Flush()
}

type intakeJSONOutput struct {
	Status string            `json:"status"`
	Count  int               `json:"count"`
	Issues []intakeJSONIssue `json:"issues"`
}

type intakeJSONIssue struct {
	Number     int    `json:"number"`
	Title      string `json:"title"`
	State      string `json:"state"`
	URL        string `json:"url"`
	Repository string `json:"repository"`
}

func outputIntakeJSON(cmd *cobra.Command, issues []api.Issue, status string) error {
	output := intakeJSONOutput{
		Status: status,
		Count:  len(issues),
		Issues: make([]intakeJSONIssue, 0, len(issues)),
	}

	for _, issue := range issues {
		output.Issues = append(output.Issues, intakeJSONIssue{
			Number:     issue.Number,
			Title:      issue.Title,
			State:      issue.State,
			URL:        issue.URL,
			Repository: fmt.Sprintf("%s/%s", issue.Repository.Owner, issue.Repository.Name),
		})
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(output)
}
