package commands

import (
	"fmt"

	"nexa-tools/internal/generator"
)

func init() {
	Register("create-project", CreateProject)
}

func CreateProject(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: nexa create-project <project-name>")
	}

	generator.CreateProject(args[0])
	return nil
}
