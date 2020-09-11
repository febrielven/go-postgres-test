package main

import (
	"fmt"
)

func main() {
	// var name = [4]string{"febri", "anto", "rezki", "razka"}

	// fmt.Println(" Hello Word")
	// fmt.Println("Isi Elemen \t", name)

	// var list1 = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	// linearsort(list1, 200)
	var list2 = []int{10, 7, 34, 97, 2, 43, 23, 13, 9, 1}
	selectionSort(list2)

	for i := 0; i < len(list2); i++ {
		fmt.Println(list2[i], " ")
	}

	fmt.Print("\n")

}

func linearsort(list []int, key int) bool {
	for _, data := range list {
		if data == key {
			fmt.Println("angka yang anda cari tersedia \t", key)
			return true
		}

	}

	fmt.Println("Anga tidak tersedia")
	return false
}

func selectionSort(list []int) {
	for i := 0; i < len(list); i++ {
		minIndex := i
		for j := i + 1; j < len(list); j++ {
			// fmt.Println(list[j], " ")
			if list[minIndex] > list[j] {
				minIndex = j

				// fmt.Println(minIndex, " ")
			}
		}

		tmp := list[i]

		list[i] = list[minIndex]

		list[minIndex] = tmp
		// fmt.Println(list[minIndex], " ")
	}
}
