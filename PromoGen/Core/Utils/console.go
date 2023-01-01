package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Run(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Clear() {
	switch runtime.GOOS {
	case "darwin":
		Run("clear")
	case "linux":
		Run("clear")
	case "windows":
		Run("cmd", "/c", "cls")
	default:
		Run("clear")
	}
}

func Title(t ...any) {
	Run(fmt.Sprintf("title %v", t...))
}
