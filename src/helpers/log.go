package helpers

import (
	"fmt"

	"github.com/ttacon/chalk"
)

// Log logs in green.
func Log(s string) {
	fmt.Println(chalk.Green.Color(s))
}

// LogErr logs in red.
func LogErr(s string) {
	fmt.Println(chalk.Red.Color(s))
}

// LogInfo logs in magenta.
func LogInfo(s string) {
	fmt.Println(chalk.Magenta.Color(s))
}

// LogWarn logs in yellow.
func LogWarn(s string) {
	fmt.Println(chalk.Yellow.Color(s))
}
