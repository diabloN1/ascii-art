package myfunctions

import (
	"fmt"
	"os"
	"strings"
)

func WriteResult(asciiChars map[byte]string) ([]string, error) {
	result := []string{}
	lineToWrite := 0
	str := os.Args[1]
	inWord := false
	letterCount := 0
	for i := 0; i < len(str); i++ {
		// Handle non ascii char.
		if str[i] < 32 || str[i] > 126 {
			return result, fmt.Errorf("A none ascii char has been found !!")
		}

		// Handle \n.
		if i+1 < len(str) && str[i] == '\\' && str[i+1] == 'n' {
			if !inWord {
				result = append(result, "")
				lineToWrite++
			} else if OnlyNewLinesRemaining(str[i:]) {
				result = append(result, "")
				lineToWrite++
			}
			inWord = false
			i++
			continue
		}
		
		// Prepare Slice to write character.
		if letterCount == 0 {
			newLineAscii := []string{"", "", "", "", "", "", "", ""}
			result = append(result, newLineAscii...)
		} else if !inWord {
			newLineAscii := []string{"", "", "", "", "", "", "", ""}
			result = append(result, newLineAscii...)
			lineToWrite += 8
		}

		//Filling the letter in the result slice.
		letterCount++
		inWord = true
		asciiChar := strings.Split(asciiChars[str[i]], "\n")
		for i, line := range asciiChar {
			result[i+lineToWrite] += line
		}
	}
	return result, nil
}

func OnlyNewLinesRemaining(str string) bool {
	for i := 0; i < len(str); i++ {
		if i+1 < len(str) && str[i] == '\\' && str[i+1] == 'n' {
			i++
			continue
		} else {
			return false
		}
	}
	return true
}