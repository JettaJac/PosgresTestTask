package model

import "testing"

// TestUser is a test user for testing
func TestCommand(t *testing.T) *Command {
	return &Command{
		Name: "test", //TODO:  времено, потом убрать
		Script:    "#!/bin/bash\necho \"Hello, World_Test\"",
	}
}
