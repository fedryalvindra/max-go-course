package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output(){
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	courseRatings.output()

	for index, username := range userNames {
		// ...
		fmt.Println(index, username)
	}

	for key, course := range courseRatings {
		fmt.Println(key, course)
	}
}