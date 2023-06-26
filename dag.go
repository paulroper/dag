package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/paulroper/dag/deps"
	"github.com/paulroper/dag/git"
	"github.com/paulroper/dag/logging"
	"github.com/paulroper/dag/output"

	"github.com/urfave/cli/v2"
)

func dag(logger logging.Logger, repo git.RepositoryInterrogator) error {
	// Step one - Pull all the file changes on this branch from Git
	changedFiles, err := repo.GetChangedFiles()
	if err != nil {
		return errors.New("FAILED TO FETCH CHANGED FILES FROM REPO")
	}

	// Step two - Filter the changes to code in apps or libs
	// TODO: The paths to check can be moved into config
	filteredFiles := []string{}
	for _, changedFile := range changedFiles {
		if strings.HasPrefix(changedFile, "apps/") || strings.HasPrefix(changedFile, "libs/") {
			filteredFiles = append(filteredFiles, changedFile)
		}
	}

	// Step three - Load all the deps files in the repo so we have a full list of modules and their deps
	depsMap, err := deps.GetDepsMap(filteredFiles)
	if err != nil {
		return errors.New("FAILED TO LOAD DEPS MAP")
	}

	logger.LogDebug(fmt.Sprintf("Deps map is %v", depsMap))

	// Step four - Work out what we need to build
	modulesToBuild, err := deps.GetModulesToBuild(filteredFiles, depsMap)
	if err != nil {
		return errors.New("FAILED TO LOAD BUILD MAP")
	}

	logger.LogDebug(fmt.Sprintf("Modules to build are %v", modulesToBuild))
	logger.Log(fmt.Sprintf("Found %d modules to build", len(modulesToBuild)))

	// Step five - Write the build list to a file
	output.WriteToFile(modulesToBuild)

	return nil
}

func main() {
	app := &cli.App{
		Name:  "dag",
		Usage: "create a dependency graph for a specified repo",
		Action: func(c *cli.Context) error {
			debug := c.Bool("debug")
			repositoryPath := c.String("repository")

			baseBranch := c.String("baseBranch")
			workingBranch := c.String("workingBranch")

			log := logging.Log{Debug: debug}
			log.LogDebug(
				fmt.Sprintf("Repository is %s", repositoryPath),
			)

			repo := git.Repository{
				BaseBranch:     baseBranch,
				Log:            log,
				RepositoryPath: repositoryPath,
				WorkingBranch:  workingBranch,
			}

			err := dag(log, repo)
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

			&cli.StringFlag{
				Name:     "baseBranch",
				Required: true,
				Value:    "",
				Usage:    "base branch to compare your changes to",
			},

			&cli.StringFlag{
				Name:     "workingBranch",
				Required: true,
				Value:    "",
				Usage:    "branch containing changes you want to build a dag for",
			},
		},
	}

	app.Run(os.Args)
}
