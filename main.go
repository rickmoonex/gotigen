package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rickmoonex/gotigen/internal/client"
	"github.com/rickmoonex/gotigen/internal/gen"
	"github.com/rickmoonex/gotigen/internal/parser"
)

func main() {
	collectionName := os.Args[1]

	client, err := client.NewClient("localhost", 9200, "FZbT0lGtYsLnRgp0gKKBPg")
	if err != nil {
		log.Fatal(err)
	}

	enums, err := client.GetEnums(collectionName)
	if err != nil {
		log.Fatal(err)
	}

	parsedEnums, err := parser.ParseThingsDBEnums(*enums)
	if err != nil {
		log.Fatal(err)
	}

	ems := *parsedEnums
	fmt.Println(string(gen.GenerateParsedEnums(ems[0])))
}
