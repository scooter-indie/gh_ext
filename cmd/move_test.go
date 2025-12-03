package cmd

import (
	"bytes"
	"testing"
)

func TestMoveCommand_Exists(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"move", "--help"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("move command should exist: %v", err)
	}

	output := buf.String()
	if !bytes.Contains([]byte(output), []byte("move")) {
		t.Error("Expected help output to mention 'move'")
	}
}

func TestMoveCommand_RequiresIssueNumber(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"move"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err == nil {
		t.Error("Expected error when issue number not provided")
	}
}

func TestMoveCommand_HasStatusFlag(t *testing.T) {
	cmd := NewRootCommand()
	moveCmd, _, err := cmd.Find([]string{"move"})
	if err != nil {
		t.Fatalf("move command not found: %v", err)
	}

	flag := moveCmd.Flags().Lookup("status")
	if flag == nil {
		t.Error("Expected --status flag to exist")
	}
}

func TestMoveCommand_HasPriorityFlag(t *testing.T) {
	cmd := NewRootCommand()
	moveCmd, _, err := cmd.Find([]string{"move"})
	if err != nil {
		t.Fatalf("move command not found: %v", err)
	}

	flag := moveCmd.Flags().Lookup("priority")
	if flag == nil {
		t.Error("Expected --priority flag to exist")
	}
}

func TestMoveCommand_RequiresAtLeastOneFlag(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"move", "123"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err == nil {
		t.Error("Expected error when no field flags provided")
	}
}

// Recursive operation tests

func TestMoveCommand_HasRecursiveFlag(t *testing.T) {
	cmd := NewRootCommand()
	moveCmd, _, err := cmd.Find([]string{"move"})
	if err != nil {
		t.Fatalf("move command not found: %v", err)
	}

	flag := moveCmd.Flags().Lookup("recursive")
	if flag == nil {
		t.Error("Expected --recursive flag to exist")
	}

	// Verify short flag
	if flag.Shorthand != "r" {
		t.Errorf("Expected --recursive shorthand to be 'r', got '%s'", flag.Shorthand)
	}
}

func TestMoveCommand_HasDepthFlag(t *testing.T) {
	cmd := NewRootCommand()
	moveCmd, _, err := cmd.Find([]string{"move"})
	if err != nil {
		t.Fatalf("move command not found: %v", err)
	}

	flag := moveCmd.Flags().Lookup("depth")
	if flag == nil {
		t.Error("Expected --depth flag to exist")
	}

	// Verify default value
	if flag.DefValue != "10" {
		t.Errorf("Expected --depth default to be 10, got '%s'", flag.DefValue)
	}
}

func TestMoveCommand_HasDryRunFlag(t *testing.T) {
	cmd := NewRootCommand()
	moveCmd, _, err := cmd.Find([]string{"move"})
	if err != nil {
		t.Fatalf("move command not found: %v", err)
	}

	flag := moveCmd.Flags().Lookup("dry-run")
	if flag == nil {
		t.Error("Expected --dry-run flag to exist")
	}
}

func TestMoveCommand_HasYesFlag(t *testing.T) {
	cmd := NewRootCommand()
	moveCmd, _, err := cmd.Find([]string{"move"})
	if err != nil {
		t.Fatalf("move command not found: %v", err)
	}

	flag := moveCmd.Flags().Lookup("yes")
	if flag == nil {
		t.Error("Expected --yes flag to exist")
	}

	// Verify short flag
	if flag.Shorthand != "y" {
		t.Errorf("Expected --yes shorthand to be 'y', got '%s'", flag.Shorthand)
	}
}

func TestMoveCommand_RecursiveHelpText(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"move", "--help"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("move help failed: %v", err)
	}

	output := buf.String()
	if !bytes.Contains([]byte(output), []byte("--recursive")) {
		t.Error("Expected help to mention --recursive flag")
	}
	if !bytes.Contains([]byte(output), []byte("sub-issues")) {
		t.Error("Expected help to mention sub-issues")
	}
}
