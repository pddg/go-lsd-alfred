package main

import (
	"fmt"
	"encoding/json"
	"os"
	"github.com/urfave/cli"
	"github.com/pddg/go-lsd-alfred/models"
	sc "github.com/pddg/go-lsd-alfred/scraping"
)

func main() {
	app := cli.NewApp()
	app.Name = "alfred-lsd"
	app.Usage = "Search a given word or phrase in Life Science Dictionary (https://lsd-project.jp/weblsd/)"
	app.Version = "0.1.0"
	app.Commands = []cli.Command {
		{
			Name:		"begin",
			Category:	"translate",
			Usage:		"Search word beginning with a given argument.",
			Action: 	search,
		},
		{
			Name:		"end",
			Category:	"translate",
			Usage:		"Search word ending with a given argument.",
			Action: 	search,
		},
		{
			Name:		"include",
			Category:	"translate",
			Usage:		"Search word containing a given argument.",
			Action: 	search,
		},
		{
			Name:		"equal",
			Category:	"translate",
			Usage:		"Search word equaring a given argument.",
			Action: 	search,
		},
	}
	app.Run(os.Args)
}

func search(c *cli.Context) {
	resp := new(models.Response)
	doc, err := sc.GetPage(c.Command.Name, c.Args().First())
	if err != nil {
		resp.Items = append(resp.Items, *sc.CreateError(err, "Couldn't get web page."))
	} else {
		sc.ScrapeMeaning(doc, &resp.Items)
		resp.Items = append(resp.Items, *sc.CreateOrigin(c.Command.Name, c.Args().First()))
	}
	r, err := json.Marshal(resp)
	if err != nil {
		resp.Items = append(resp.Items, *sc.CreateError(err, "Couldn't parse json."))
	}
	fmt.Println(string(r))
}
