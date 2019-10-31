package main

import (
	"log"

	"github.com/gotoolkit/peony"
)

func main() {

	app, err := peony.New()
	if err != nil {
		log.Fatal(err)
	}
	app.Start()
}
