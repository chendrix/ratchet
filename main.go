package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	"path/filepath"
)

func main() {
	cmd := &Cmd{}

	parser := flags.NewParser(cmd, flags.HelpFlag|flags.PassDoubleDash)
	parser.NamespaceDelimiter = "-"

	_, err := parser.Parse()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	cmd.Execute()
}


// Cmd represents ratchet's CLI
type Cmd struct {
	CmdPositionalArgs `positional-args:"yes"`
}

// CmdPositionalArgs represents the positional args for ratchet's CLI
type CmdPositionalArgs struct {
	RelativePath string `positional-arg-name:"PATH" default:"." Description:"Path to the project to run against"`
}

const (
	// RatchetFilename is the name of the file that consumers can provide to customize ratchet
	RatchetFilename = ".ratchet.yml"
)

// Execute runs the ratchet command
func (c *Cmd) Execute() {
	ratchetPath := filepath.Join(c.RelativePath, RatchetFilename)

	if _, err := os.Stat(ratchetPath); os.IsNotExist(err) {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Found %s in %s", RatchetFilename, c.RelativePath)
}