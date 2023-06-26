package git

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/paulroper/dag/logging"
)

type RepositoryInterrogator interface {
	GetChangedFiles() ([]string, error)
}

type Repository struct {
	Log            logging.Logger
	RepositoryPath string
	BaseBranch     string
	WorkingBranch  string
}

func (repo Repository) GetChangedFiles() ([]string, error) {
	r, err := git.PlainOpen(repo.RepositoryPath)
	if err != nil {
		repo.Log.Log((fmt.Sprintf("Failed to open repo: %s", err)))
		return []string{}, nil
	}

	workingBranchHead, err := r.Reference(
		normaliseGitRef(repo.WorkingBranch),
		false,
	)

	if err != nil {
		repo.Log.Log((fmt.Sprintf("Failed to get HEAD ref for branch %s: %s", repo.WorkingBranch, err)))
		return []string{}, nil
	}

	baseBranchHead, err := r.Reference(
		normaliseGitRef(repo.BaseBranch),
		false,
	)

	if err != nil {
		repo.Log.Log((fmt.Sprintf("Failed to get HEAD ref for branch %s: %s", repo.BaseBranch, err)))
		return []string{}, nil
	}

	workingBranchHeadCommit, _ := r.CommitObject(workingBranchHead.Hash())
	if err != nil {
		repo.Log.Log((fmt.Sprintf("Failed to get HEAD commit for branch %s: %s", repo.WorkingBranch, err)))
		return []string{}, nil
	}

	baseBranchHeadCommit, _ := r.CommitObject(baseBranchHead.Hash())
	if err != nil {
		repo.Log.Log((fmt.Sprintf("Failed to get HEAD commit for branch %s: %s", repo.BaseBranch, err)))
		return []string{}, nil
	}

	patch, _ := workingBranchHeadCommit.Patch(baseBranchHeadCommit)
	filePatches := patch.FilePatches()

	changedFiles := []string{}

	for _, filePatch := range filePatches {
		from, to := filePatch.Files()

		if from == nil {
			changedFiles = append(changedFiles, to.Path())
		} else {
			changedFiles = append(changedFiles, from.Path())
		}
	}

	repo.Log.LogDebug(fmt.Sprintf("Detected changes in files: %s", changedFiles))

	return changedFiles, nil
}

func normaliseGitRef(ref string) plumbing.ReferenceName {
	if strings.HasPrefix(ref, "origin/") {
		return plumbing.ReferenceName(fmt.Sprintf("refs/remotes/%s", ref))
	}

	return plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", ref))
}
