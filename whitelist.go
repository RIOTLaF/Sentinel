package main

import (
	"fmt"
	"whitelist/Handlers"
)

/*
GENERATE A HASH WITH 30 WORDS EASILY

import "whitelist/gen"
gen.Salt(30)
*/

func Init() int {
	Handlers.Define_Salty("easel-usage-broom-draw-prescribe-hazy-arming-compost-exerciser-hexagon-unsubtly-excuse-uphold-revival-atrophy-identity-mutual-comfy-debating-grandkid")
	// Define a new Salty if you want a secure hash you can use the hash generator
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
