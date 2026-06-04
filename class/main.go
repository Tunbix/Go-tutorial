package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)
	//userNames := []string{}

	userNames[0] = "Alice"

	userNames = append(userNames, "Bob")
	userNames = append(userNames, "Charlie")

	fmt.Println(userNames)

	courseRatings := floatMap{}

	courseRatings["Go Programming"] = 4.5
	courseRatings["Python"] = 4.7
	courseRatings["Java"] = 4.3

	courseRatings.output()

	//fmt.Println(courseRatings) 

	for course, rating := range courseRatings {
		fmt.Println("Course:", course)
		fmt.Println("Rating:", rating)
	}

	for index, name := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Name:", name)
	}
}