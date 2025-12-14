package commands

import "fmt"

// Version adalah versi resmi nx
// Update ini setiap release
const Version = "0.1.0"

func init() {
	Register("version", VersionCmd)
	Register("--version", VersionCmd)
}

func VersionCmd(args []string) error {
	fmt.Println("nx version", Version)
	return nil
}
