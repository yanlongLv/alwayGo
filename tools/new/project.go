package main

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/alwayGo/tools/base"
)

const (
	serviceURL = "https://github.com/go-kratos/service-layout.git"
)

type project struct {
	dir         string
	fileName    string
	projectName string
}

func NewProject(dir, fileName, projectName string) *project {
	return &project{dir, fileName, projectName}
}

func (p *project) Generate(ctx context.Context) error {
	to := path.Join(p.dir, p.projectName)
	if _, err := os.Stat(to); !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists", p.projectName)
	}
	repo := base.NewRepo(serviceURL)
	mod, err := base.ModulePath(path.Join(repo.Path(), "go.mod"))
	if err != nil {
		return err
	}
	if err = repo.CopyTo(ctx, to, []string{mod, p.projectName}, []string{".git", ".github"}); err != nil {
		return err
	}
	os.Rename(path.Join(to, "cmd", "server"), path.Join(to, "cmd", p.projectName))

	fmt.Printf("creating service %s\n", p.projectName)
	return nil
}
