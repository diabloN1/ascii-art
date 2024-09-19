package myfunctions

import (
	"strings"
)

func BytesToAsciiMap(style []byte) map[byte]string {
	chars := make(map[byte]string)
	intValue := 32
	splitedStyle := strings.Split(string(style[1:len(style)-1]), "\n\n")
	// Range char by char standard to fill each ascii separatly in the map.
	for _, v := range splitedStyle {
		chars[byte(intValue)] = v
		intValue++
	}
	return chars
}
