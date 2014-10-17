package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"strings"

	"github.com/codegangsta/cli"
)

// newCommand returns a command populated from the context
func newCommand(context *cli.Context) (c command, err error) {
	var (
		cmd  string
		args = []string(context.Args())
	)

	switch len(args) {
	default:
		cmd = strings.Join(args, " ")
	case 0:
		raw, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return c, err
		}
		cmd = string(raw)
	}

	if cmd == "" {
		return c, fmt.Errorf("no command specified")
	}
	return command{
		Cmd:      cmd,
		User:     context.GlobalString("user"),
		Identity: context.GlobalString("identity"),
	}, nil
}

// command to run over an SSH connection
type command struct {
	// User is the user to run the command as
	User string

	// Cmd is the pared command string that will be executed
	Cmd string

	// Identity is the SSH key to identify as which is commonly
	// the private keypair i.e. id_rsa
	Identity string
}

// String returns a pretty printed string of the command
func (c command) String() string {
	return fmt.Sprintf("user: %s command: %s", c.User, c.Cmd)
}
