package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)

	whatIsIt = reverseMsg(sd)

	whatIsIt = strings.ReplaceAll(whatIsIt, ":", " ")

	fmt.Printf("%s. Okay!, I accept the invitation.", whatIsIt)
}

func reverseMsg(msg []byte) string {
	runes := []rune(string(msg))
	for start, end := 0, len(msg)-1; start < end; start, end = start+1, end-1 {
		runes[start], runes[end] = runes[end], runes[start]
	}
	return string(runes)
}
