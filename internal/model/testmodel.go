package model

import (
	"fmt"
	"testing"
)

var (
	TestResult = "Hello, World_Model_Test"
)

// TestUser is a test user for testing
func TestCommand(t *testing.T) *Command {
	quely := fmt.Sprintf("#!/bin/bash\necho \"%s\"", TestResult)
	return &Command{
		Script: quely,
	}
}
