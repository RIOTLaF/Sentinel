package gen

import (
	"fmt"
	"whitelist/phrase"
)

func Salt(count int) string {
	ps := phrase.Default
	Default := 8
	if count < Default {
		fmt.Println("Not enough | Using default")
	} else {
		Default = count
	}

	return ps.Generate(Default).String()
}
