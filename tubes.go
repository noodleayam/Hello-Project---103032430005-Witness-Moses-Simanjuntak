package main

import "fmt"

// ini comment

// please ini gua cuma mau tambah comment doang

type DataSampah struct {
	ID      int
	Jenis   string
	Berat   float64
	Lokasi  string
	Tanggal string
}

const NMAX int = 1000

type databaseSampah [NMAX]DataSampah
type cthData [10]DataSampah

var n int = 0
var idCounter int = 1
var cth cthData

func main() {
	var pilihan int
	var db databaseSampah

	for {
		fmt.Println("\n======= Aplikasi Pendataan Sampah Daerah =======")
		fmt.Println("0. Contoh Data Sampah")
		fmt.Println("1. Tampilkan Data Sampah")
		fmt.Println("2. Edit Data Sampah")
		fmt.Println("3. Tambah Data Sampah")
		fmt.Println("4. Urutkan Data")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu ( 0 / 1 / 2 / 3 / 4 / 5 ): ")
		fmt.Scan(&pilihan)
		fmt.Println("===============================================")

		switch pilihan {
		case 0:
			dataContoh()
		case 1:
			menu_tampil_sampah(&db)
		case 2:
			menu_edit_sampah(&db)
		case 3:
			menu_tambah_sampah(&db)
		case 4:
			menu_sorting(&db)
		case 5:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini. Sampai jumpa!")
			fmt.Println("===============================================")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func dataContoh() {
	// Contoh data sampah yang untuk ditambahkan ke database
	var pilihan int
	cth[0] = DataSampah{ID: 3, Jenis: "Organik", Berat: 2.5, Lokasi: "Jakarta_Pusat", Tanggal: "15-01-2025"}
	cth[1] = DataSampah{ID: 1, Jenis: "Anorganik", Berat: 1.2, Lokasi: "Jakarta_Barat", Tanggal: "10-01-2025"}
	cth[2] = DataSampah{ID: 2, Jenis: "B3", Berat: 0.8, Lokasi: "Jakarta_Timur", Tanggal: "20-01-2025"}

	fmt.Println("\n======= Contoh Data Sampah =======")
	fmt.Println("Berikut adalah contoh data sampah yang dapat Anda gunakan sebagai referensi:")
	fmt.Printf("%-5s %-15s %-10s %-20s %-12s\n", "ID", "Jenis", "Berat(kg)", "Lokasi", "Tanggal")
	fmt.Println("----------------------------------------------------------------------")
	for i := 0; i < 3; i++ {
		fmt.Printf("%-5d %-15s %-10.2f %-20s %-12s\n",
			cth[i].ID, cth[i].Jenis, cth[i].Berat, cth[i].Lokasi, cth[i].Tanggal)
	}
	fmt.Println("NOTES: -> Data ini tidak masuk dalam database.")
	fmt.Println("       -> Jadikan data ini sebagai acuan saat menambahkan data ke database.")
	fmt.Println("===================================")

	fmt.Println("Kembali ke menu utama?")
	fmt.Println("(1) Ya")
	fmt.Println("(0) Tidak (tampilkan lagi)")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		fmt.Println("Kembali ke menu utama...")
		return
	case 0:
		dataContoh()
	default:
		fmt.Println("Pilihan tidak valid. Kembali ke menu utama...")
		return
	}
}

func menu_tampil_sampah(db *databaseSampah) {
	var pilihan int
	fmt.Println("\n======= Data Sampah =======")
	if n <= 0 {
		fmt.Println("Belum ada data sampah. Kembali ke menu utama...")
		return
	}
	fmt.Println("Data sampah yang tersedia:")
	fmt.Println("----------------------------------------------------------------------")
	selectionSortByID(db)

	fmt.Printf("%-5s %-15s %-10s %-20s %-12s\n", "ID", "Jenis", "Berat(kg)", "Lokasi", "Tanggal")
	fmt.Println("----------------------------------------------------------------------")

	for i := 0; i < n; i++ {
		fmt.Printf("%-5d %-15s %-10.2f %-20s %-12s\n",
			db[i].ID, db[i].Jenis, db[i].Berat, db[i].Lokasi, db[i].Tanggal)
	}
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("Total data sampah: %d\n", n)
	fmt.Println("===============================================")

	fmt.Println("Kembali ke menu utama?")
	fmt.Println("(1) Ya")
	fmt.Println("(0) Tidak (tampilkan lagi)")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Println("Kembali ke menu utama...")
		return
	case 0:
		menu_tampil_sampah(db)
	default:
		fmt.Println("Pilihan tidak valid. Kembali ke menu utama...")
		return
	}
}

func menu_tambah_sampah(db *databaseSampah) {
	var jumlahTambah int

	fmt.Println("\n======= Tambah Data Sampah =======")
	fmt.Print("Berapa data sampah yang ingin ditambahkan? ")
	fmt.Scan(&jumlahTambah)

	if jumlahTambah <= 0 {
		fmt.Println("Jumlah data yang ingin ditambah harus lebih dari 0. Kembali ke menu utama...")
		return
	}

	if n+jumlahTambah > NMAX {
		fmt.Printf("Tidak dapat menambah %d data. Database hanya bisa menampung %d data lagi.\n",
			jumlahTambah, NMAX-n)
		return
	}

	dataAwal := n

	for i := 0; i < jumlahTambah; i++ {
		if n >= NMAX {
			fmt.Println("Database penuh! Tidak dapat menambah data lagi.")
			break
		}

		var data DataSampah
		data.ID = idCounter

		for {
			found := false
			for j := 0; j < n; j++ {
				if db[j].ID == data.ID {
					found = true
					break
				}
			}
			if !found {
				break
			}
			data.ID++
		}

		fmt.Printf("\n--- Data ke-%d ---\n", i+1)
		fmt.Print("Jenis sampah (Organik / Anorganik / B3): ")
		fmt.Scan(&data.Jenis)
		fmt.Print("Berat (kg): ")
		fmt.Scan(&data.Berat)
		fmt.Print("Lokasi: ")
		fmt.Scan(&data.Lokasi)

		data.Tanggal = inputTanggal()

		db[n] = data
		n++
		idCounter = data.ID + 1

		fmt.Printf("Data sampah berhasil ditambahkan dengan ID: %d\n", data.ID)
	}

	selectionSortByID(db)
	fmt.Printf("\nTotal %d data sampah berhasil ditambahkan!\n", n-dataAwal)
	fmt.Println("===============================================")
}

func menu_edit_sampah(db *databaseSampah) {
	fmt.Println("\n======= Edit Data Sampah =======")
	if n == 0 {
		fmt.Println("Belum ada data sampah untuk diedit. Kembali ke menu utama...")
		fmt.Println("===============================================")
		return
	}

	var id int
	fmt.Print("Masukkan ID data yang ingin diedit: ")
	fmt.Scan(&id)

	index := cariDataByID(*db, id)
	if index == -1 {
		fmt.Println("Data dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Println("\nData saat ini:")
	fmt.Printf("ID: %d, Jenis: %s, Berat: %.2f kg, Lokasi: %s, Tanggal: %s\n",
		db[index].ID, db[index].Jenis, db[index].Berat, db[index].Lokasi, db[index].Tanggal)

	fmt.Print("Jenis sampah baru: ")
	fmt.Scan(&db[index].Jenis)
	fmt.Print("Berat baru (kg): ")
	fmt.Scan(&db[index].Berat)
	fmt.Print("Lokasi baru: ")
	fmt.Scan(&db[index].Lokasi)

	fmt.Print("Tanggal baru: ")
	db[index].Tanggal = inputTanggal()

	selectionSortByID(db)
	fmt.Println("Data berhasil diupdate!")
}

// Fungsi untuk mengurutkan data sampah berdasarkan berat atau tanggal
// Menggunakan Insertion Sort untuk berat dan tanggal
func menu_sorting(db *databaseSampah) {
	if n == 0 {
		fmt.Println("Belum ada data sampah untuk diurutkan.")
		return
	}

	var pilihan int
	fmt.Println("\n======= Menu Pengurutan =======")
	fmt.Println("1. Urutkan berdasarkan Berat")
	fmt.Println("2. Urutkan berdasarkan Tanggal")
	fmt.Print("Pilih jenis pengurutan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		insertionSortByBerat(db)
		fmt.Println("Data berhasil diurutkan berdasarkan Jenis Sampah!")
	case 2:
		insertionSortByTanggal(db)
		fmt.Println("Data berhasil diurutkan berdasarkan Tanggal!")
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// Fungsi untuk mengurutkan data sampah berdasarkan ID menggunakan Selection Sort
// Mengurutkan dari ID terkecil ke terbesar
func selectionSortByID(db *databaseSampah) {
	for i := 1; i < n; i++ {
		minIdx := i - 1
		// Temukan indeks dengan ID terkecil
		for j := i; j < n; j++ {
			if db[j].ID < db[minIdx].ID {
				minIdx = j
			}
		}
		// Tukar data pada indeks i dengan indeks minIdx

		db[i-1], db[minIdx] = db[minIdx], db[i-1]

	}
}

// Fungsi untuk mengurutkan data berdasarkan berat menggunakan Insertion Sort
// Mengurutkan dari berat terkecil ke terbesar
func insertionSortByBerat(db *databaseSampah) {
	for i := 1; i < n; i++ {
		key := db[i]
		j := i - 1

		// Urutkan berdasarkan berat
		for j >= 0 && db[j].Berat > key.Berat {
			db[j+1] = db[j]
			j--
		}
		db[j+1] = key
	}
	fmt.Println("Data berhasil diurutkan berdasarkan Berat!")
	fmt.Println("===============================================")
	fmt.Println("Data sampah yang telah diurutkan:")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("%-5s %-15s %-10s %-20s %-12s\n", "ID", "Jenis", "Berat(kg)", "Lokasi", "Tanggal")
	for i := 0; i < n; i++ {
		fmt.Printf("%-5d %-15s %-10.2f %-20s %-12s\n",
			db[i].ID, db[i].Jenis, db[i].Berat, db[i].Lokasi, db[i].Tanggal)
	}
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("Total data sampah: %d\n", n)
	fmt.Println("===============================================")
	fmt.Println("Kembali ke menu utama?")
	fmt.Println("(1) Ya")
	fmt.Println("(0) Tidak (tampilkan lagi)")
	fmt.Print("Pilih: ")
	var pilihan int
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		fmt.Println("Kembali ke menu utama...")
		return
	case 0:
		menu_tampil_sampah(db)
	default:
		fmt.Println("Pilihan tidak valid. Kembali ke menu utama...")
		return
	}
}

// Fungsi untuk mengkonversi tanggal dari format DD-MM-YYYY ke YYYY-MM-DD
// Diperlukan untuk sorting tanggal secara leksikografis
func konversiTanggalUntukSort(tanggal string) string {
	// Input: DD-MM-YYYY
	// Output: YYYY-MM-DD
	if len(tanggal) != 10 {
		return tanggal
	}

	dd := tanggal[0:2]
	mm := tanggal[3:5]
	yyyy := tanggal[6:10]

	return yyyy + "-" + mm + "-" + dd
}

// Fungsi untuk mengurutkan data berdasarkan tanggal menggunakan Insertion Sort
// Menggunakan konversi tanggal untuk perbandingan yang benar
func insertionSortByTanggal(db *databaseSampah) {
	for i := 1; i < n; i++ {
		key := db[i]
		j := i - 1

		tanggalKey := konversiTanggalUntukSort(key.Tanggal)
		for j >= 0 && konversiTanggalUntukSort(db[j].Tanggal) > tanggalKey {
			db[j+1] = db[j]
			j--
		}
		db[j+1] = key
	}

	fmt.Println("Data berhasil diurutkan berdasarkan Tanggal!")
	fmt.Println("===============================================")
	fmt.Println("Data sampah yang telah diurutkan:")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("%-5s %-15s %-10s %-20s %-12s\n", "ID", "Jenis", "Berat(kg)", "Lokasi", "Tanggal")
	for i := 0; i < n; i++ {
		fmt.Printf("%-5d %-15s %-10.2f %-20s %-12s\n",
			db[i].ID, db[i].Jenis, db[i].Berat, db[i].Lokasi, db[i].Tanggal)
	}
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("Total data sampah: %d\n", n)
	fmt.Println("===============================================")

	fmt.Println("Kembali ke menu utama?")
	fmt.Println("(1) Ya")
	fmt.Println("(0) Tidak (tampilkan lagi)")
	fmt.Print("Pilih: ")
	var pilihan int
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		fmt.Println("Kembali ke menu utama...")
		return
	case 0:
		menu_tampil_sampah(db)
	default:
		fmt.Println("Pilihan tidak valid. Kembali ke menu utama...")
		return
	}
}

func validasiTanggal(tanggal string) bool {
	// Cek panjang string
	if len(tanggal) != 10 {
		return false
	}

	// Cek format DD-MM-YYYY (posisi tanda '-')
	if tanggal[2] != '-' || tanggal[5] != '-' {
		return false
	}

	// Cek apakah DD, MM, YYYY adalah angka
	dd := tanggal[0:2]
	mm := tanggal[3:5]
	yyyy := tanggal[6:10]

	if len(dd) != 2 || len(mm) != 2 || len(yyyy) != 4 {
		return false
	}

	// Cek apakah semua karakter adalah digit
	for _, char := range dd + mm + yyyy {
		if char < '0' || char > '9' {
			return false
		}
	}

	return true
}

func inputTanggal() string {
	var tanggal string
	for {
		fmt.Print("Tanggal (DD-MM-YYYY): ")
		fmt.Scan(&tanggal)

		if validasiTanggal(tanggal) {
			return tanggal
		} else {
			fmt.Println("Format tanggal salah! Gunakan format DD-MM-YYYY (contoh: 15-01-2025)")
		}
	}
}

// Fungsi untuk mencari data sampah berdasarkan ID menggunakan Binary Search
// Data harus sudah diurutkan berdasarkan ID sebelum melakukan pencarian
func cariDataByID(db databaseSampah, id int) int {
	selectionSortByID(&db)

	left := 0
	right := n - 1

	for left <= right {
		mid := (left + right) / 2
		if db[mid].ID == id {
			return mid
		}
		if db[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
