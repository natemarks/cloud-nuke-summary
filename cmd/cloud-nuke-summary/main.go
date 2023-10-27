package main

import (
	"log"
	"os"

	"github.com/natemarks/cloud-nuke-summary/summary"
)

func main() {
	argsWithoutProg := os.Args[1:]
	contents, err := summary.GetContentsFromFile(argsWithoutProg[0])
	if err != nil {
		log.Fatal(err)
	}
	summary.PrintReport(contents)
}
