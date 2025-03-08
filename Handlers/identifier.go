package Handlers

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/denisbrodbeck/machineid"
)

func Simple_sha256(encode []byte) string {
	hash := sha256.Sum256(encode)
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
	hwid, err := machineid.ID()

	if err != nil {
		log.Fatal(err.Error())
	}

	return map[string]any{
		"ip":   getip(),
		"hwid": Simple_sha256([]byte(hwid)),
	}
}
