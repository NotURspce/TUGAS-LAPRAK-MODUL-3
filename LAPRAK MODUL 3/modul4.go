//TUGAS PRAKTIKUM (MANDIRI)
package main
import "fmt"

func main() {
	var emas, sisa, perak, tembaga, koin int

	// BAGIAN OUTPUT
	fmt.Print("Masukan koin kamu : ")
	fmt.Scan(&koin) //Membaca total koin user
	emas = koin / 9 //menghitung pertukaran koin emas (emas = 9koin)
	sisa = koin % 9 // menghitung sisa koin setelah petukaran
	perak = sisa / 3 // menghitung pertukaran koin perak dari sisa konversi sebelum nya (perak = 3koin)
	tembaga = koin % 3 //menghitung pertukaran koin tembabaga dari hasil sisa pertukaran perak (tembaga = 1koin)

	//BAGIAN OUTPUT
	fmt.Print("Hasil pertukaranmu : ")
	fmt.Println(emas, perak, tembaga)
}