package main

import "fmt"

const NMAX = 100

type RiwayatTidur struct {
    Tanggal     string
    JamTidur    string
    JamBangun   string
    DurasiTidur float64
}

type dataTidur [NMAX]RiwayatTidur

func main() {
    var data dataTidur
    var pilihan int
    for {
        fmt.Println("\n=== APLIKASI PEMANTAUAN KESEHATAN DAN POLA TIDUR ===")
        fmt.Println("1. Tambah Data Tidur")
        fmt.Println("2. Ubah Data Tidur")
        fmt.Println("3. Hapus Data Tidur")
        fmt.Println("4. Cari Data Tidur")
        fmt.Println("5. Urutkan Data Tidur")
        fmt.Println("6. Laporan Pola Tidur")
        fmt.Println("0. Keluar")
        fmt.Print("Pilih menu: ")
        fmt.Scan(&pilihan)
        switch pilihan {
        case 1:
            data = tambahData(data)
        case 2:
            data = ubahData(data)
        case 3:
            data = hapusData(data)
        case 4:
            cariData(data)
        case 5:
            data = urutkanData(data)
        case 6:
            tampilkanLaporan(data)
        case 0:
            fmt.Println("Terima kasih! Selamat istirahat.")
            return
        default:
            fmt.Println("Pilihan tidak valid. Silahkan coba lagi!")
        }
    }
}

func hitungJumlahData(dt dataTidur) int {
    count := 0
    for i := 0; i < NMAX; i++ {
        if dt[i].Tanggal != "" {
            count++
        }
    }
    return count
}

func validasiTanggal(tanggal string) bool {
    var tahun, bulan, hari int
    if _, err := fmt.Sscanf(tanggal, "%d-%d-%d", &tahun, &bulan, &hari); err != nil {
        return false
    }
    if tahun < 2025 || bulan < 1 || bulan > 12 || hari < 1 || hari > 31 {
        return false
    }
    hariBulan := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    if bulan == 2 && ((tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0) {
        if hari > 29 {
            return false
        }
    } else {
        if hari > hariBulan[bulan-1] {
            return false
        }
    }
    return true
}

func validasiJam(jam string) bool {
    var h, m int
    if _, err := fmt.Sscanf(jam, "%d:%d", &h, &m); err != nil {
        return false
    }
    if h < 0 || h > 23 || m < 0 || m > 59 {
        return false
    }
    return true
}

func hitungDurasi(jamTidur, jamBangun string) float64 {
    var jam1, menit1, jam2, menit2 int
    fmt.Sscanf(jamTidur, "%d:%d", &jam1, &menit1)
    fmt.Sscanf(jamBangun, "%d:%d", &jam2, &menit2)
    totalMenitTidur := (jam2*60 + menit2) - (jam1*60 + menit1)
    if totalMenitTidur < 0 {
        totalMenitTidur += 24 * 60
    }
    return float64(totalMenitTidur) / 60.0
}

func tambahData(dt dataTidur) dataTidur {
    jumlahData := hitungJumlahData(dt)
    if jumlahData >= NMAX {
        fmt.Println("Data sudah penuh.")
        return dt
    }
    var tanggal, jamTidur, jamBangun string
    valid := false
    for !valid {
        fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
        fmt.Scan(&tanggal)
        if validasiTanggal(tanggal) {
            valid = true
        } else {
            fmt.Println("Format tanggal tidak valid atau tanggal tidak ada. Coba lagi.")
        }
    }
    valid = false
    for !valid {
        fmt.Print("Masukkan jam tidur (HH:MM): ")
        fmt.Scan(&jamTidur)
        if validasiJam(jamTidur) {
            valid = true
        } else {
            fmt.Println("Format jam tidak valid. Coba lagi.")
        }
    }
    valid = false
    for !valid {
        fmt.Print("Masukkan jam bangun (HH:MM): ")
        fmt.Scan(&jamBangun)
        if validasiJam(jamBangun) {
            valid = true
        } else {
            fmt.Println("Format jam tidak valid. Coba lagi.")
        }
    }
    durasi := hitungDurasi(jamTidur, jamBangun)
    dt[jumlahData] = RiwayatTidur{
        Tanggal:     tanggal,
        JamTidur:    jamTidur,
        JamBangun:   jamBangun,
        DurasiTidur: durasi,
    }
    fmt.Println("Data berhasil ditambahkan!")
    if durasi < 7 {
        fmt.Println("⚠️ Saran: Durasi tidur Anda kurang dari 7 jam. Usahakan untuk tidur lebih lama demi kesehatan.")
    } else if durasi > 9 {
        fmt.Println("⚠️ Saran: Durasi tidur Anda lebih dari 9 jam. Tidur terlalu lama juga kurang baik untuk kesehatan.")
    } else {
        fmt.Println("💬 Feedback: Durasi tidur Anda teratur (7-9 jam). Pertahankan pola tidur ini!")
    }
    var jamTidurJam, jamTidurMenit int
    fmt.Sscanf(jamTidur, "%d:%d", &jamTidurJam, &jamTidurMenit)
    if jamTidurJam >= 23 {
        fmt.Println("⚠️ Saran: Anda tidur di atas jam 11 malam. Usahakan untuk tidur lebih awal agar lebih sehat.")
    } else {
        fmt.Println("💬 Feedback: Anda sudah tidur sebelum jam 11 malam. Pertahankan kebiasaan baik ini!")
    }
    return dt
}

func ubahData(dt dataTidur) dataTidur {
    jumlahData := hitungJumlahData(dt)
    if jumlahData == 0 {
        fmt.Println("Belum ada data untuk diubah.")
        return dt
    }
    var tanggal string
    valid := false
    for !valid {
        fmt.Print("Masukkan tanggal data yang ingin diubah (YYYY-MM-DD): ")
        fmt.Scan(&tanggal)
        if validasiTanggal(tanggal) {
            valid = true
        } else {
            fmt.Println("Format tanggal tidak valid atau tanggal tidak ada. Coba lagi.")
        }
    }
    idx := -1
    for i := 0; i < jumlahData; i++ {
        if dt[i].Tanggal == tanggal {
            idx = i
        }
    }
    if idx == -1 {
        fmt.Println("Data tidak ditemukan.")
        return dt
    }
    fmt.Println("Data ditemukan:")
    fmt.Println("Jam Tidur Sebelumnya:", dt[idx].JamTidur)
    fmt.Println("Jam Bangun Sebelumnya:", dt[idx].JamBangun)
    var jamTidurBaru, jamBangunBaru string
    valid = false
    for !valid {
        fmt.Print("Masukkan jam tidur baru (HH:MM): ")
        fmt.Scan(&jamTidurBaru)
        if validasiJam(jamTidurBaru) {
            valid = true
        } else {
            fmt.Println("Format jam tidak valid. Coba lagi.")
        }
    }
    valid = false
    for !valid {
        fmt.Print("Masukkan jam bangun baru (HH:MM): ")
        fmt.Scan(&jamBangunBaru)
        if validasiJam(jamBangunBaru) {
            valid = true
        } else {
            fmt.Println("Format jam tidak valid. Coba lagi.")
        }
    }
    durasiBaru := hitungDurasi(jamTidurBaru, jamBangunBaru)
    dt[idx].JamTidur = jamTidurBaru
    dt[idx].JamBangun = jamBangunBaru
    dt[idx].DurasiTidur = durasiBaru
    fmt.Println("Data berhasil diperbarui!")
    return dt
}

func hapusData(dt dataTidur) dataTidur {
    jumlahData := hitungJumlahData(dt)
    if jumlahData == 0 {
        fmt.Println("Belum ada data untuk dihapus.")
        return dt
    }
    var tanggal string
    valid := false
    for !valid {
        fmt.Print("Masukkan tanggal data yang ingin dihapus (YYYY-MM-DD): ")
        fmt.Scan(&tanggal)
        if validasiTanggal(tanggal) {
            valid = true
        } else {
            fmt.Println("Format tanggal tidak valid atau tanggal tidak ada. Coba lagi.")
        }
    }
    idx := -1
    for i := 0; i < jumlahData; i++ {
        if dt[i].Tanggal == tanggal {
            idx = i
        }
    }
    if idx == -1 {
        fmt.Println("Data tidak ditemukan.")
        return dt
    }
    var konfirmasi string
    fmt.Printf("Apakah Anda yakin ingin menghapus data tanggal %s? (y/n): ", tanggal)
    fmt.Scan(&konfirmasi)
    if konfirmasi != "y" && konfirmasi != "Y" {
        fmt.Println("Penghapusan dibatalkan.")
        return dt
    }
    for i := idx; i < jumlahData-1; i++ {
        dt[i] = dt[i+1]
    }
    dt[jumlahData-1] = RiwayatTidur{}
    fmt.Println("Data berhasil dihapus!")
    return dt
}

func cariData(dt dataTidur) {
    jumlahData := hitungJumlahData(dt)
    if jumlahData == 0 {
        fmt.Println("Tidak ada data untuk dicari.")
        return
    }
    var tanggal string
    var metode int
    valid := false
    for !valid {
        fmt.Print("Masukkan tanggal yang dicari (YYYY-MM-DD): ")
        fmt.Scan(&tanggal)
        if validasiTanggal(tanggal) {
            valid = true
        } else {
            fmt.Println("Format tanggal tidak valid atau tanggal tidak ada. Coba lagi.")
        }
    }
    fmt.Println("Pilih metode pencarian:")
    fmt.Println("1. Sequential Search")
    fmt.Println("2. Binary Search")
    fmt.Print("Pilihan: ")
    fmt.Scan(&metode)
    var idx int = -1
    if metode == 1 {
        for i := 0; i < jumlahData; i++ {
            if dt[i].Tanggal == tanggal {
                idx = i
            }
        }
    } else if metode == 2 {
        // Pastikan data sudah terurut berdasarkan tanggal ASCENDING sebelum binary search
        dt = selectionSort(dt, 1, true)
        left := 0
        right := jumlahData - 1
        for left <= right && idx == -1 {
            mid := (left + right) / 2
            if dt[mid].Tanggal == tanggal {
                idx = mid
            } else if dt[mid].Tanggal < tanggal {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    } else {
        fmt.Println("Pilihan tidak valid.")
        return
    }
    if idx != -1 {
        fmt.Println("\nData telah ditemukan.")
        fmt.Println("Tanggal     :", dt[idx].Tanggal)
        fmt.Println("Jam Tidur   :", dt[idx].JamTidur)
        fmt.Println("Jam Bangun  :", dt[idx].JamBangun)
        fmt.Printf("Durasi Tidur: %.2f jam\n", dt[idx].DurasiTidur)
    } else {
        fmt.Println("Data tidak ditemukan.")
    }
}

func urutkanData(dt dataTidur) dataTidur {
    jumlahData := hitungJumlahData(dt)
    if jumlahData == 0 {
        fmt.Println("Tidak ada data untuk diurutkan.")
        return dt
    }
    var metode, kriteria, urutan int
    fmt.Println("Pilih metode pengurutan:")
    fmt.Println("1. Selection Sort")
    fmt.Println("2. Insertion Sort")
    fmt.Print("Pilihan: ")
    fmt.Scan(&metode)
    fmt.Println("Urut berdasarkan:")
    fmt.Println("1. Tanggal")
    fmt.Println("2. Durasi Tidur")
    fmt.Print("Pilihan: ")
    fmt.Scan(&kriteria)
    fmt.Println("Urutan:")
    fmt.Println("1. Ascending")
    fmt.Println("2. Descending")
    fmt.Print("Pilihan: ")
    fmt.Scan(&urutan)
    ascending := urutan == 1
    if metode == 1 {
        dt = selectionSort(dt, kriteria, ascending)
    } else if metode == 2 {
        dt = insertionSort(dt, kriteria, ascending)
    } else {
        fmt.Println("Metode tidak valid.")
        return dt
    }
    fmt.Println("Data berhasil diurutkan!")
    return dt
}

func selectionSort(dt dataTidur, kriteria int, ascending bool) dataTidur {
    jumlahData := hitungJumlahData(dt)
    for i := 0; i < jumlahData-1; i++ {
        idxTerpilih := i
        for j := i + 1; j < jumlahData; j++ {
            if banding(dt[j], dt[idxTerpilih], kriteria, ascending) {
                idxTerpilih = j
            }
        }
        dt[i], dt[idxTerpilih] = dt[idxTerpilih], dt[i]
    }
    return dt
}

func insertionSort(dt dataTidur, kriteria int, ascending bool) dataTidur {
    jumlahData := hitungJumlahData(dt)
    for i := 1; i < jumlahData; i++ {
        temp := dt[i]
        j := i - 1
        for j >= 0 && banding(temp, dt[j], kriteria, ascending) {
            dt[j+1] = dt[j]
            j--
        }
        dt[j+1] = temp
    }
    return dt
}

func banding(a, b RiwayatTidur, kriteria int, ascending bool) bool {
    if kriteria == 1 {
        if ascending {
            return a.Tanggal < b.Tanggal
        } else {
            return a.Tanggal > b.Tanggal
        }
    } else {
        if ascending {
            return a.DurasiTidur < b.DurasiTidur
        } else {
            return a.DurasiTidur > b.DurasiTidur
        }
    }
}

func nilaiEkstrim(dt dataTidur, jumlahData int) (int, int) {
    var maxIdx, minIdx int
    start := 0
    if jumlahData > 7 {
        start = jumlahData - 7
    }
    maxIdx, minIdx = start, start
    for i := start + 1; i < jumlahData; i++ {
        if dt[i].DurasiTidur > dt[maxIdx].DurasiTidur {
            maxIdx = i
        }
        if dt[i].DurasiTidur < dt[minIdx].DurasiTidur {
            minIdx = i
        }
    }
    return maxIdx, minIdx
}

func tampilkanLaporan(dt dataTidur) {
    jumlahData := hitungJumlahData(dt)
    if jumlahData == 0 {
        fmt.Println("Tidak ada data untuk ditampilkan.")
        return
    }
    dt = selectionSort(dt, 1, true)
    fmt.Println("===== Laporan Pola Tidur =====")
    fmt.Println("\n📊 Rekap 7 Hari Terakhir:")
    start := 0
    if jumlahData > 7 {
        start = jumlahData - 7
    }
    total7Hari := 0.0
    jumlahHari := 0
    for i := start; i < jumlahData; i++ {
        fmt.Printf("- %s: %.2f jam\n", dt[i].Tanggal, dt[i].DurasiTidur)
        total7Hari += dt[i].DurasiTidur
        jumlahHari++
    }
    if jumlahHari > 0 {
        maxIdx, minIdx := nilaiEkstrim(dt, jumlahData)
        fmt.Printf("\n📈 Durasi tidur terlama: %s (%.2f jam)\n", dt[maxIdx].Tanggal, dt[maxIdx].DurasiTidur)
        fmt.Printf("📉 Durasi tidur tersingkat: %s (%.2f jam)\n", dt[minIdx].Tanggal, dt[minIdx].DurasiTidur)
    }
    totalAll := 0.0
    for i := 0; i < jumlahData; i++ {
        totalAll += dt[i].DurasiTidur
    }
    rataRata := totalAll / float64(jumlahData)
    fmt.Printf("\n Rata-rata durasi tidur dalam %d hari: %.2f jam\n", jumlahData, rataRata)
    if jumlahHari > 0 {
        fmt.Printf(" Rata-rata durasi tidur %d hari terakhir: %.2f jam\n", jumlahHari, total7Hari/float64(jumlahHari))
    }
    fmt.Println()
    if rataRata < 7 {
        fmt.Println("⚠️ Rata-rata durasi tidur kurang dari 7 jam. Perlu perbaikan pola tidur.")
    } else if rataRata > 9 {
        fmt.Println("⚠️ Rata-rata durasi tidur lebih dari 9 jam. Perhatikan pola tidur berlebih.")
    } else {
        fmt.Println("💬 Rata-rata durasi tidur dalam rentang sehat. Pertahankan!")
    }
}