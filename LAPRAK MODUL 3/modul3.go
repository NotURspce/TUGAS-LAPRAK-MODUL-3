//TUGAS PRAKTIKUM (TERBIMBING)
package main
import "fmt"

func main() {
	var kecepatan, menit, jam, totaljarak int

	//bagian input
	fmt.Print("Masukan kecepatan rata rata mu: ")
	fmt.Scan(&kecepatan) //membaca inputan nilai dari user
	totaljarak = 100 + 60 + 170 // menjumlahkan total jarak perjalanan
	menit = (totaljarak * 60) / kecepatan // menghitung estimasi total waktu dlm satuan mnt
	jam = menit / 60 // ini untuk konversi menit jadi jam 
	menit = menit % 60 // mencari sisa mnt yg sudh dikonvrsi ke jam menggunakan modulus

	//bagian output
	fmt.Print("Kamu memakan waktu : ")
	fmt.Println(jam, menit)

}