package main

import (
	"fmt"
	"os"

	"github.com/juliankoehn/zonk/http"
	"github.com/juliankoehn/zonk/tmpl"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

var VERSION = "dev"

func runCli() *cli.App {
	app := cli.NewApp()
	app.Name = "zonk"
	app.Usage = "Zonk provides tools to convert file to golang"
	app.Version = VERSION
	app.Author = "Julian Koehn"
	app.Commands = []cli.Command{
		http.HTTPCommand,
		tmpl.TmplCommand,
	}

	return app
}

func runPrompt() error {
	prompt := promptui.Select{
		Label: "Select Command",
		Items: []string{"http"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return err
	}

	if result == "http" {
		return http.RunPrompt()
	}

	return nil
}

func main() {
	if len(os.Args) > 1 {
		if err := runCli().Run(os.Args); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		if err := runPrompt(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
