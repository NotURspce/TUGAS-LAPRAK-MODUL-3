//TUGAS PRAKTIKUM (MANDIRI)
package main 
import "fmt"

func main () {
	var gp, lembur, bl, pt, gb int

	//BAGIAN INPUT GAJI POKOK
	fmt.Print("Masukan Gaji pokokmu : ")
	fmt.Scan(&gp) // membaca input dari user berupa gaji pokok

	//BAGIAN INPUT LEMBUR
	fmt.Print("Masukan Berapa jam lembur anda : ")
	fmt.Scan(&lembur)
	
	//BAGIAN PERHITUNGAN
	bl = lembur * 45000 //menghitung bonus lembur =  lama waktu x 45k /jam
	//karena tipe dta int hasil perkalian desimal di konv secara otomatis kebawah
	pt = int(float64(gp) * 0.055) // perhitungan potongan , 5.5% (2%+3.5%) dari gajipokok
	gb = gp + bl - pt // menghitung gaji bersih setelah menghitung bonus dan pengurangan potongan
	
	//BAGIAN OUTPUT
	fmt.Print("Gaji bersihmu adalah : ")
	fmt.Println(gb) //memunculkan hasil gaji bersih di terminal

}