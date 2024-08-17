package main

import (
	"fmt"
	"os"
)

// Struct Siswa
type Siswa struct {
	Id        string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Variabel Dummy Nama Siswa
var siswas = []Siswa{
	{"1", "Kahfi", "Jl. Jakarta No. 1", "Programmer", "Tuntutan Pekerjaan"},
	{"1", "Muzadi", "Jl. Jakarta No. 2", "Data Scientist", "Ingin Tambah Pintar"},
	{"3", "Yusuf", "Jl. Jakarta No. 3", "Web Developer", "Biar Makin Keren"},
	{"4", "Kahfi", "Jl. Jakarta No. 4", "Programmer", "Kepengen Aja"},
	{"5", "Reza", "Jl. Jakarta No. 5", "Programmer", "Coba-Coba"},
	{"6", "Riza", "Jl. Jakarta No. 6", "QA", "Ikut Teman"},
}

// Function Untuk Print Nama Siswa
func printNamaSiswa(siswa Siswa) {
	fmt.Printf("ID : %s\n", siswa.Id)
	fmt.Printf("Nama : %s\n", siswa.Nama)
	fmt.Printf("Alamat : %s\n", siswa.Alamat)
	fmt.Printf("Pekerjaan : %s\n", siswa.Pekerjaan)
	fmt.Printf("Alasan Memilih Kelas Golang : %s\n", siswa.Alasan)
}

func main() {
	// Code Untuk Cek Apakah Terminal Ada Parameter Nya Atau Tidak
	if len(os.Args) < 2 {
		fmt.Println("Anda Harus Memasukkan Parameter Terlebih Dahulu")
		return
	}

	// Variabel Untuk Simpan Parameter Dari Terminal
	namaSiswa := os.Args[1]

	// Variabel Untuk Simpan Data Yg Ditemukan Di Terminal
	var hasilPencarian []Siswa

	// Looping Untuk Cek Apakah Ada Data Siswa Yg Cocok Dari Parameter dan Variabel Siswa
	for _, siswa := range siswas {
		if siswa.Nama == namaSiswa || siswa.Id == namaSiswa {
			// Append data yang ditemukan ke slice baru
			hasilPencarian = append(hasilPencarian, siswa)
		}
	}

	// Print Data
	if len(hasilPencarian) > 0 {
		fmt.Println("Data yang ditemukan:")
		fmt.Println()

		for _, siswa := range hasilPencarian {
			printNamaSiswa(siswa)
			fmt.Println()
		}
	} else {
		fmt.Printf("Data Dengan Parameter %s Tidak Ditemukan\n", namaSiswa)
	}
}
