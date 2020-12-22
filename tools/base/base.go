package base

import (
	"context"
	"errors"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git/v5"
)

//Repo is new repo
type Repo struct {
	url  string
	home string
}

//NewRepo is new Repo install
func NewRepo(url string) *Repo {
	return &Repo{
		url:  url,
		home: "./",
	}
}

//Path get repo path
func (r *Repo) Path() string {
	start := strings.LastIndex(r.url, "/")
	end := strings.LastIndex(r.url, ".git")
	return path.Join(r.home, r.url[start+1:end])
}

// Pull 根据本地分支代码更新远程分支
func (r *Repo) Pull(ctx context.Context, url string) error {
	repo, err := git.PlainOpen(r.Path())
	if err != nil {
		return err
	}
	workTree, err := repo.Worktree()
	if err != nil {
		return err
	}

	if err = workTree.PullContext(ctx, &git.PullOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
	}); errors.Is(err, git.NoErrAlreadyUpToDate) {
		return nil
	}
	return nil
}

// Clone remote branch
func (r *Repo) Clone(ctx context.Context) error {
	if _, err := os.Stat(r.Path()); !os.IsNotExist(err) {
		return r.Pull(ctx, r.url)
	}
	_, err := git.PlainCloneContext(ctx, r.Path(), false, &git.CloneOptions{
		URL:      r.url,
		Progress: os.Stdout,
	})
	return err
}

// CopyTo clone failure need to copy
func (r *Repo) CopyTo(ctx context.Context, to string, replaces, ignores []string) error {
	if err := r.Clone(ctx); err != nil {
		return err
	}
	return r.CopyTo(r.Path(), to)
}
