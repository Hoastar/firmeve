package main

import (
	"fmt"
	"github.com/firmeve/firmeve"
	"github.com/firmeve/firmeve/http"
	"github.com/firmeve/firmeve/kernel/contract"
	"github.com/firmeve/firmeve/render"
)

func main() {
	firmeve.RunWithSupportFunc(
		application,
		firmeve.WithConfigPath("../config.yaml"),
		firmeve.WithProviders([]contract.Provider{
			new(http.Provider),
		}),
		firmeve.WithCommands([]contract.Command{
			new(http.Command),
		}),
	)
}

func application(application contract.Application) {
	router := application.Resolve(`http.router`).(contract.HttpRouter)
	router.GET("/", func(c contract.Context) {
		fmt.Printf("%t", c.Application() == firmeve.Application)
		c.RenderWith(200, render.JSON, map[string]string{
			"ctx_application":    fmt.Sprintf("%p", c.Firmeve()),
			"global_application": fmt.Sprintf("%p", firmeve.Application),
		})
	})
}
