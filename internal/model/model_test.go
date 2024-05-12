package model_test

import (
	// "errors"
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

// func TestCommand_Validate2(t *testing.T) {

// 	testCases := []struct {
// 		name    string
// 		command func() *model.Command
// 		isValid bool
// 	}{
// 		{
// 			name: "valid",
// 			command: func() *model.Command {
// 				c := model.TestCommand(t)
// 				c.Script = "#!/bin/bash\necho \"Hello, World Test!!!\""
// 				return c
// 			},
// 			isValid: true,
// 		},
// 		{
// 			name: "empty script",
// 			command: func() *model.Command {
// 				c := model.TestCommand(t)
// 				c.Script = ""
// 				return c
// 			},
// 			isValid: false,
// 		},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			validate := validator.New()
// 			err := validate.Struct(tc.command)
// 			if tc.isValid {
// 				assert.NoError(t, err)
// 			} else {
// 				assert.Error(t, err)
// 			}
// 		})

// 	}
// }
