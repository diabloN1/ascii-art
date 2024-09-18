package main

import (
	"asciiArt/myFunctions"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [STRING]")
		fmt.Println()
		fmt.Println("EX: go run . something")
		return
	}
	standard, err := myfunctions.Read("standard.txt")
	if err != nil {
		return
	}
	asciiChars := myfunctions.BytesToAsciiMap(standard)
	result, err := myfunctions.WriteResult(asciiChars)
	if err != nil {
		log.Println(err)
		return
	}
	myfunctions.PrintResult(result)
}
