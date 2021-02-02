package util

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func ClearConsole() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func Scan(text string) string {
	fmt.Print(text)
	var input string
	fmt.Scanln(&input)
	return input
}

func ScanAsIntWithDefault(text string, defaultValue int) int {
	numericText := Scan(text)
	number, err := strconv.Atoi(numericText)
	if err != nil {
		return defaultValue
	}
	return number
}
