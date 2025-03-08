package Handlers

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func ReadJSONFile(file string) map[string]any {
	myfile, err := os.Open(file)
	text, newerr := io.ReadAll(myfile)
	var table map[string]any

	if err != nil {
		log.Fatal("Error: " + err.Error())
	}

	if newerr != nil {
		log.Fatal("Error: " + newerr.Error())
	}

	n := json.Unmarshal(text, &table)

	if n != nil {
		log.Fatal(n.Error())
	}

	return table
}
