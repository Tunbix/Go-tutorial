package main

import "fmt"

type Product struct {
	id string
	name string
	price float64
}

func main() {
	//1)
	var hobbies [3]string = [3]string{"coding", "gaming", "traveling"}
	fmt.Println(hobbies)

	//2)
	println(hobbies[0])
	println(hobbies[1:3])

	//3)
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	//4)
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	//5)
	courseGoals := []string{"Learn Go", "Build Projects", "Contribute to Open Source"}
	fmt.Println(courseGoals)

	//6)
	courseGoals[1] = "Master Go!"
	courseGoals = append(courseGoals, "Become proficient in Go!")
	fmt.Println(courseGoals)

	//7)
	products := []Product{
		{
			"first-product",
			"A first product",
			19.99,
		},

		{
			"second-product",
			"A second product",
			129.99,
		},
	}
	fmt.Println(products)

	newProduct := Product{
	 "third-product",
	 "A third product",
	  349.99,
	}
	products = append(products, newProduct)
	fmt.Println(products)
}