package main

import (
	"log"

	"github.com/gotoolkit/peony"
)

func main() {
	app := peony.New()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
