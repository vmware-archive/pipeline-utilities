package main

import (
	"fmt"
	"os"

	"github.com/pivotalservices/pipeline-utilities/commands"
	flags "github.com/jessevdk/go-flags"
)

func main() {
	parser := flags.NewParser(&commands.Manager, flags.HelpFlag)
	parser.NamespaceDelimiter = "-"

	_, err := parser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
