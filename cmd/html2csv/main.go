package main

import (
	"log"

	app "github.com/vzvu3k6k/html2csv"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
