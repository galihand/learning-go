package main

import "fmt"

/*
	- type is a alias for data type
	- basicly we can create alias for some data type
	- so we can understand more easily
*/

func main() {
	type NoKTP string // declare NoKTP as alias of string
	type Married bool // declare Married as alias of bool

	var (
		ktpGalih  NoKTP   = "12345195342352"
		isMarried Married = true
	)

	fmt.Println(ktpGalih, isMarried)
}
