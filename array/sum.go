package main

import "fmt"

func Sum(numbers [5]int) int{
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum

}

func main()  {
	num := [5]int{1,2,3,4,5}
	fmt.Println(Sum(num))
}