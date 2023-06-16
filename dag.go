package main

import (
	"fmt"
	"os"

	"github.com/paulroper/dag/git"
	"github.com/paulroper/dag/log"

	"github.com/urfave/cli/v2"
)

func dag(repositoryPath string) (string, error) {
	var changedFiles = git.GetChangedFiles(repositoryPath)
	return fmt.Sprintf("Found %d files", len(changedFiles)), nil
}

func main() {
	app := &cli.App{
		Name:  "dag",
		Usage: "create a dependency graph for a specified repo",
		Action: func(c *cli.Context) error {
			debug := c.Bool("debug")
			repositoryPath := c.String("repository")

			logger := log.Logger{Debug: debug}
			logger.LogDebug(
				fmt.Sprintf("Repository is %s", repositoryPath),
			)

			_, err := dag(repositoryPath)
			if err != nil {
				logger.LogError("Oh no!")
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Value: false,
				Usage: "print debug messages",
			},

			&cli.StringFlag{
				Name:     "repository",
				Required: true,
				Value:    "",
				Usage:    "path to repository to create dag for",
			},
		},
	}

	app.Run(os.Args)
}
