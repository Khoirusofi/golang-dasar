package main

func main() {
	days := []string{"asa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	switch days[0] {
	case "senin":
		println("Hari ini Senin")
	case "asa":
		println("Hari ini Selasa")
	default:
		println("Hari ini Bukan Senin")
	}

	switch length := len(days); length < 5 {
	case true:
		println("Panjang banget")
	case false:
		println("Dikit banget")
	}
}