package main

import "fmt"

func main() {
	var name string
	name = "Galih"

	// other option for declare and intialize variable
	var isMale = true
	umur := 26 // ":" just for initialize. for reassigned doesn't need ":"

	// declaration multiple variable
	var (
		origin   = "Yogyakarta"
		religion = "islam"
	)
	isMaried, fullname := true, "Galih Andyan Anindita"

	// const value is not changeable and doesn't triggered error if wasn't used
	// const is not a variable but we can do multiple declaration as variable
	const address = "Yogyakarta, Indonesia"

	fmt.Println(name, umur, isMale, origin, religion, isMaried, fullname)
	fmt.Println(address)
}
