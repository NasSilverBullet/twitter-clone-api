package main

import (
	"fmt"
	"os"

	"github.com/NasSilverBullet/twitter-clone-api/app/frameworks"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func run() error {
	logger := frameworks.NewLogger(os.Stdout, os.Stderr)
	if err := frameworks.LoadEnv(logger); err != nil {
		return err
	}

	sqlHandler, err := frameworks.NewSQLHandler()
	if err != nil {
		return err
	}

	if err := frameworks.Routes(sqlHandler); err != nil {
		return err
	}

	return nil
}
