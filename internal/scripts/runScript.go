package scripts

import (
	"fmt"
	"os/exec"
)

// Run executes a bash script
func Run(script string) ([]byte, error) {
	const op = "scripts.Run"
	result, err := exec.Command("bash", "-c", script).Output()
	if err != nil || script == "" {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return result, nil
}
