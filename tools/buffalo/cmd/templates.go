package cmd

var mainGo = `
package main

import (
	"log"	
	"time"

	"{{.}}/app"

	"github.com/gobuffalo/buffalo/servers"
)

const (
	// responseTimeout used for request Read, Write and Idle.
	responseTimeout = 30 * time.Second
)

// We initialize some server settings to avoid long running
// requests that would kill our DB.
var server = &servers.Simple{
	Server: &http.Server{
		ReadTimeout:  responseTimeout,
		WriteTimeout: responseTimeout,
		IdleTimeout:  responseTimeout,
	},
}

func main() {
	bapp := app.New()
	if err := bapp.Serve(server); err != nil {
		log.Fatal(err)
	}
}
`
