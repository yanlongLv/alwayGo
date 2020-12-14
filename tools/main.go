package main

import (
	"fmt"
	"os"
	"path"

	"github.com/alwayGo/tools/pkg"
	"github.com/urfave/cli/v2"
)

// /Users/yanlonglv/go/bin/packr2
func main() {
	app := cli.NewApp()
	app.Name = "create new project"
	app.Usage = "init project and product project directory"
	app.Version = "1.0.0"
	app.Commands = []*cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "create new project and product new directory",
			Action:  newProject,
		}, {
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build project",
			Action:  build,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func newProject(ctx *cli.Context) error {
	projectName := ctx.Args().Slice()[0]
	pwd, _ := os.Getwd()
	projectfile := path.Join(pwd, projectName)
	return pkg.Create(projectfile, projectName)
}

func build(ctx *cli.Context) error {
	fmt.Print("you are build")
	return nil
}
