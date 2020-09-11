package main

import "fmt"

func main() {

	var in int = 10
	var out int = 2
	log, total := calculate(in, out)
	fmt.Println("biaya pakir lama: ", log, " dengan biaya: ", total)
}

func calculate(in int, out int) (int, int) {
	var log int
	var total int

	if out >= in {
		log = out - in
	} else {
		log = (12 - in) + out
	}

	if log > 2 {
		total = 2000 + (log-2)*500
	} else {
		total = log * 2000
	}

	return log, total

}
