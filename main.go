package main

import (
	"fmt"
	"strings"
)

func repeat() {
	fmt.Println("Enter a key")
	fmt.Println(`Type "Close" to finish the program`)

	var key string

	fmt.Scanln(&key)
	if VerifyKey(key) {
		fmt.Println("Sucess")
	} else {
		fmt.Println("Fail, try again")
	}

	if strings.ToLower(key) != "close" {
		repeat()
	}
}

func main() {
	Init()
	repeat()
}
