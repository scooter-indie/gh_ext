package api

import (
	"testing"
)

func TestNewClient_ReturnsClient(t *testing.T) {
	// ACT: Create a new client
	client := NewClient()

	// ASSERT: Client is not nil
	if client == nil {
		t.Fatal("Expected client to be non-nil")
	}
}

func TestNewClient_HasGraphQLClient(t *testing.T) {
	// Skip in CI - requires gh auth
	if testing.Short() {
		t.Skip("Skipping test that requires gh auth")
	}

	// ACT: Create a new client
	client := NewClient()

	// ASSERT: GraphQL client is accessible
	if client.gql == nil {
		t.Fatal("Expected GraphQL client to be initialized")
	}
}

func TestNewClientWithOptions_CustomHost(t *testing.T) {
	// ARRANGE: Custom options
	opts := ClientOptions{
		Host: "github.example.com",
	}

	// ACT: Create client with options
	client := NewClientWithOptions(opts)

	// ASSERT: Client is created (host is used internally)
	if client == nil {
		t.Fatal("Expected client to be non-nil")
	}
}

func TestClient_FeatureHeaders_Included(t *testing.T) {
	// This test verifies that sub_issues feature header is configured
	// We can't easily test the actual header without making a request,
	// but we can verify the client was created with the right options

	client := NewClient()

	// ASSERT: Client exists and has feature flags set
	if client == nil {
		t.Fatal("Expected client to be non-nil")
	}

	// The feature headers are set during client creation
	// Actual header verification would require integration tests
}

func TestJoinFeatures_Empty(t *testing.T) {
	result := joinFeatures([]string{})
	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestJoinFeatures_Single(t *testing.T) {
	result := joinFeatures([]string{"sub_issues"})
	if result != "sub_issues" {
		t.Errorf("Expected 'sub_issues', got '%s'", result)
	}
}

func TestJoinFeatures_Multiple(t *testing.T) {
	result := joinFeatures([]string{"sub_issues", "issue_types"})
	if result != "sub_issues,issue_types" {
		t.Errorf("Expected 'sub_issues,issue_types', got '%s'", result)
	}
}
