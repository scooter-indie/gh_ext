package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/scooter-indie/gh-pmu/internal/api"
)

func TestSubCommand_Exists(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"sub", "--help"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("sub command should exist: %v", err)
	}

	output := buf.String()
	if !bytes.Contains([]byte(output), []byte("sub")) {
		t.Error("Expected help output to mention 'sub'")
	}
}

func TestSubAddCommand_Exists(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"sub", "add", "--help"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("sub add command should exist: %v", err)
	}

	output := buf.String()
	if !bytes.Contains([]byte(output), []byte("add")) {
		t.Error("Expected help output to mention 'add'")
	}
}

func TestSubAddCommand_RequiresTwoArgs(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"sub", "add", "123"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err == nil {
		t.Error("Expected error when only one argument provided")
	}
}

func TestSubAddCommand_RequiresParentAndChild(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"sub", "add"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err == nil {
		t.Error("Expected error when no arguments provided")
	}
}

func TestSubCreateCommand_Exists(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"sub", "create", "--help"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("sub create command should exist: %v", err)
	}

	output := buf.String()
	if !bytes.Contains([]byte(output), []byte("create")) {
		t.Error("Expected help output to mention 'create'")
	}
}

func TestSubCreateCommand_HasParentFlag(t *testing.T) {
	cmd := NewRootCommand()
	subCmd, _, err := cmd.Find([]string{"sub", "create"})
	if err != nil {
		t.Fatalf("sub create command not found: %v", err)
	}

	flag := subCmd.Flags().Lookup("parent")
	if flag == nil {
		t.Fatal("Expected --parent flag to exist")
	}
}

func TestSubCreateCommand_HasTitleFlag(t *testing.T) {
	cmd := NewRootCommand()
	subCmd, _, err := cmd.Find([]string{"sub", "create"})
	if err != nil {
		t.Fatalf("sub create command not found: %v", err)
	}

	flag := subCmd.Flags().Lookup("title")
	if flag == nil {
		t.Fatal("Expected --title flag to exist")
	}
}

func TestSubCreateCommand_RequiresParentFlag(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"sub", "create", "--title", "Test"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err == nil {
		t.Error("Expected error when --parent not provided")
	}
}

func TestSubCreateCommand_RequiresTitleFlag(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"sub", "create", "--parent", "123"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err == nil {
		t.Error("Expected error when --title not provided")
	}
}

func TestSubListCommand_Exists(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"sub", "list", "--help"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("sub list command should exist: %v", err)
	}

	output := buf.String()
	if !bytes.Contains([]byte(output), []byte("list")) {
		t.Error("Expected help output to mention 'list'")
	}
}

func TestSubListCommand_RequiresParentArg(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"sub", "list"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err == nil {
		t.Error("Expected error when parent issue not provided")
	}
}

func TestSubListCommand_HasJSONFlag(t *testing.T) {
	cmd := NewRootCommand()
	subCmd, _, err := cmd.Find([]string{"sub", "list"})
	if err != nil {
		t.Fatalf("sub list command not found: %v", err)
	}

	flag := subCmd.Flags().Lookup("json")
	if flag == nil {
		t.Fatal("Expected --json flag to exist")
	}
}

func TestSubRemoveCommand_Exists(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"sub", "remove", "--help"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("sub remove command should exist: %v", err)
	}

	output := buf.String()
	if !bytes.Contains([]byte(output), []byte("remove")) {
		t.Error("Expected help output to mention 'remove'")
	}
}

func TestSubRemoveCommand_RequiresTwoArgs(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"sub", "remove", "123"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err == nil {
		t.Error("Expected error when only one argument provided")
	}
}

func TestSubRemoveCommand_RequiresParentAndChild(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"sub", "remove"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err == nil {
		t.Error("Expected error when no arguments provided")
	}
}

// Cross-repository sub-issue tests

func TestSubCreateCommand_HasRepoFlag(t *testing.T) {
	cmd := NewRootCommand()
	subCmd, _, err := cmd.Find([]string{"sub", "create"})
	if err != nil {
		t.Fatalf("sub create command not found: %v", err)
	}

	flag := subCmd.Flags().Lookup("repo")
	if flag == nil {
		t.Fatal("Expected --repo flag to exist")
	}

	// Verify short flag
	if flag.Shorthand != "R" {
		t.Errorf("Expected --repo shorthand to be 'R', got '%s'", flag.Shorthand)
	}
}

func TestSubCreateCommand_RepoFlagHelpText(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"sub", "create", "--help"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("sub create help failed: %v", err)
	}

	output := buf.String()
	// Verify cross-repo example is shown
	if !bytes.Contains([]byte(output), []byte("--repo owner/repo2")) {
		t.Error("Expected help to show cross-repo example with --repo flag")
	}
}

func TestSubAddCommand_HelpShowsCrossRepoExample(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"sub", "add", "--help"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("sub add help failed: %v", err)
	}

	output := buf.String()
	// Verify cross-repo format is documented
	if !bytes.Contains([]byte(output), []byte("owner/repo#")) {
		t.Error("Expected help to document owner/repo#number format for cross-repo")
	}
}

// ============================================================================
// subCreateOptions Tests
// ============================================================================

func TestSubCreateOptions_Defaults(t *testing.T) {
	// Get the command to verify default values
	cmd := NewRootCommand()
	subCmd, _, err := cmd.Find([]string{"sub", "create"})
	if err != nil {
		t.Fatalf("sub create command not found: %v", err)
	}

	// Check inherit-labels default (should be true)
	inheritLabels := subCmd.Flags().Lookup("inherit-labels")
	if inheritLabels == nil {
		t.Fatal("Expected --inherit-labels flag to exist")
	}
	if inheritLabels.DefValue != "true" {
		t.Errorf("Expected --inherit-labels default to be 'true', got '%s'", inheritLabels.DefValue)
	}

	// Check inherit-assignees default (should be false)
	inheritAssign := subCmd.Flags().Lookup("inherit-assignees")
	if inheritAssign == nil {
		t.Fatal("Expected --inherit-assignees flag to exist")
	}
	if inheritAssign.DefValue != "false" {
		t.Errorf("Expected --inherit-assignees default to be 'false', got '%s'", inheritAssign.DefValue)
	}

	// Check inherit-milestone default (should be true)
	inheritMilestone := subCmd.Flags().Lookup("inherit-milestone")
	if inheritMilestone == nil {
		t.Fatal("Expected --inherit-milestone flag to exist")
	}
	if inheritMilestone.DefValue != "true" {
		t.Errorf("Expected --inherit-milestone default to be 'true', got '%s'", inheritMilestone.DefValue)
	}
}

func TestSubCreateCommand_HasBodyFlag(t *testing.T) {
	cmd := NewRootCommand()
	subCmd, _, err := cmd.Find([]string{"sub", "create"})
	if err != nil {
		t.Fatalf("sub create command not found: %v", err)
	}

	flag := subCmd.Flags().Lookup("body")
	if flag == nil {
		t.Fatal("Expected --body flag to exist")
	}

	if flag.Shorthand != "b" {
		t.Errorf("Expected --body shorthand to be 'b', got '%s'", flag.Shorthand)
	}
}

// ============================================================================
// outputSubListJSON Tests
// ============================================================================

func TestOutputSubListJSON_EmptyList(t *testing.T) {
	parent := &api.Issue{
		Number: 10,
		Title:  "Parent Issue",
	}
	subIssues := []api.SubIssue{}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := outputSubListJSON(subIssues, parent)

	w.Close()
	os.Stdout = oldStdout

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	output, _ := io.ReadAll(r)

	var result SubListJSONOutput
	if err := json.Unmarshal(output, &result); err != nil {
		t.Fatalf("Failed to parse JSON output: %v", err)
	}

	if result.Parent.Number != 10 {
		t.Errorf("Expected parent number 10, got %d", result.Parent.Number)
	}
	if result.Parent.Title != "Parent Issue" {
		t.Errorf("Expected parent title 'Parent Issue', got '%s'", result.Parent.Title)
	}
	if result.Summary.Total != 0 {
		t.Errorf("Expected total 0, got %d", result.Summary.Total)
	}
	if result.Summary.Open != 0 {
		t.Errorf("Expected open 0, got %d", result.Summary.Open)
	}
	if result.Summary.Closed != 0 {
		t.Errorf("Expected closed 0, got %d", result.Summary.Closed)
	}
	if len(result.SubIssues) != 0 {
		t.Errorf("Expected 0 sub-issues, got %d", len(result.SubIssues))
	}
}

func TestOutputSubListJSON_SummaryCounts(t *testing.T) {
	parent := &api.Issue{
		Number: 10,
		Title:  "Parent Issue",
	}
	subIssues := []api.SubIssue{
		{Number: 1, Title: "Open 1", State: "OPEN", Repository: api.Repository{Owner: "owner", Name: "repo"}},
		{Number: 2, Title: "Open 2", State: "OPEN", Repository: api.Repository{Owner: "owner", Name: "repo"}},
		{Number: 3, Title: "Closed 1", State: "CLOSED", Repository: api.Repository{Owner: "owner", Name: "repo"}},
	}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := outputSubListJSON(subIssues, parent)

	w.Close()
	os.Stdout = oldStdout

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	output, _ := io.ReadAll(r)

	var result SubListJSONOutput
	if err := json.Unmarshal(output, &result); err != nil {
		t.Fatalf("Failed to parse JSON output: %v", err)
	}

	// Verify summary counts match acceptance criteria
	if result.Summary.Total != 3 {
		t.Errorf("Expected total=3, got %d", result.Summary.Total)
	}
	if result.Summary.Open != 2 {
		t.Errorf("Expected open=2, got %d", result.Summary.Open)
	}
	if result.Summary.Closed != 1 {
		t.Errorf("Expected closed=1, got %d", result.Summary.Closed)
	}
}

func TestOutputSubListJSON_RepositoryField(t *testing.T) {
	parent := &api.Issue{
		Number: 10,
		Title:  "Parent Issue",
	}
	subIssues := []api.SubIssue{
		{
			Number:     1,
			Title:      "Sub in other repo",
			State:      "OPEN",
			URL:        "https://github.com/other/repo/issues/1",
			Repository: api.Repository{Owner: "other", Name: "repo"},
		},
		{
			Number:     2,
			Title:      "Sub with empty repo",
			State:      "OPEN",
			URL:        "https://github.com/owner/repo/issues/2",
			Repository: api.Repository{Owner: "", Name: ""},
		},
	}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := outputSubListJSON(subIssues, parent)

	w.Close()
	os.Stdout = oldStdout

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	output, _ := io.ReadAll(r)

	var result SubListJSONOutput
	if err := json.Unmarshal(output, &result); err != nil {
		t.Fatalf("Failed to parse JSON output: %v", err)
	}

	if len(result.SubIssues) != 2 {
		t.Fatalf("Expected 2 sub-issues, got %d", len(result.SubIssues))
	}

	// First sub-issue should have full repo
	if result.SubIssues[0].Repository != "other/repo" {
		t.Errorf("Expected repository 'other/repo', got '%s'", result.SubIssues[0].Repository)
	}

	// Second sub-issue with empty repo should have empty string
	if result.SubIssues[1].Repository != "" {
		t.Errorf("Expected empty repository, got '%s'", result.SubIssues[1].Repository)
	}
}

func TestOutputSubListJSON_AllFields(t *testing.T) {
	parent := &api.Issue{
		Number: 42,
		Title:  "Epic Issue",
	}
	subIssues := []api.SubIssue{
		{
			Number:     100,
			Title:      "Task One",
			State:      "OPEN",
			URL:        "https://github.com/owner/repo/issues/100",
			Repository: api.Repository{Owner: "owner", Name: "repo"},
		},
	}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := outputSubListJSON(subIssues, parent)

	w.Close()
	os.Stdout = oldStdout

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	output, _ := io.ReadAll(r)

	var result SubListJSONOutput
	if err := json.Unmarshal(output, &result); err != nil {
		t.Fatalf("Failed to parse JSON output: %v", err)
	}

	// Verify all fields are present
	sub := result.SubIssues[0]
	if sub.Number != 100 {
		t.Errorf("Expected number 100, got %d", sub.Number)
	}
	if sub.Title != "Task One" {
		t.Errorf("Expected title 'Task One', got '%s'", sub.Title)
	}
	if sub.State != "OPEN" {
		t.Errorf("Expected state 'OPEN', got '%s'", sub.State)
	}
	if sub.URL != "https://github.com/owner/repo/issues/100" {
		t.Errorf("Expected URL 'https://github.com/owner/repo/issues/100', got '%s'", sub.URL)
	}
	if sub.Repository != "owner/repo" {
		t.Errorf("Expected repository 'owner/repo', got '%s'", sub.Repository)
	}
}

// ============================================================================
// outputSubListTable Tests
// ============================================================================

func TestOutputSubListTable_EmptyList(t *testing.T) {
	parent := &api.Issue{
		Number: 10,
		Title:  "Parent Issue",
		Repository: api.Repository{
			Owner: "owner",
			Name:  "repo",
		},
	}
	subIssues := []api.SubIssue{}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := outputSubListTable(subIssues, parent)

	w.Close()
	os.Stdout = oldStdout

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	output, _ := io.ReadAll(r)
	outputStr := string(output)

	if !strings.Contains(outputStr, "No sub-issues found") {
		t.Error("Expected 'No sub-issues found' message for empty list")
	}
}

func TestOutputSubListTable_SingleRepo(t *testing.T) {
	parent := &api.Issue{
		Number: 10,
		Title:  "Parent Issue",
		Repository: api.Repository{
			Owner: "owner",
			Name:  "repo",
		},
	}
	subIssues := []api.SubIssue{
		{Number: 1, Title: "Task 1", State: "OPEN", Repository: api.Repository{Owner: "owner", Name: "repo"}},
		{Number: 2, Title: "Task 2", State: "CLOSED", Repository: api.Repository{Owner: "owner", Name: "repo"}},
	}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := outputSubListTable(subIssues, parent)

	w.Close()
	os.Stdout = oldStdout

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	output, _ := io.ReadAll(r)
	outputStr := string(output)

	// Should show parent info
	if !strings.Contains(outputStr, "#10") {
		t.Error("Expected parent issue number in output")
	}
	if !strings.Contains(outputStr, "Parent Issue") {
		t.Error("Expected parent title in output")
	}

	// Should show sub-issues without repo prefix (same repo)
	if !strings.Contains(outputStr, "#1") {
		t.Error("Expected sub-issue #1 in output")
	}
	if !strings.Contains(outputStr, "#2") {
		t.Error("Expected sub-issue #2 in output")
	}

	// Should show progress
	if !strings.Contains(outputStr, "1/2 complete") {
		t.Error("Expected progress '1/2 complete' in output")
	}

	// Check state indicators
	if !strings.Contains(outputStr, "[ ]") {
		t.Error("Expected open indicator '[ ]' in output")
	}
	if !strings.Contains(outputStr, "[x]") {
		t.Error("Expected closed indicator '[x]' in output")
	}
}

func TestOutputSubListTable_CrossRepo(t *testing.T) {
	parent := &api.Issue{
		Number: 10,
		Title:  "Parent Issue",
		Repository: api.Repository{
			Owner: "owner",
			Name:  "repo",
		},
	}
	subIssues := []api.SubIssue{
		{Number: 1, Title: "Same repo task", State: "OPEN", Repository: api.Repository{Owner: "owner", Name: "repo"}},
		{Number: 100, Title: "Cross repo task", State: "OPEN", Repository: api.Repository{Owner: "other", Name: "project"}},
	}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := outputSubListTable(subIssues, parent)

	w.Close()
	os.Stdout = oldStdout

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	output, _ := io.ReadAll(r)
	outputStr := string(output)

	// Should show repo info for cross-repo sub-issues
	if !strings.Contains(outputStr, "other/project#100") {
		t.Error("Expected cross-repo sub-issue to show 'other/project#100'")
	}

	// Should also show repo for same-repo when there are cross-repo issues
	if !strings.Contains(outputStr, "owner/repo#1") {
		t.Error("Expected same-repo sub-issue to show 'owner/repo#1' when cross-repo exists")
	}
}

func TestOutputSubListTable_ProgressCalculation(t *testing.T) {
	parent := &api.Issue{
		Number:     10,
		Title:      "Parent Issue",
		Repository: api.Repository{Owner: "owner", Name: "repo"},
	}

	tests := []struct {
		name     string
		subs     []api.SubIssue
		expected string
	}{
		{
			name: "all open",
			subs: []api.SubIssue{
				{Number: 1, State: "OPEN", Repository: api.Repository{Owner: "owner", Name: "repo"}},
				{Number: 2, State: "OPEN", Repository: api.Repository{Owner: "owner", Name: "repo"}},
				{Number: 3, State: "OPEN", Repository: api.Repository{Owner: "owner", Name: "repo"}},
			},
			expected: "0/3 complete",
		},
		{
			name: "all closed",
			subs: []api.SubIssue{
				{Number: 1, State: "CLOSED", Repository: api.Repository{Owner: "owner", Name: "repo"}},
				{Number: 2, State: "CLOSED", Repository: api.Repository{Owner: "owner", Name: "repo"}},
			},
			expected: "2/2 complete",
		},
		{
			name: "mixed",
			subs: []api.SubIssue{
				{Number: 1, State: "CLOSED", Repository: api.Repository{Owner: "owner", Name: "repo"}},
				{Number: 2, State: "OPEN", Repository: api.Repository{Owner: "owner", Name: "repo"}},
				{Number: 3, State: "CLOSED", Repository: api.Repository{Owner: "owner", Name: "repo"}},
				{Number: 4, State: "OPEN", Repository: api.Repository{Owner: "owner", Name: "repo"}},
			},
			expected: "2/4 complete",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := outputSubListTable(tt.subs, parent)

			w.Close()
			os.Stdout = oldStdout

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			output, _ := io.ReadAll(r)
			outputStr := string(output)

			if !strings.Contains(outputStr, tt.expected) {
				t.Errorf("Expected progress '%s' in output, got: %s", tt.expected, outputStr)
			}
		})
	}
}

// ============================================================================
// Additional Command Flag Tests
// ============================================================================

func TestSubCreateCommand_FlagShorthands(t *testing.T) {
	cmd := NewRootCommand()
	subCmd, _, err := cmd.Find([]string{"sub", "create"})
	if err != nil {
		t.Fatalf("sub create command not found: %v", err)
	}

	tests := []struct {
		flag      string
		shorthand string
	}{
		{"parent", "p"},
		{"title", "t"},
		{"body", "b"},
		{"repo", "R"},
	}

	for _, tt := range tests {
		t.Run(tt.flag, func(t *testing.T) {
			flag := subCmd.Flags().Lookup(tt.flag)
			if flag == nil {
				t.Fatalf("Expected --%s flag to exist", tt.flag)
			}
			if flag.Shorthand != tt.shorthand {
				t.Errorf("Expected --%s shorthand to be '%s', got '%s'", tt.flag, tt.shorthand, flag.Shorthand)
			}
		})
	}
}
