package cli

import (
	"errors"
	"fmt"
	"os"
)

type Runner interface {
	Name() string
	Init(args []string) error
	Run() error
}

func Root(args []string) error {
	if len(args) < 1 {
		return errors.New("you must pass a sub-command")
	}

	cmds := []Runner{
		NewGreeter(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		fmt.Println(cmd.Name(), os.Args[2])
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("unknown subcommand: %s", subcommand)
}
