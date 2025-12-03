package cmd

import (
	"bytes"
	"testing"
)

func TestIntakeCommand(t *testing.T) {
	t.Run("has correct command structure", func(t *testing.T) {
		cmd := newIntakeCommand()

		if cmd.Use != "intake" {
			t.Errorf("expected Use to be 'intake', got %s", cmd.Use)
		}

		if cmd.Short == "" {
			t.Error("expected Short description to be set")
		}

		// Check aliases
		if len(cmd.Aliases) == 0 || cmd.Aliases[0] != "in" {
			t.Error("expected 'in' alias")
		}
	})

	t.Run("has required flags", func(t *testing.T) {
		cmd := newIntakeCommand()

		// Check --apply flag
		applyFlag := cmd.Flags().Lookup("apply")
		if applyFlag == nil {
			t.Error("expected --apply flag")
		}

		// Check --dry-run flag
		dryRunFlag := cmd.Flags().Lookup("dry-run")
		if dryRunFlag == nil {
			t.Error("expected --dry-run flag")
		}

		// Check --json flag
		jsonFlag := cmd.Flags().Lookup("json")
		if jsonFlag == nil {
			t.Error("expected --json flag")
		}
	})

	t.Run("command is registered in root", func(t *testing.T) {
		root := NewRootCommand()
		buf := new(bytes.Buffer)
		root.SetOut(buf)
		root.SetArgs([]string{"intake", "--help"})
		err := root.Execute()
		if err != nil {
			t.Errorf("intake command not registered: %v", err)
		}
	})
}

func TestIntakeOptions(t *testing.T) {
	t.Run("default options", func(t *testing.T) {
		opts := &intakeOptions{}

		if opts.apply {
			t.Error("apply should be false by default")
		}
		if opts.dryRun {
			t.Error("dryRun should be false by default")
		}
		if opts.json {
			t.Error("json should be false by default")
		}
	})
}
