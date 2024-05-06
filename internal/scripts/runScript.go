package scripts

import (
	"fmt"
	"os/exec"
)

func Run(script string) ([]byte, error) { // TODO:  сделать лог, записать в него, в принципе лог возвращаеться и обрабатываеться в месте вызова
	const op = "srorage.sqlstore.runScript"
	result, err := exec.Command("bash", "-c", script).Output()
	if err != nil {
		// server.error(w, r, http.StatusConflict, err)
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// s.respond(w, r, code, map[string]string{"error": err.Error()})
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	return result, nil
}
