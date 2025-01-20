package main

func main() {

	var nilaiUas = 80
	var nilaiUts = 80
	var nilaiAbsen = 90

	var lulusNilaiUas = nilaiUas > 75
	var lulusNilaiUts = nilaiUts > 75
	var lulusNilaiAbsen = nilaiAbsen > 75

	var result = lulusNilaiUas && lulusNilaiUts && lulusNilaiAbsen

	println(result)
	// Output
	// true
}