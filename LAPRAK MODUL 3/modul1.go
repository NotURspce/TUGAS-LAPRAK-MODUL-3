//TUGAS PRAKTIKUM (PENDAHULUAN)
package main
import "fmt"

func main() {
	var N, Jam, Sisa, Menit, Detik int

	//bagian input
	fmt.Print("Masukan detik nya : ")
	fmt.Scan(&N) //ini tuh fungsi nya buat membaca angka masukan dari pengguna terus di simpen di var n
	Jam = N / 3600 // untuk menghitung jumlah jam (1jam = 3600dtk)
	Sisa = N % 3600 // menghitung sisa detik setelah di ambil jamnya menggunakan oprasi modulo 
	Menit = Sisa / 60 //menghitung sisa menit dari var sisa (1 mnt = 60dtk)
	Detik = Sisa % 60 // menghitung sisa detik akhir setelah di ambil oleh menit

	//bagian output
	fmt.Print("Hasil konversi : ")
	fmt.Println(Jam, Menit, Detik) // menampilkan hasil konversi (output)
}