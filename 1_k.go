// 2311102310
package main

import "fmt"

func main() {
	var n, i int
	var notiket float64
	fmt.Println("Masukkan jumlah peserta :")
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
		fmt.Println("Masukkan nomor tiket peserta :", i)
		fmt.Scan(&notiket)

		/*fungsi penentuan hadiah
		  if semua digit pada nomor tiketnya adalah angka yang sama dan genap = Hadiah utama
		  if semua digit pada nomor tiketnya adalah angka ganjil = Hadiah sembakko
		  else = Hadiah konsol*/

		/*if notiket == 55466 && notiket == 66787 {
			fmt.Println("Hadiah konsol")
		} else {
			fmt.Println("Hadiah utama")
		}*/
	}
}
