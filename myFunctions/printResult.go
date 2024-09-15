package myfunctions

import "fmt"

func PrintResult(result []string) {
	for _, v := range result {
		fmt.Println(v)
	}
}
