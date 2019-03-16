package main

import (
	"log"
	"os"

	"github.com/maximelamure/restli-go/generator"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("<output directory> <pdsc/idl directory>...")
	}

	err := generator.Generate(os.Args[1], os.Args[2:])
	if err != nil {
		log.Println(err)
	}
}
