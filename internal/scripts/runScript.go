package scripts

import (
	"fmt"
	"os/exec"
)

func Run(script string) ([]byte, error) { // TODO:  сделать лог, записать в него, в принципе лог возвращаеться и обрабатываеться в месте вызова
	const op = "scripts.Run"
	result, err := exec.Command("bash", "-c", script).Output()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return result, nil
}
