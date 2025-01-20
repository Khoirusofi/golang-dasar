package main

func main() {
	var nilai32 int32 = 32767
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	println(nilai32, nilai64, nilai16)

	// Output
	// 32767 32767 32767

	var name = "M Khoirusofi"
	var e byte = name[5]
	var eString = string(e)
	println(name, e, eString)

	// Output
	// M Khoirusofi 105 i

}