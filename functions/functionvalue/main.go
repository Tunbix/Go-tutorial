package functionvalue

import "fmt"

type transform func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumber(&numbers, double)
	tripled := transformNumber(&numbers, triple)

	fmt.Println(tripled)
	fmt.Println(doubled)
}

func transformNumber(numbers *[]int, transformFunc transform) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transformFunc(val))
	}

	return dNumbers
}

func double(numbers int) int {
	return numbers * 2
}

func triple(numbers int) int {
	return numbers * 3
}
