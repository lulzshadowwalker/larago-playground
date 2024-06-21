package main

import (
	"fmt"
	"os"

	"github.com/lulzshadowwalker/larago/internal/cli"
)

func main() {
	if err := cli.Root(os.Args[1:]); err != nil {
		// TODO: use Charm logging library
		fmt.Println(err)
		os.Exit(1)
	}
}
