package main

import (
	"os"

	"github.com/scooter-indie/gh-pm/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
