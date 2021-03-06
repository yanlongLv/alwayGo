package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/urfave/cli/v2"
)

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
			Action:  create,
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

func create(ctx *cli.Context) error {
	projectName := ctx.Args().Slice()[0]
	pwd, _ := os.Getwd()
	projectfile := path.Join(pwd, projectName)
	p := NewProject(pwd, projectfile, projectName)
	ct, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if err := p.Generate(ct); err != nil {
		return err
	}
	return nil
}

func build(ctx *cli.Context) error {
	fmt.Print("you are build")
	return nil
}
