package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 100000
		nilai64 int64 = int64(nilai32)
		nilai8  int8  = int8(nilai64) // this will go wrong, bcs max value of int8 is 127
	)

	fmt.Println(nilai32, nilai64, nilai8)

	name := "Galih"
	firstChar := name[0] // this will get byte of the first char. (byte is alias of uint8)
	firstAlphabet := string(firstChar)

	fmt.Println(name, firstChar, firstAlphabet)
}
