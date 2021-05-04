package main

import "fmt"

/*
	- is like object in javascript
	- Map is data container that we can defince data type of the index
	- Map is data container with key-value. which key is unique
	- We can put a lot of data inside map in condition the key is unique
	- If we put data with the same value, it will replace recent value of that key
*/

func main() {
	// create map map[key-type]value-type{data}
	person := map[string]string{
		"name":    "Galih",
		"Address": "Yogya",
	}

	person["title"] = "engineer" // add new data to map or replace current data if key is exist

	fmt.Println(person)
	fmt.Println(person["name"]) // get value for some key
	fmt.Println(len(person))    // get length of map

}
