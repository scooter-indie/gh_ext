package cmd

import (
	"bytes"
	"testing"
)

func TestTriageCommand(t *testing.T) {
	t.Run("has correct command structure", func(t *testing.T) {
		cmd := newTriageCommand()

		if cmd.Use != "triage [config-name]" {
			t.Errorf("expected Use to be 'triage [config-name]', got %s", cmd.Use)
		}

		if cmd.Short == "" {
			t.Error("expected Short description to be set")
		}
	})

	t.Run("has required flags", func(t *testing.T) {
		cmd := newTriageCommand()

		// Check --dry-run flag
		dryRunFlag := cmd.Flags().Lookup("dry-run")
		if dryRunFlag == nil {
			t.Error("expected --dry-run flag")
		}

		// Check --interactive flag
		interactiveFlag := cmd.Flags().Lookup("interactive")
		if interactiveFlag == nil {
			t.Error("expected --interactive flag")
		}

		// Check --json flag
		jsonFlag := cmd.Flags().Lookup("json")
		if jsonFlag == nil {
			t.Error("expected --json flag")
		}

		// Check --list flag
		listFlag := cmd.Flags().Lookup("list")
		if listFlag == nil {
			t.Error("expected --list flag")
		}
	})

	t.Run("command is registered in root", func(t *testing.T) {
		root := NewRootCommand()
		buf := new(bytes.Buffer)
		root.SetOut(buf)
		root.SetArgs([]string{"triage", "--help"})
		err := root.Execute()
		if err != nil {
			t.Errorf("triage command not registered: %v", err)
		}
	})
}

func TestTriageOptions(t *testing.T) {
	t.Run("default options", func(t *testing.T) {
		opts := &triageOptions{}

		if opts.dryRun {
			t.Error("dryRun should be false by default")
		}
		if opts.interactive {
			t.Error("interactive should be false by default")
		}
		if opts.json {
			t.Error("json should be false by default")
		}
		if opts.list {
			t.Error("list should be false by default")
		}
	})
}
