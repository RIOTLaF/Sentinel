package Handlers

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var Salty string = "deflator-wackiness-widow-revisit-tricky-freckled-prepaid-tarnish-stopping-zesty-gauze-smog-uncapped-curliness-defiling-chance-pursuit-boots-recoil-cure"

// Default salty

func Define_Salty(NewSalty string) {
	count := strings.Count(NewSalty, "-")
	if count < 8 {
		fmt.Println("Not a secure hash using default")
		return
	}
	Salty = NewSalty
}

func Simple_sha256(encode []byte) string {
	localsalt := []byte(Salty)
	hash := sha256.Sum256(append(localsalt, encode...))
	return fmt.Sprintf("%x", hash)
}

func getip() string {
	api, err := http.Get("https://api.ipify.org/")
	if err != nil {
		log.Fatal(err.Error())
	}

	encoded, newerr := io.ReadAll(api.Body)
	if newerr != nil {
		log.Fatal(newerr.Error())
	}

	return Simple_sha256(encoded)
}

func Getinfo() map[string]any {
	return map[string]any{
		"ip": getip(),
	}
}
