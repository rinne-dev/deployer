package main

import (
	"log"
	"os"
	"rinne.dev/deployer/app"
)

func main() {
	err := app.App().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
