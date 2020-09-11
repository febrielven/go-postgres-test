package main

import "fmt"

func main() {
	var in int = 10
	var out int = 2

	lama := hitung(in, out)
	fmt.Println("lama bekerja = \n", lama)
}

func hitung(in int, out int) int {

	var lama int
	if out >= in {
		lama = out - in
	} else {
		lama = (12 - in) + out
	}

	return lama

}
