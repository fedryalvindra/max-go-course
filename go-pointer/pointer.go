package main

import "fmt"

func main() {
	age := 32

	var agePointer *int
	// & to get address value
	agePointer = &age

	// * to get value behind pointer
	fmt.Println("age", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
