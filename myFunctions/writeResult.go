package myfunctions

import (
	"fmt"
	"os"
	"strings"
)

func WriteResult(asciiChars map[int]string) ([]string, error) {
	result := []string{}
	lineToWrite := 0
	str := os.Args[1]
	isOn := false
	for i := 0; i < len(str); i++ {
		//Handle non ascii char
		if str[i] < 32 || str[i] > 126 {
			return result, fmt.Errorf("A none ascii char has been found !!")
		}

		//Handle \n
		if str[i] == '\\' && str[i+1] == 'n' {
			if i == 0 {
				result = append(result,	"")
				lineToWrite++
				isOn = true
			} else if len(str) == i+2 || (len(str) > i + 3 && str[i+2] == '\\' && str[i+3] == 'n'){
				result = append(result,	"")
				lineToWrite++
			} else {
				fmt.Println("ok")
				newLineAscii := []string{"","","","","","","",""}
				result = append(result,	newLineAscii...)
				lineToWrite += 8
			}
			i++
			continue
		} else if i == 0 || isOn {
			newLineAscii := []string{"","","","","","","",""}
			result = append(result,	newLineAscii...)
			isOn = false
		}
		
		asciiChar := strings.Split(asciiChars[int(str[i])], "\n")
		for i, line := range asciiChar {
			result[i + lineToWrite] += line
		}
	}
	return result, nil
}
