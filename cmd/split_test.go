package cmd

import (
	"bytes"
	"testing"
)

func TestSplitCommand(t *testing.T) {
	t.Run("has correct command structure", func(t *testing.T) {
		cmd := newSplitCommand()

		if cmd.Use != "split <issue> [tasks...]" {
			t.Errorf("expected Use to be 'split <issue> [tasks...]', got %s", cmd.Use)
		}

		if cmd.Short == "" {
			t.Error("expected Short description to be set")
		}
	})

	t.Run("has required flags", func(t *testing.T) {
		cmd := newSplitCommand()

		// Check --from flag
		fromFlag := cmd.Flags().Lookup("from")
		if fromFlag == nil {
			t.Error("expected --from flag")
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
		root.SetArgs([]string{"split", "--help"})
		err := root.Execute()
		if err != nil {
			t.Errorf("split command not registered: %v", err)
		}
	})
}

func TestSplitOptions(t *testing.T) {
	t.Run("default options", func(t *testing.T) {
		opts := &splitOptions{}

		if opts.from != "" {
			t.Error("from should be empty by default")
		}
		if opts.dryRun {
			t.Error("dryRun should be false by default")
		}
		if opts.json {
			t.Error("json should be false by default")
		}
	})
}

func TestParseChecklist(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name: "simple checklist",
			input: `# Epic Title

Some description here.

## Tasks
- [ ] Task one
- [ ] Task two
- [ ] Task three
`,
			expected: []string{"Task one", "Task two", "Task three"},
		},
		{
			name: "mixed checked and unchecked",
			input: `- [x] Completed task
- [ ] Pending task
- [ ] Another pending
`,
			expected: []string{"Pending task", "Another pending"},
		},
		{
			name: "with nested content",
			input: `- [ ] Main task
  - Some notes
  - More notes
- [ ] Second task
`,
			expected: []string{"Main task", "Second task"},
		},
		{
			name:     "no checklist items",
			input:    "Just some text without any checklist",
			expected: []string{},
		},
		{
			name: "checklist with extra whitespace",
			input: `- [ ]   Task with leading space
- [ ]	Task with tab
`,
			expected: []string{"Task with leading space", "Task with tab"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseChecklist(tt.input)

			if len(result) != len(tt.expected) {
				t.Errorf("expected %d items, got %d", len(tt.expected), len(result))
				t.Errorf("got: %v", result)
				return
			}

			for i, expected := range tt.expected {
				if result[i] != expected {
					t.Errorf("item %d: expected %q, got %q", i, expected, result[i])
				}
			}
		})
	}
}
