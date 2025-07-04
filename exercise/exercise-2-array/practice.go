package main

import "fmt"

func main() {

	// 1).
	hobbies := [3]string{"Football", "Boxing", "Run"}
	for _, hobby := range hobbies {
		fmt.Println(hobby)
	}

	// 2).
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// 3).
	newSlices := hobbies[:2]
	fmt.Println(newSlices)

	// 4).
	reSlice := newSlices[1:3]
	fmt.Println(reSlice)

	// 5).
	goals := []string{"Great Developer", "Billionare"}
	fmt.Println(goals)

	// 6).
	goals[1] = "Miliarder"
	goals = append(goals, "Boxer")
	fmt.Println(goals)

	// 7).
	type Product struct {
		id    int
		title string
		price float64
	}

	products := []Product{{
		id:    1,
		title: "Book",
		price: 100,
	}, {
		id:    2,
		title: "Pencil",
		price: 150,
	}}

	products[1] = Product{
		id:    2,
		title: "Eraser",
		price: 100,
	}

	products = append(products, Product{id: 3, title: "Pen", price: 1000})

	fmt.Println(products)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
