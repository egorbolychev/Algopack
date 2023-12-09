package main

import (
	"algopack/cmd/algopack"
	"log"
)

func main() {
	if err := algopack.RootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
