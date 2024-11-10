package main

import "fmt"

func main() {
	var n, m int
	var bagi, sisa int
	fmt.Print("Masukkan Pembilang : ")
	fmt.Scan(&n)
	fmt.Print("Masukkan penyebut : ")
	fmt.Scan(&m)

	bagi = n / m
	sisa = n % m

	fmt.Println("Hasil dari pembagian adalah : ", bagi)
	fmt.Println("Sisa dari hasil bagi adalah : ", sisa)
}
