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
	Salt := "easel-usage-broom-draw-prescribe-hazy-arming-compost-exerciser-hexagon-unsubtly-excuse-uphold-revival-atrophy-identity-mutual-comfy-debating-grandkid"
	Handlers.Define_Salty(Salt)
	// Define a new Salty if you want a secure hash you can use the hash generator
	config := Handlers.ReadJSONFile("Data/config.json")
	basetype := config["type"].(string)
	version := config["version"]
	debug := config["debug"].(bool)

	fmt.Println("Running whitelist version:", version)
	fmt.Println("Type: ", basetype+"\n")

	if debug {
		info := Handlers.Getinfo()
		typing := info[basetype]
		fmt.Println("HWID: ", typing)
		fmt.Println("Salt: " + Salt)
	}

	return 0
}

func VerifyKey(key string) bool {
	info := Handlers.Getinfo()
	keys := Handlers.ReadJSONFile("Data/Keys.json")
	config := Handlers.ReadJSONFile("Data/Config.json")
	basetype := config["type"].(string)
	debug := config["debug"].(bool)

	typing := info[basetype]
	hash := Handlers.Simple_sha256([]byte(key))

	if debug {
		fmt.Println("Validing key: " + key)
		fmt.Println("Hash form: " + hash)
	}

	return keys[hash] == typing
}
