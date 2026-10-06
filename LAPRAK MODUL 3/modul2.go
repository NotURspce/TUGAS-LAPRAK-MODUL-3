//TUGAS PRAKTIKUM (PENDAHULUAN)
package main
import "fmt"

func main() {
	var n, hari int

	//bagian input
	fmt.Print("Masukan angka : ")
	fmt.Scan(&n) //membaca masukan angka dari user dan di simpan di var n
	hari = (4 + n - 1) % 7 + 1 // bagian rumus dan menghitung

	//bagian output
	fmt.Print("ini angka nya : ")
	fmt.Println(hari) // ini untuk memunculkan hasil perhitungan (output)
	
}