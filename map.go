package main

import "fmt"

func main() {
	days := map[string]string{
		"1": "Senin",
		"2": "Selasa",
		"3": "Rabu",
		"4": "Kamis",
		"5": "Jumat",
		"6": "Sabtu",
		"7": "Minggu",
	}
	
	fmt.Println(days["1"]) // Senin
	fmt.Println(days) // map[1:Senin 2:Selasa 3:Rabu 4:Kamis 5:Jumat 6:Sabtu 7:Minggu

	mounth := map[string]string{
		"1":  "Januari",
    "2":  "Februari",
    "3":  "Maret",
    "4":  "April",
    "5":  "Mei",
    "6":  "Juni",
    "7":  "Juli",
    "8":  "Agustus",
    "9":  "September",
    "10": "Oktober",
    "11": "November",
    "12": "Desember",
  }
	
	fmt.Println(mounth["1"]) // Januari
	fmt.Println(mounth) // map[1:Januari 10:Oktober 11:November 12:Desember 2:Februari 3:Maret 4:April 5:Mei 6:Juni 7:Juli 8:Agustus 9:September]

	delete(mounth, "2")
	fmt.Println(mounth) // map[1:Januari 10:Oktober 11:November 12:Desember 3:Maret 4:April 5:Mei 6:Juni 7:Juli 8:Agustus 9:September]

}