package main

import (
	"log"
	"os"
	"asciiArt/myFunctions"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Usage: go run . [STRING]")
		log.Println()
		log.Println("EX: go run . something")
		return
	}
	standard, err := myfunctions.Read("standard.txt")
	if err != nil {
		return
	}
	asciiChars := myfunctions.BytesToAsciiMap([]byte(standard))
	result, err := myfunctions.WriteResult(asciiChars)
	if err != nil {
		log.Println(err)
		return
	}
	myfunctions.PrintResult(result)
}
