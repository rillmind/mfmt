// Means my formating
package mfmt

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

// A simple scanner that scans a intire line from the standard input and returns a string
func Scanln() string {
	reader := bufio.NewReader(os.Stdin)

	texto, _ := reader.ReadString('\n')

	return strings.TrimSpace(texto)
}

// Cleans the console/terminal
func Cleanse() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}
