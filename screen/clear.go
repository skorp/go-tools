package screen

import (
	"os"
	"os/exec"
	"runtime"
)

// Clear runs the command clear / cls
func Clear() {
	var c *exec.Cmd
	switch os := runtime.GOOS; os {
	case "darwin", "linux":
		c = exec.Command("clear")
	default:
		c = exec.Command("cls")
	}
	c.Stdout = os.Stdout
	c.Run()
}
