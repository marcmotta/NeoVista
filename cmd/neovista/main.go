// cmd/neovista/main.go
package main

import (
	"flag"
	"log"
	"os"

	"neovista/internal/neovista"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := neovista.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
