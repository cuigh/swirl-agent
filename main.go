package main

import (
	"github.com/cuigh/auxo/app"
	"github.com/cuigh/auxo/app/flag"
	_ "github.com/cuigh/auxo/cache/memory"
	"github.com/cuigh/auxo/data/valid"
	"github.com/cuigh/auxo/net/web"
	"github.com/cuigh/auxo/net/web/filter"
	"github.com/cuigh/swirl-agent/controller"
)

func main() {
	app.Name = "swirl-agent"
	app.Version = "0.0.1"
	app.Desc = "An extension tool for Swirl(https://github.com/cuigh/swirl)"
	app.Action = func(ctx *app.Context) {
		app.Run(server())
	}
	app.Flags.Register(flag.All)
	app.Start()
}

func server() *web.Server {
	ws := web.Auto()

	// set render
	ws.Validator = &valid.Validator{Tag: "valid"}

	// register global filters
	ws.Use(filter.NewRecover())

	// register controllers
	ws.Handle("/container", controller.Container())
	ws.Handle("/image", controller.Image())
	ws.Handle("/volume", controller.Volume())

	return ws
}
