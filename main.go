package main

import (
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/husobee/vestigo"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "users-api"
	app.Usage = ""

	app.Action = func(c *cli.Context) error {
		serve()
		return nil
	}

	app.Run(os.Args)
}

func serve() {
	r := vestigo.NewRouter()

	r.Get("/__health", Health())
	r.Get("/__gtg", GTG())

	http.Handle("/", r)
	log.Info("Users API started!")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.WithError(err).Panic("Couldn't set up HTTP listener")
	}
}
