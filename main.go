package main

import (
	"fmt"
	"go-cli/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Start")

	appCLI := app.Generate()

	if err := appCLI.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
