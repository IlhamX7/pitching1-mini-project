package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Buku struct {
	No             int
	Judul, Penulis string
	Rating         int
}

func create(daftarBuku []Buku) []Buku {
	fmt.Println("")
	fmt.Println("Silahkan tambah buku favorit")
	fmt.Println("")

	var judul, penulis string
	var rating int

	fmt.Print("Masukkan judul buku baru: ")
	judul = strings.TrimSpace(judul)
	reader := bufio.NewReader(os.Stdin)
	judul, _ = reader.ReadString('\n')
	judul = strings.TrimSpace(judul)
	fmt.Scanln(&judul)

	fmt.Print("Masukkan nama penulis buku baru: ")
	penulis = strings.TrimSpace(penulis)
	reader1 := bufio.NewReader(os.Stdin)
	penulis, _ = reader1.ReadString('\n')
	penulis = strings.TrimSpace(penulis)
	fmt.Scanln(&penulis)

	fmt.Print("Masukkan nilai rating buku baru (1-5): ")
	fmt.Scanln(&rating)

	maxValue := daftarBuku[0].No

	for _, buku := range daftarBuku {
		if buku.No > maxValue {
			maxValue = buku.No
		}
	}

	fmt.Println()
	newBuku := Buku{
		No:      maxValue + 1,
		Judul:   judul,
		Penulis: penulis,
		Rating:  rating,
	}
	daftarBuku = append(daftarBuku, newBuku)

	fmt.Println("Success add new book")
	fmt.Println(daftarBuku)
	fmt.Println("")

	var next int
	fmt.Println("1. Lanjut pilih menu")
	fmt.Println("2. Keluar")
	fmt.Print("Selanjutnya pilih no yg di inginkan: ")
	fmt.Scan(&next)
	if next == 1 {
		welcome(daftarBuku)
	} else if next == 2 {
		return daftarBuku
	} else {
		fmt.Println("Maaf nomer yang anda pilih tidak tersedia")
	}
	return daftarBuku
}

func read(daftarBuku []Buku) []Buku {
	var searchJudul string
	fmt.Println("")
	fmt.Print("Masukkan judul buku yang ingin dicari: ")
	searchJudul = strings.TrimSpace(searchJudul)
	reader := bufio.NewReader(os.Stdin)
	searchJudul, _ = reader.ReadString('\n')
	searchJudul = strings.TrimSpace(searchJudul)
	fmt.Scanln(&searchJudul)
	fmt.Println("")

	var found bool
	for _, buku := range daftarBuku {
		if strings.Contains(buku.Judul, searchJudul) {
			fmt.Printf("Buku ditemukan:\nNo: %d\nJudul: %s\nPenulis: %s\nRating: %d\n", buku.No, buku.Judul, buku.Penulis, buku.Rating)
			found = true
			break
		}
	}
	fmt.Println("")
	var next int
	fmt.Println("1. Lanjut pilih menu")
	fmt.Println("2. Keluar")
	fmt.Print("Selanjutnya pilih no yg di inginkan: ")
	fmt.Scan(&next)
	if next == 1 {
		welcome(daftarBuku)
	} else if next == 2 {
		return daftarBuku
	} else {
		fmt.Println("Maaf nomer yang anda pilih tidak tersedia")
	}

	if !found {
		fmt.Println("Buku tidak ditemukan.")
		fmt.Println("")
		var next int
		fmt.Println("1. Lanjut pilih menu")
		fmt.Println("2. Keluar")
		fmt.Print("Selanjutnya pilih no yg di inginkan: ")
		fmt.Scan(&next)
		if next == 1 {
			welcome(daftarBuku)
		} else if next == 2 {
			return daftarBuku
		} else {
			fmt.Println("Maaf nomer yang anda pilih tidak tersedia")
		}
	}
	return daftarBuku
}

func getAll(daftarBuku []Buku) []Buku {
	fmt.Println("")
	fmt.Println("Berikut list buku yang tersedia:")
	fmt.Println("")
	for _, buku := range daftarBuku {
		fmt.Printf("No: %d\nJudul: %s\nPenulis: %s\nRating: %d\n\n", buku.No, buku.Judul, buku.Penulis, buku.Rating)
	}
	var next int
	fmt.Println("1. Lanjut pilih menu")
	fmt.Println("2. Keluar")
	fmt.Print("Selanjutnya pilih no yg di inginkan: ")
	fmt.Scan(&next)
	if next == 1 {
		welcome(daftarBuku)
	} else if next == 2 {
		return daftarBuku
	} else {
		fmt.Println("Maaf nomer yang anda pilih tidak tersedia")
	}
	return daftarBuku
}

func update(daftarBuku []Buku) []Buku {
	var updateNo int
	fmt.Println("")
	fmt.Print("Masukkan no buku yang ingin diupdate: ")
	fmt.Scanln(&updateNo)
	fmt.Println("")

	var found bool
	var foundIndex int = -1
	for i, buku := range daftarBuku {
		if buku.No == updateNo {
			found = true
			foundIndex = i
			break
		}
	}

	if !found {
		fmt.Println("Buku tidak ditemukan.")
		fmt.Println("")
		var next int
		fmt.Println("1. Lanjut pilih menu")
		fmt.Println("2. Keluar")
		fmt.Print("Selanjutnya pilih no yg di inginkan: ")
		fmt.Scan(&next)
		if next == 1 {
			welcome(daftarBuku)
		} else if next == 2 {
			return daftarBuku
		} else {
			fmt.Println("Maaf nomer yang anda pilih tidak tersedia")
		}
	}

	var newJudul, newPenulis string
	var newRating int

	fmt.Print("")
	fmt.Print("Masukkan judul baru: ")
	newJudul = strings.TrimSpace(newJudul)
	reader := bufio.NewReader(os.Stdin)
	newJudul, _ = reader.ReadString('\n')
	newJudul = strings.TrimSpace(newJudul)
	fmt.Scanln(&newJudul)
	newJudul = strings.TrimSpace(newJudul)

	fmt.Print("Masukkan nama penulis baru: ")
	newPenulis = strings.TrimSpace(newPenulis)
	reader1 := bufio.NewReader(os.Stdin)
	newPenulis, _ = reader1.ReadString('\n')
	newPenulis = strings.TrimSpace(newPenulis)
	fmt.Scanln(&newPenulis)

	fmt.Print("Masukkan nilai rating baru (1-5): ")
	fmt.Scanln(&newRating)

	daftarBuku[foundIndex].Judul = newJudul
	daftarBuku[foundIndex].Penulis = newPenulis
	daftarBuku[foundIndex].Rating = newRating

	fmt.Println("")
	fmt.Println("Success Update List of Books")
	fmt.Println("")
	fmt.Println("List of Books:")
	for i, buku := range daftarBuku {
		fmt.Printf("Buku %d:\nNo: %d\nJudul: %s\nPenulis: %s\nRating: %d\n\n", i+1, buku.No, buku.Judul, buku.Penulis, buku.Rating)
	}

	var next int
	fmt.Println("1. Lanjut pilih menu")
	fmt.Println("2. Keluar")
	fmt.Print("Selanjutnya pilih no yg di inginkan: ")
	fmt.Scan(&next)
	if next == 1 {
		welcome(daftarBuku)
	} else if next == 2 {
		return daftarBuku
	} else {
		fmt.Println("Maaf nomer yang anda pilih tidak tersedia")
	}

	return daftarBuku
}

func delete(daftarBuku []Buku) []Buku {
	var deleteNo int
	fmt.Println("")
	fmt.Print("Masukkan no buku yang ingin dihapus: ")
	fmt.Scanln(&deleteNo)
	fmt.Println("")

	var found bool
	var foundIndex int = -1
	for i, buku := range daftarBuku {
		if buku.No == deleteNo {
			found = true
			foundIndex = i
			break
		}
	}

	if !found {
		fmt.Println("Buku tidak ditemukan.")
		fmt.Println("")
		var next int
		fmt.Println("1. Lanjut pilih menu")
		fmt.Println("2. Keluar")
		fmt.Print("Selanjutnya pilih no yg di inginkan: ")
		fmt.Scan(&next)
		if next == 1 {
			welcome(daftarBuku)
		} else if next == 2 {
			return daftarBuku
		} else {
			fmt.Println("Maaf nomer yang anda pilih tidak tersedia")
		}
	}

	daftarBuku = append(daftarBuku[:foundIndex], daftarBuku[foundIndex+1:]...)

	fmt.Println("Success Delete One Data from List of Books")
	fmt.Println("")
	fmt.Println("List of Books:")
	for i, buku := range daftarBuku {
		fmt.Printf("Buku %d:\nNo: %d\nJudul: %s\nPenulis: %s\nRating: %d\n\n", i+1, buku.No, buku.Judul, buku.Penulis, buku.Rating)
	}

	var next int
	fmt.Println("1. Lanjut pilih menu")
	fmt.Println("2. Keluar")
	fmt.Print("Selanjutnya pilih no yg di inginkan: ")
	fmt.Scan(&next)
	if next == 1 {
		welcome(daftarBuku)
	} else if next == 2 {
		return daftarBuku
	} else {
		fmt.Println("Maaf nomer yang anda pilih tidak tersedia")
	}
	return daftarBuku
}

func pilihan(no int, daftarBuku []Buku) int {

	fmt.Scan(&no)
	if no == 1 {
		create(daftarBuku)
	} else if no == 2 {
		read(daftarBuku)
	} else if no == 3 {
		getAll(daftarBuku)
	} else if no == 4 {
		update(daftarBuku)
	} else if no == 5 {
		delete(daftarBuku)
	} else {
		fmt.Println("")
		fmt.Println("Maaf nomer yang anda pilih tidak tersedia")
		return no
	}
	return no
}

func welcome(data []Buku) {
	fmt.Println("")
	fmt.Println("Selamat datang di toko buku")
	fmt.Println("")
	fmt.Println("Berikut pilihan menu yang tersedia:")
	fmt.Println("1. Tambah buku favorit")
	fmt.Println("2. Cari buku favorit")
	fmt.Println("3. Lihat semua buku favorit")
	fmt.Println("4. Ubah buku favorit")
	fmt.Println("5. Hapus buku favorit")
	fmt.Println("")

	var no int
	fmt.Print("Silahkan pilih no pilihan yang anda inginkan: ")
	pilihan(no, data)
}

func main() {
	var daftarBuku []Buku

	buku1 := Buku{
		No:      1,
		Judul:   "Harry Potter",
		Penulis: "J. K. Rowling",
		Rating:  5,
	}
	buku2 := Buku{
		No:      2,
		Judul:   "Laskar Pelangi",
		Penulis: "Andrea Hirata",
		Rating:  4,
	}
	buku3 := Buku{
		No:      3,
		Judul:   "The Chronicles of Narnia",
		Penulis: "C.S. Lewis",
		Rating:  5,
	}
	buku4 := Buku{
		No:      4,
		Judul:   "Ayat-Ayat Cinta",
		Penulis: "Habiburrahman El Shirazy",
		Rating:  4,
	}
	daftarBuku = append(daftarBuku, buku1, buku2, buku3, buku4)

	welcome(daftarBuku)
}
