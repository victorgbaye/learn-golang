package main

// import "fmt"

func Sum(numbers []int) int{
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum

}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
// func main()  {
// 	num := []int{1,2,3,4,5}
// 	fmt.Println(Sum(num))
// }