package commands

import (
	"fmt"
	"os"
	"os/exec"
)

const NxModulePath = "github.com/brian2zz/nexa-tools/cmd/nx"

func init() {
	Register("update", Update)
}

func Update(args []string) error {
	fmt.Println("⬆️  Updating Nx...")

	cmd := exec.Command("go", "install", NxModulePath+"@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("✅ Nx updated successfully")
	fmt.Println("👉 Run `nx doctor` to verify your environment")

	return nil
}
