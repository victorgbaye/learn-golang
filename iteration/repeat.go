package main

import "fmt"

func Repeat(char string) string{
	var repeated string
	for i := 0; i < 5; i++{
		repeated += char
	}
	return repeated
}

func main(){
	fmt.Println(Repeat("a"))
}