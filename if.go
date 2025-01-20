package main

func main() {
	days := []string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	if days[0] == "senin" {
		println("Hari ini Senin")
	} else if days[0] == "selasa" {
		println("Hari ini Selasa")
	} else {
		println("Hari ini Bukan Senin")
	}

	if length := len(days); length > 4 {
		println("Array days memiliki lebih dari 4 elemen")
	} else {
		println("Array days hanya memiliki 4 elemen")
	}
}