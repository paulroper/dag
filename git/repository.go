package git

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/paulroper/dag/logging"
)

type RepositoryInterrogator interface {
	GetChangedFiles() ([]string, error)
}

type Repository struct {
	Log            logging.Logger
	RepositoryPath string
}

func (repo Repository) GetChangedFiles() ([]string, error) {
	_, err := git.PlainOpen(repo.RepositoryPath)
	if err != nil {
		repo.Log.LogError((fmt.Sprintf("Failed to open repo: %s", err)))
		return []string{}, nil
	}

	return []string{}, nil
}
