package model

import "testing"

// TestUser is a test user for testing
func TestCommand(t *testing.T) *Command {
	return &Command{
		Script: "#!/bin/bash\necho \"Hello, World\"",
	}
}
