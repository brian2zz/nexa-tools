package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func CreateProject(name string) {
	fmt.Println("🚀 Creating project:", name)

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("❌ Cannot get working directory:", err)
		return
	}

	baseDir := filepath.Dir(wd)
	projectRoot := filepath.Join(baseDir, name)

	if _, err := os.Stat(projectRoot); err == nil {
		fmt.Println("❌ Project already exists:", projectRoot)
		return
	}

	// ======================
	// Folder structure
	// ======================
	folders := []string{
		"cmd/server",
		"cmd/nexa/commands",
		"app/routes",
		"app/handlers",
		"app/services",
		"app/middleware",
		"app/config",
		"app/models",
		"model/schema",
	}

	for _, folder := range folders {
		if err := os.MkdirAll(filepath.Join(projectRoot, folder), 0755); err != nil {
			fail("Failed to create folder: " + folder)
		}
	}

	// ======================
	// Files
	// ======================
	must(createGoMod(projectRoot, name))
	must(createServerMain(projectRoot, name))
	must(createRoutes(projectRoot))
	must(createProjectNexaMain(projectRoot, name))
	must(createProjectCommandRegistry(projectRoot))
	must(createServeCommand(projectRoot))
	must(createEnvExample(projectRoot))
	must(createConfig(projectRoot))

	// ======================
	// go mod tidy
	// ======================
	fmt.Println("📦 Running go mod tidy...")
	must(runGoModTidy(projectRoot))

	fmt.Println("✅ Project created successfully at:", projectRoot)
}

/*
=====================================================
GENERATORS — SEMUA FILE PROJECT DIBUAT DI SINI
=====================================================
*/

func createGoMod(projectRoot, moduleName string) error {
	content := fmt.Sprintf(`module %s

go 1.22

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/joho/godotenv v1.5.1
)
`, moduleName)

	return os.WriteFile(filepath.Join(projectRoot, "go.mod"), []byte(content), 0644)
}

func createServerMain(projectRoot, moduleName string) error {
	content := fmt.Sprintf(`package main

import (
	"%s/app/config"
	"%s/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":" + cfg.AppPort)
}
`, moduleName, moduleName)

	return os.WriteFile(filepath.Join(projectRoot, "cmd/server/main.go"), []byte(content), 0644)
}

func createRoutes(projectRoot string) error {
	content := `package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
`
	return os.WriteFile(filepath.Join(projectRoot, "app/routes/routes.go"), []byte(content), 0644)
}

/*
======================
PROJECT NEXA CLI
======================
*/

func createProjectNexaMain(projectRoot, moduleName string) error {
	content := fmt.Sprintf(`package main

import (
	"fmt"
	"os"

	"%s/cmd/nexa/commands"
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
`, moduleName)

	return os.WriteFile(
		filepath.Join(projectRoot, "cmd/nexa/main.go"),
		[]byte(content),
		0644,
	)
}

func createProjectCommandRegistry(projectRoot string) error {
	content := `package commands

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
	fmt.Println("Nexa Project CLI")
	fmt.Println("")
	fmt.Println("Available commands:")
	for name := range registry {
		fmt.Println(" ", name)
	}
}
`
	return os.WriteFile(
		filepath.Join(projectRoot, "cmd/nexa/commands/registry.go"),
		[]byte(content),
		0644,
	)
}

func createServeCommand(projectRoot string) error {
	content := `package commands

import (
	"os"
	"os/exec"
)

func init() {
	Register("serve", Serve)
}

func Serve(args []string) error {
	cmd := exec.Command("go", "run", "cmd/server/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
`
	return os.WriteFile(
		filepath.Join(projectRoot, "cmd/nexa/commands/serve.go"),
		[]byte(content),
		0644,
	)
}

/*
======================
ENV & CONFIG
======================
*/

func createEnvExample(projectRoot string) error {
	content := `APP_NAME=MyApp
APP_ENV=development
APP_PORT=8080
`
	return os.WriteFile(filepath.Join(projectRoot, ".env.example"), []byte(content), 0644)
}

func createConfig(projectRoot string) error {
	content := `package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	AppEnv  string
	AppPort string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppName: getEnv("APP_NAME", "NexaApp"),
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: getEnv("APP_PORT", "8080"),
	}

	log.Printf("Config loaded: %s (%s)", cfg.AppName, cfg.AppEnv)
	return cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
`
	return os.WriteFile(filepath.Join(projectRoot, "app/config/config.go"), []byte(content), 0644)
}

/*
======================
UTILS
======================
*/

func runGoModTidy(projectRoot string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func must(err error) {
	if err != nil {
		fail(err.Error())
	}
}

func fail(msg string) {
	fmt.Println("❌", msg)
	os.Exit(1)
}
