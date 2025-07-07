package shellbar

import (
	"flag"
	"fmt"
	"os"
)

var Version = "0.1.0" // コンパイル時に -ldflags で上書き可能

type Command struct {
	run func() error
}

func (c *Command) Run() error {
	return c.run()
}

var (
	VersionCommand = &Command{
		run: func() error {
			fmt.Printf("sbar version %s\n", Version)
			return nil
		},
	}

	HelpCommand = &Command{
		run: func() error {
			fmt.Fprintf(os.Stderr, "Usage: %s <command> [args...]\n\n", os.Args[0])
			fmt.Fprintf(os.Stderr, "Options:\n")
			flag.PrintDefaults()
			os.Exit(1)
			return nil
		},
	}
)

func parseArgs() (*Command, error) {
	version := flag.Bool("version", false, "show version")
	flag.Parse()

	if *version {
		return VersionCommand, nil
	}

	if flag.NArg() == 0 {
		return HelpCommand, nil
	}

	return NewExternalCommand(flag.Arg(0), flag.Args()[1:]), nil
}

func NewExternalCommand(name string, args []string) *Command {
	return &Command{
		run: func() error {
			fmt.Printf("sbar: executing %s %v (not implemented yet)\n", name, args)
			return nil
		},
	}
}