package commands

import "fmt"

type Command func(args []string) error

var registry = map[string]Command{}

func Register(name string, cmd Command) {
	registry[name] = cmd
}

func Run(name string, args []string) error {
	if cmd, ok := registry[name]; ok {
		return cmd(args)
	}
	return fmt.Errorf("unknown command: %s", name)
}

func ShowHelp() {
	fmt.Println("Nx — Nexa Project Manager")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  nx <command>")
	fmt.Println("")
	fmt.Println("Available commands:")
	for name := range registry {
		// sembunyikan flag internal
		if name == "--version" {
			continue
		}
		fmt.Println(" ", name)
	}
	fmt.Println("")
	fmt.Println("Flags:")
	fmt.Println("  --version   Show nx version")
}

