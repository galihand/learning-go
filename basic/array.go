package main

import "fmt"

func main() {
	// array declaration
	var name [3]string
	name[0] = "budi"
	name[1] = "andi"
	name[2] = "joko"

	var umur = [3]int{
		1, 2, 3,
	}

	fmt.Println(name)
	fmt.Println(umur)
	fmt.Println("Panjang Array = ", len(name))
}
