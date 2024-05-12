package model_test

import (
	"main/internal/model"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestCommand_Validate(t *testing.T) {

	testCases := []struct {
		name    string
		command interface{}
		isValid bool
	}{
		{
			name: "valid",
			command: model.Command{
				Script: "#!/bin/bash\necho \"Hello, World Test!!!\"",
			},
			isValid: true,
		},
		{
			name: "empty script",
			command: model.Command{
				Script: "",
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validate := validator.New()
			err := validate.Struct(tc.command)
			if tc.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})

	}
}
