package main

import "fmt"

func main() {
	days := []string{
		"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu",
	}

	daySlice := days[0:]
	fmt.Println(daySlice[:])
	daySlice[0] = "(Senin Libur)"
	fmt.Println(daySlice)

	daySlice2 := days[0:5]
	daySlice2[4] = "(Jumat Libur)"
	fmt.Println(daySlice2)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Senin"
	newSlice[1] = "Selasa"

	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println(newSlice)

	newSlice2 := append(newSlice, "Rabu")
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))
	fmt.Println(newSlice2)

}
