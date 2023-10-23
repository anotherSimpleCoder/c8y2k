package main

import (
	"c8y2k/pages"
	"fmt"
	"os"
)

var content = ""

func main() {
	args := os.Args

	if len(args) < 2 {
		content = pages.Help()
	} else {
		arg := args[1]

		switch {
		case arg == "help":
			content = pages.Help()

		case arg == "new-component":
			content = pages.NewComponent()

		case arg == "new-widget":
			content = pages.NewWidget()

		default:
			{
				content = fmt.Sprintf("Invalid command!\n%s", pages.Help())
			}
		}
	}

	fmt.Println(content)
}
