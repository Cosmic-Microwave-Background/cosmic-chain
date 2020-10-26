package main

import (
	"os"

	"github.com/enflow.io/enf1/cmd/enf1d/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
