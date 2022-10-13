package main

import "fmt"

func main() {
	var isim string
	var boy float32
	var kilo float32

	fmt.Println("Lütfen isiminizi giriniz ↓↓")
	fmt.Scan(&isim)
	fmt.Println("")
	fmt.Println("Doğru bir sonuç için boyunuzu metre cinsinden giriniz ↓↓")
	fmt.Scan(&boy)
	fmt.Println("")
	fmt.Println("Son olarak kilonuzu giriniz ↓↓")
	fmt.Scan(&kilo)

	//var vucutKitleEndeksi float32

	// for i := 0; i < len(isimler); i++ {

	//var almanGerekenKilo float32
	//var vermenGerekenKilo float32
	//var saglikliKilo float32

	vucutKitleEndeksi := vucutKitleEndeksdeğeri(isim, kilo, boy)

	if vucutKitleEndeksi >= 18 && vucutKitleEndeksi <= 25 {

		fmt.Println(isim, "kitle endeksi normal ", vucutKitleEndeksi)
		fmt.Println("")
		fmt.Println("normal = 18-25")
	} else if vucutKitleEndeksi < 18 {
		saglikliKilo := 18 * (boy * boy)

		almanGerekenKilo := saglikliKilo - kilo
		fmt.Println(isim, "alinacak kilo : +", almanGerekenKilo)
		fmt.Println("vücut kitle endeksi", vucutKitleEndeksi)
	} else if vucutKitleEndeksi > 25 {
		saglikliKilo := 25 * (boy * boy)

		vermenGerekenKilo := kilo - saglikliKilo
		fmt.Println(isim, "vermen gereken kilo : -", vermenGerekenKilo)
		fmt.Println("vücut kitle endeksi", vucutKitleEndeksi)
	}
	fmt.Println("")
}

func vucutKitleEndeksdeğeri(isimler string, kilo float32, boy float32) (vucutKitleEndeksdeğeri float32) {
	vucutKitleEndeksi := kilo / (boy * boy)
	return vucutKitleEndeksi
}
