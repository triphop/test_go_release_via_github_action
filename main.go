package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("Hi, DEV World! 😉")

	// Print with default helper functions
	color.Cyan("Prints text in cyan.")

	// A newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")

	// These are using the default foreground colors
	color.Red("We have red")
	color.Magenta("And many others ..")
}
