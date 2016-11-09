package main

import (
	"log"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/tylerb/graceful"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
)

func run(ctx *cli.Context) error {
	log.Print("Starting...")
	log.Printf("Using config: %s", ctx.String("config"))
	baseRouter := httprouter.New()

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.UseHandler(baseRouter)
	timeout := 10 * time.Second
	graceful.Run(ctx.String("port"), timeout, n)

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "botrocity"
	app.Author = "David J Felix <felix.davidj@gmail.com>"
	app.Usage = "Run a bot with fixed responses for mattermost"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "port, p",
			Value: ":8080",
			Usage: "the host address to run botrocity on.",
		},
		cli.StringFlag{
			Name:  "config, f",
			Value: "botrocity.json",
			Usage: "The config file location.",
		},
	}
	app.Run(os.Args)
}
