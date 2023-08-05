package main

import (
	"log"
	"os"

	"github.com/integralist/fastly-cli/pkg/app"
)

func main() {
	if err := app.Run(os.Args, os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
