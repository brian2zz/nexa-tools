package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {
	Register("install", Install)
}

func Install(args []string) error {
	fmt.Println("🔧 Installing Nexa project CLI...")

	// Pastikan di root project
	if _, err := os.Stat("go.mod"); err != nil {
		return fmt.Errorf("go.mod not found. run this command from project root")
	}

	if _, err := os.Stat("cmd/nexa/main.go"); err != nil {
		return fmt.Errorf("this is not a Nexa project (cmd/nexa not found)")
	}

	cmd := exec.Command("go", "install", "./cmd/nexa")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("✅ Nexa project CLI installed")
	fmt.Println("👉 You can now run: nexa serve")

	return nil
}
