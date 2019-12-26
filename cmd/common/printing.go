package common

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Part1(format string, a ...interface{}) {
	fmt.Printf(
		"Part 1: %s\n",
		fmt.Sprintf(format, a...),
	)
}

func Part2(format string, a ...interface{}) {
	fmt.Printf(
		"Part 2: %s\n",
		fmt.Sprintf(format, a...),
	)
}

func Debug(format string, a ...interface{}) {
	fmt.Printf(
		"DEBUG: %s\n",
		fmt.Sprintf(format, a...),
	)
}

func Print(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

// ClearScreen clears the terminal screen. Only works on Linux.
func ClearScreen() {
	PanicIf(runtime.GOOS != `linux` && runtime.GOOS != `darwin`, "Only linux is supported")

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
