package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	sum := sumup(22, 10, 9, 3, 8)
	anotherSum := sumup(22, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val // sum = sum + val
	}
	return sum
}