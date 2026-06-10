package anonymous

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}

	double := createTransformFunc(2)
	triple := createTransformFunc(3)

	transformed := transformNumber(&numbers, func(num int) int {
		return num * 2
	})

	doubled := transformNumber(&numbers, double)
	tripled := transformNumber(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumber(numbers *[]int, transformFunc func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transformFunc(val))
	}

	return dNumbers
}

func createTransformFunc(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}