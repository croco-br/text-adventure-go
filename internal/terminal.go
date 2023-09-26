package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var red = "\033[31m"
var blue = "\033[34m"
var reset = "\033[0m"
var green = "\033[32m"

func ProcessInput() string {
	var input string
	fmt.Print("\n> ")
	fmt.Scanln(&input)
	return strings.ToLower(input)
}

func ClearScreen() {
	cmd := exec.Command("clear") // for Unix-like systems
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Info(s string) {
	fmt.Printf("%s%s%s\n", reset, s, reset)
}

func Important(s string) {
	fmt.Printf("%s%s%s\n", blue, s, reset)
}

func Error(s string) {
	fmt.Printf("%s%s%s\n", red, s, reset)
}

func Success(s string) {
	fmt.Printf("%s%s%s\n", green, s, reset)
}
