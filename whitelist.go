package main

import (
	"fmt"
	"whitelist/Handlers"
)

func Init() int {
	config := Handlers.ReadJSONFile("Data/config.json")
	basetype := config["type"].(string)

	version := config["version"]

	fmt.Println("Running whitelist version:", version)
	fmt.Println("Type: ", basetype+"\n")

	return 0
}

func VerifyKey(key string) bool {
	info := Handlers.Getinfo()
	keys := Handlers.ReadJSONFile("Data/Keys.json")
	config := Handlers.ReadJSONFile("Data/Config.json")
	basetype := config["type"].(string)

	typing := info[basetype]
	hash := Handlers.Simple_sha256([]byte(key))

	return keys[hash] == typing
}
