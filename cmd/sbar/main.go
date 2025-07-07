package main

import (
	"os"

	"github.com/kawaken/shellbar"
)

func main() {
	s := &shellbar.Shellbar{}
	if err := s.Run(); err != nil {
		os.Exit(1)
	}
}