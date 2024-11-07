package main

import "fmt"

func main() {
	age := 32 // regular variable
	// var agePointer *int
	agePointer := &age

	fmt.Println("Age:", *agePointer) // same as age
	fmt.Println("agePointer:", agePointer)
	// adultYears := getAdultYears(age)
	// adultYears := getAdultYears(agePointer)
	// fmt.Println(adultYears)
	aditAgeToAdultYears(agePointer)
	fmt.Println(age)
}

// func getAdultYears(age int) int {
// 	return age - 18
// }
func aditAgeToAdultYears(age *int) {
	*age = *age - 18
}

// PS.: avoid this methods to small values like int, numb, string...
