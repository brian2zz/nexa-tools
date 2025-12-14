package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func init() {
	Register("doctor", Doctor)
}

func Doctor(args []string) error {
	fmt.Println("🩺 Nx Doctor")
	fmt.Println("────────────────────────────")

	ok := true

	// ======================
	// OS & ARCH
	// ======================
	fmt.Printf("OS:   %s\n", runtime.GOOS)
	fmt.Printf("ARCH: %s\n\n", runtime.GOARCH)

	// ======================
	// Go check
	// ======================
	goBin, err := exec.LookPath("go")
	if err != nil {
		fmt.Println("❌ Go: NOT FOUND")
		fmt.Println("   👉 Install Go from https://go.dev/dl/")
		ok = false
	} else {
		fmt.Println("✅ Go: found at", goBin)

		versionOut, err := exec.Command("go", "version").Output()
		if err != nil {
			fmt.Println("❌ Go version: failed to detect")
			ok = false
		} else {
			version := strings.TrimSpace(string(versionOut))
			fmt.Println("ℹ️  ", version)

			if !isGoVersionSupported(version) {
				fmt.Println("⚠️  Go version < 1.22 (recommended)")
				ok = false
			}
		}
	}

	fmt.Println()

	// ======================
	// GOPATH / GOBIN
	// ======================
	goPath := os.Getenv("GOPATH")
	goBinEnv := os.Getenv("GOBIN")

	if goPath == "" {
		fmt.Println("⚠️  GOPATH: not set (default will be used)")
	} else {
		fmt.Println("✅ GOPATH:", goPath)
	}

	if goBinEnv != "" {
		fmt.Println("✅ GOBIN:", goBinEnv)
	} else if goPath != "" {
		fmt.Println("ℹ️  GOBIN: using", goPath+"/bin")
	} else {
		fmt.Println("⚠️  GOBIN: not detected")
		ok = false
	}

	fmt.Println()

	// ======================
	// PATH check
	// ======================
	pathEnv := os.Getenv("PATH")
	binPath := goBinEnv

	if binPath == "" && goPath != "" {
		binPath = goPath + string(os.PathSeparator) + "bin"
	}

	if binPath != "" && strings.Contains(pathEnv, binPath) {
		fmt.Println("✅ PATH contains Go bin:", binPath)
	} else {
		fmt.Println("⚠️  PATH does NOT contain Go bin")
		fmt.Println("   👉 Add Go bin to PATH:", binPath)
		ok = false
	}

	fmt.Println("────────────────────────────")

	if ok {
		fmt.Println("✅ Environment looks good!")
	} else {
		fmt.Println("⚠️  Issues detected. Fix them before continuing.")
	}

	return nil
}

// ======================================================
// Utils
// ======================================================

func isGoVersionSupported(version string) bool {
	// contoh: "go version go1.22.1 windows/amd64"
	parts := strings.Fields(version)
	if len(parts) < 3 {
		return false
	}

	v := strings.TrimPrefix(parts[2], "go")
	vParts := strings.Split(v, ".")
	if len(vParts) < 2 {
		return false
	}

	major := vParts[0]
	minor := vParts[1]

	if major != "1" {
		return false
	}

	return minor >= "22"
}
