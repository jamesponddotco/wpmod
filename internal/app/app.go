// Package app is the main package for the application.
package app

import (
	"fmt"
	"os"

	"git.sr.ht/~jamesponddotco/wpmod/internal/meta"
	"github.com/urfave/cli/v2"
)

// Run is the entry point for the application.
func Run(args []string) int {
	app := cli.NewApp()
	app.Name = meta.Name
	app.Version = meta.Version
	app.Usage = meta.Description
	app.HideHelpCommand = true

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "path",
			Aliases:  []string{"p"},
			Usage:    "path to the WordPress installation",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "user",
			Aliases:  []string{"u"},
			Usage:    "user for file ownership",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "group",
			Aliases:  []string{"g"},
			Usage:    "group for file ownership",
			Required: true,
		},
		&cli.BoolFlag{
			Name:    "strict",
			Aliases: []string{"s"},
			Usage:   "enable strict file permission mode",
		},
	}

	app.Action = PermissionAction

	if err := app.Run(args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)

		return 1
	}

	return 0
}
