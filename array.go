package main

func main() {
	var days [3]string
	days[0] = "Senin"
	days[1] = "Selasa"
	days[2] = "Rabu"

	println(days[0]) // Senin
	println(days[1]) // Selasa
	println(days[2]) // Rabu

	var days1 = [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	println(days1[0])   // Senin
	println(len(days1)) // 7

	var months = [12]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
	}

	println(months[11])  // 12
	println(len(months)) // 12
	println(cap(months)) // 12

}