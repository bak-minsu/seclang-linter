package main

import (
	"fmt"
	"os"

	"github.com/bak-minsu/seclang-linter/cmd/seclang-linter/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}
