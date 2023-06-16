package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/paulroper/dag/git"
	"github.com/paulroper/dag/logging"

	"github.com/urfave/cli/v2"
)

func dag(log logging.Logger, repo git.RepositoryInterrogator) (string, error) {
	changedFiles, err := repo.GetChangedFiles()
	if err != nil {
		return "", errors.New("FAILED TO FETCH CHANGED FILES FROM REPO")
	}

	return fmt.Sprintf("Found %d files", len(changedFiles)), nil
}

func main() {
	app := &cli.App{
		Name:  "dag",
		Usage: "create a dependency graph for a specified repo",
		Action: func(c *cli.Context) error {
			debug := c.Bool("debug")
			repositoryPath := c.String("repository")

			log := logging.Log{Debug: debug}
			log.LogDebug(
				fmt.Sprintf("Repository is %s", repositoryPath),
			)

			repo := git.Repository{Log: log, RepositoryPath: repositoryPath}

			_, err := dag(log, repo)
			if err != nil {
				os.Exit(1)
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
