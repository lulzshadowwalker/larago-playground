package cli

import (
	"flag"
	"fmt"
)

type Greeter struct {
	fs   *flag.FlagSet
	name string
}

func NewGreeter() Greeter {
	g := Greeter{
		name: "greet",
		fs:   flag.NewFlagSet("greet", flag.ContinueOnError),
	}

	g.fs.StringVar(&g.name, "name", "World", "name of the person to be greeted")
	return g
}

func (g Greeter) Name() string {
	return g.name
}

func (g Greeter) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g Greeter) Run() error {
	fmt.Printf("hello, %s", g.name)
	return nil
}
