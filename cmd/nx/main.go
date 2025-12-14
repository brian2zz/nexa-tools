package main

import (
	"fmt"
	"os"

	"nexa-tools/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		commands.ShowHelp()
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	if err := commands.Run(cmd, args); err != nil {
		fmt.Println("❌", err)
	}
}
