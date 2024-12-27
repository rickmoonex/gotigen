package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rickmoonex/gotigen/internal/client"
	"github.com/rickmoonex/gotigen/internal/parser"
)

func main() {
	collectionName := os.Args[1]

	client, err := client.NewClient("localhost", 9200, "FZbT0lGtYsLnRgp0gKKBPg")
	if err != nil {
		log.Fatal(err)
	}

	types, err := client.GetTypes(collectionName)
	if err != nil {
		log.Fatal(err)
	}

	parsedTypes := parser.ParseThingsDBTypes(*types)

	data, _ := json.Marshal(parsedTypes)

	fmt.Println(string(data))
}
