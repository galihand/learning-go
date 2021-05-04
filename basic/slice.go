package main

import "fmt"

/*
	- Slice just like array but more flexible
	- Slice is connected with array.
	-	When you create slice it's also create array in the background
	- Slice has has:
		- Pointer = pointed to array data
		- Length = length of slice. cant greater than capacity
		- Capacity = capacity of slice
	- Becareful when create array or slice
		- Array [...] / [5]
		- Slice []
*/

func main() {
	var months = [...]string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sept", "Oct", "Nov", "Dec",
	}

	// create slice from array
	var slice1 = months[3:7] // slice months index 3 to 6

	// create slice

	fmt.Println(slice1)
	fmt.Println(len(slice1)) // length of slice1
	fmt.Println(cap(slice1)) // capacity of slice1

	months[3] = "April" // reassigned array will effect the slice
	fmt.Println(slice1)

	slice1[0] = "Apr" // reassigned slice will effect the array
	fmt.Println(months)

	// will append element in end of slice or replace the next value array after slice
	// But if capacity is doesn't enough append will generate new array
	// So append doesnt give an impact either to the original slice or original array
	slice2 := months[10:]
	slice3 := append(slice2, "something")
	slice3[1] = "December"

	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	// create slice from the beginning
	newSlice := make([]string, 2, 5) //make(tipe, length, capacity)
	newSlice[0] = "zero"
	newSlice[1] = "one"
	otherSlice := []string{
		"other1", "other2",
	}

	fmt.Println(otherSlice, len(otherSlice), cap(otherSlice))
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice
	// make sure length and capacity is equal or greater than src slice
	coplyOfNewSlice := make([]string, len(newSlice), cap(newSlice))
	copy(coplyOfNewSlice, newSlice)

	fmt.Println(coplyOfNewSlice)

}
