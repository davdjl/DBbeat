package main

import (
	"os"

	"github.com/davdjl/dbbeat/cmd"

	_ "github.com/davdjl/dbbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
