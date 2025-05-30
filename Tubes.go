package main

import (
    "fmt"
    "strings"
)

const NMAX = 100


type Crypto struct {
    ID        int
    Name      string
    Symbol    string
    Price     float64
    MarketCap float64
}

type User struct {
    Username string
    Password string
}

type OwnedCrypto struct {
    CryptoID int
    Amount   float64
}


var cryptoList [NMAX]Crypto
var cryptoCount int

var userList = [NMAX]User{
    {"Gustavo", "1234"},
    {"Alice", "pass"},
}
var userCount = 2

var userHoldings [NMAX][NMAX]OwnedCrypto
var holdingsCount [NMAX]int
var currentUserIndex int

func main() {
    var choice int

    fmt.Println("=== Selamat Datang di Aplikasi Simulasi Perdagangan Kripto ===")
    fmt.Println("1. Start")
    fmt.Println("2. Keluar")
    fmt.Print("Pilih menu: ")
    fmt.Scan(&choice)

    if choice == 2 {
        fmt.Println("Terima kasih!")
        return
    } else if choice != 1 {
        fmt.Println("Pilihan tidak valid!")
        return
    }

    if !loginUser() {
        return
    }

    initSampleData()

    for {
        fmt.Println("\n=== Menu Utama ===")
        fmt.Println("1. Tampilkan Daftar Crypto")
        fmt.Println("2. Tambah Crypto Baru")
        fmt.Println("3. Edit Data Crypto")
        fmt.Println("4. Hapus Data Crypto")
        fmt.Println("5. Cari Crypto dengan Harga > $10 (Sequential Search)") // <<-- ganti teks menu
        fmt.Println("6. Cari Crypto (Binary Search)")
        fmt.Println("7. Lihat Crypto Saya")
        fmt.Println("8. Keluar")
        fmt.Print("Pilih menu: ")
        fmt.Scan(&choice)

        switch choice {
        case 1:
            displayCryptoMenu()
        case 2:
            addCrypto()
        case 3:
            editCrypto()
        case 4:
            deleteCrypto()
        case 5:
            searchByPriceAbove10() // <<-- panggil fungsi baru di sini
        case 6:
            searchBinary()
        case 7:
            viewMyCrypto()
        case 8:
            fmt.Println("Terima kasih!")
            return
        default:
            fmt.Println("Pilihan tidak valid!")
        }
    }
}

func loginUser() bool {
    var uname, pass string
    fmt.Print("Masukkan username: ")
    fmt.Scan(&uname)
    fmt.Print("Masukkan password: ")
    fmt.Scan(&pass)

    for i := 0; i < userCount; i++ {
        if userList[i].Username == uname && userList[i].Password == pass {
            currentUserIndex = i
            fmt.Printf("=== Selamat datang, %s ===\n", uname)
            return true
        }
    }

    fmt.Println("Username atau password salah!")
    return false
}


func initSampleData() {
    cryptoList[0] = Crypto{1, "Bitcoin", "BTC", 65000, 1200000000}
    cryptoList[1] = Crypto{2, "Ethereum", "ETH", 3200, 380000000}
    cryptoList[2] = Crypto{3, "Dogecoin", "DOGE", 0.15, 20000000}
    cryptoCount = 3
}

func addCrypto() {
    if cryptoCount >= NMAX {
        fmt.Println("Database penuh!")
        return
    }

    var name, symbol string
    var price, marketCap float64

    fmt.Print("Masukkan nama crypto: ")
    fmt.Scan(&name)
    fmt.Print("Masukkan simbol: ")
    fmt.Scan(&symbol)
    fmt.Print("Masukkan harga: ")
    fmt.Scan(&price)
    fmt.Print("Masukkan market cap: ")
    fmt.Scan(&marketCap)

    cryptoCount++
    cryptoList[cryptoCount-1] = Crypto{cryptoCount, name, strings.ToUpper(symbol), price, marketCap}
    fmt.Println("Data berhasil ditambahkan!")
}

func editCrypto() {
    var keyword string
    fmt.Print("Masukkan nama crypto yang ingin diedit: ")
    fmt.Scan(&keyword)

    idx := sequentialSearch(keyword)
    if idx == -1 {
        fmt.Println("Crypto tidak ditemukan!")
        return
    }

    fmt.Print("Masukkan harga baru: ")
    fmt.Scan(&cryptoList[idx].Price)
    fmt.Print("Masukkan market cap baru: ")
    fmt.Scan(&cryptoList[idx].MarketCap)

    fmt.Println("Data berhasil diperbarui!")
}

func deleteCrypto() {
    var keyword string
    fmt.Print("Masukkan nama crypto yang ingin dihapus: ")
    fmt.Scan(&keyword)

    idx := sequentialSearch(keyword)
    if idx == -1 {
        fmt.Println("Crypto tidak ditemukan!")
        return
    }

    for i := idx; i < cryptoCount-1; i++ {
        cryptoList[i] = cryptoList[i+1]
    }
    cryptoCount--

    fmt.Println("Data berhasil dihapus!")
}

func sequentialSearch(keyword string) int {
    for i := 0; i < cryptoCount; i++ {
        if strings.EqualFold(cryptoList[i].Name, keyword) {
            return i
        }
    }
    return -1
}

func searchSequential() {
    var keyword string
    fmt.Print("Masukkan nama crypto: ")
    fmt.Scan(&keyword)

    idx := sequentialSearch(keyword)
    if idx == -1 {
        fmt.Println("Crypto tidak ditemukan!")
    } else {
        printCrypto(cryptoList[idx])
    }
}

func binarySearch(keyword string) int {
    left := 0
    right := cryptoCount - 1

    for left <= right {
        mid := (left + right) / 2
        if strings.EqualFold(cryptoList[mid].Name, keyword) {
            return mid
        } else if strings.ToLower(keyword) < strings.ToLower(cryptoList[mid].Name) {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return -1
}

func searchBinary() {
    insertionSortByName(true)

    var keyword string
    fmt.Print("Masukkan nama crypto: ")
    fmt.Scan(&keyword)

    idx := binarySearch(keyword)
    if idx == -1 {
        fmt.Println("Crypto tidak ditemukan!")
    } else {
        printCrypto(cryptoList[idx])
    }
}

func displayCryptoMenu() {
    var choice, asc int

    fmt.Println("\nUrutkan berdasarkan:")
    fmt.Println("1. Nama (Insertion Sort)")
    fmt.Println("2. Harga (Selection Sort)")
    fmt.Print("Pilih: ")
    fmt.Scan(&choice)

    fmt.Print("Urutan (1: ASC, 2: DESC): ")
    fmt.Scan(&asc)

    isAsc := asc == 1

    if choice == 1 {
        insertionSortByName(isAsc)
    } else if choice == 2 {
        selectionSortByPrice(isAsc)
    } else {
        fmt.Println("Pilihan salah!")
        return
    }

    displayCryptoList()
}

func displayCryptoList() {
    fmt.Println("\nID | Nama | Simbol | Harga | MarketCap")
    for i := 0; i < cryptoCount; i++ {
        printCrypto(cryptoList[i])
    }
}

func printCrypto(c Crypto) {
    fmt.Printf("%d | %s | %s | $%.2f | $%.2f\n", c.ID, c.Name, c.Symbol, c.Price, c.MarketCap)
}

func selectionSortByPrice(ascending bool) {
    for i := 0; i < cryptoCount-1; i++ {
        idx := i
        for j := i + 1; j < cryptoCount; j++ {
            if ascending && cryptoList[j].Price < cryptoList[idx].Price || !ascending && cryptoList[j].Price > cryptoList[idx].Price {
                idx = j
            }
        }
        cryptoList[i], cryptoList[idx] = cryptoList[idx], cryptoList[i]
    }
}

func insertionSortByName(ascending bool) {
    for i := 1; i < cryptoCount; i++ {
        temp := cryptoList[i]
        j := i - 1
        for j >= 0 && ((ascending && strings.ToLower(cryptoList[j].Name) > strings.ToLower(temp.Name)) || (!ascending && strings.ToLower(cryptoList[j].Name) < strings.ToLower(temp.Name))) {
            cryptoList[j+1] = cryptoList[j]
            j--
        }
        cryptoList[j+1] = temp
    }
}

func viewMyCrypto() {
    for {
        fmt.Printf("\n=== Menu Crypto %s ===\n", userList[currentUserIndex].Username)
        fmt.Println("1. Lihat Kepemilikan Crypto")
        fmt.Println("2. Tambah Koin Baru")
        fmt.Println("3. Edit Koin")
        fmt.Println("4. Hapus Koin")
        fmt.Println("5. Transaksi Koin (Beli/Jual)")
        fmt.Println("6. Kembali ke Menu Utama")
        fmt.Print("Pilih menu: ")

        var pilihan int
        fmt.Scan(&pilihan)

        switch pilihan {
        case 1:
            if holdingsCount[currentUserIndex] == 0 {
                fmt.Println("Anda belum memiliki crypto.")
            } else {
                fmt.Println("Nama | Simbol | Jumlah | Total ($)")
                for i := 0; i < holdingsCount[currentUserIndex]; i++ {
                    oc := userHoldings[currentUserIndex][i]
                    cIdx := findCryptoByID(oc.CryptoID)
                    if cIdx != -1 {
                        c := cryptoList[cIdx]
                        total := oc.Amount * c.Price
                        fmt.Printf("%s | %s | %.4f | %.2f\n", c.Name, c.Symbol, oc.Amount, total)
                    }
                }
            }
        case 2:
            addCryptoToUser()
        case 3:
            editCryptoUser()
        case 4:
            deleteCryptoUser()
        case 5:
            simulateTrade()
        case 6:
            return
        default:
            fmt.Println("Pilihan tidak valid!")
        }
    }
}

func findCryptoByID(id int) int {
    for i := 0; i < cryptoCount; i++ {
        if cryptoList[i].ID == id {
            return i
        }
    }
    return -1
}

func addCryptoToUser() {
    if cryptoCount == 0 {
        fmt.Println("Belum ada crypto yang tersedia.")
        return
    }
    var keyword string
    fmt.Print("Masukkan nama crypto yang ingin ditambah: ")
    fmt.Scan(&keyword)
    idx := sequentialSearch(keyword)
    if idx == -1 {
        fmt.Println("Crypto tidak ditemukan.")
        return
    }
    var amount float64
    fmt.Print("Masukkan jumlah yang ingin ditambahkan: ")
    fmt.Scan(&amount)
    if amount <= 0 {
        fmt.Println("Jumlah harus positif.")
        return
    }
    
    for i := 0; i < holdingsCount[currentUserIndex]; i++ {
        if userHoldings[currentUserIndex][i].CryptoID == cryptoList[idx].ID {
            userHoldings[currentUserIndex][i].Amount += amount
            fmt.Println("Jumlah crypto berhasil ditambahkan.")
            return
        }
    }
    
    userHoldings[currentUserIndex][holdingsCount[currentUserIndex]] = OwnedCrypto{
        CryptoID: cryptoList[idx].ID,
        Amount:   amount,
    }
    holdingsCount[currentUserIndex]++
    fmt.Println("Crypto berhasil ditambahkan ke kepemilikan Anda.")
}

func editCryptoUser() {
    if holdingsCount[currentUserIndex] == 0 {
        fmt.Println("Anda belum memiliki crypto.")
        return
    }
    var keyword string
    fmt.Print("Masukkan nama crypto yang ingin diedit jumlahnya: ")
    fmt.Scan(&keyword)
    cIdx := sequentialSearch(keyword)
    if cIdx == -1 {
        fmt.Println("Crypto tidak ditemukan.")
        return
    }
    
    for i := 0; i < holdingsCount[currentUserIndex]; i++ {
        if userHoldings[currentUserIndex][i].CryptoID == cryptoList[cIdx].ID {
            var newAmount float64
            fmt.Print("Masukkan jumlah baru: ")
            fmt.Scan(&newAmount)
            if newAmount < 0 {
                fmt.Println("Jumlah tidak boleh negatif.")
                return
            }
            userHoldings[currentUserIndex][i].Amount = newAmount
            fmt.Println("Jumlah berhasil diubah.")
            return
        }
    }
    fmt.Println("Anda tidak memiliki crypto ini.")
}

func deleteCryptoUser() {
    if holdingsCount[currentUserIndex] == 0 {
        fmt.Println("Anda belum memiliki crypto.")
        return
    }
    var keyword string
    fmt.Print("Masukkan nama crypto yang ingin dihapus dari kepemilikan: ")
    fmt.Scan(&keyword)
    cIdx := sequentialSearch(keyword)
    if cIdx == -1 {
        fmt.Println("Crypto tidak ditemukan.")
        return
    }
    
    for i := 0; i < holdingsCount[currentUserIndex]; i++ {
        if userHoldings[currentUserIndex][i].CryptoID == cryptoList[cIdx].ID {
            for j := i; j < holdingsCount[currentUserIndex]-1; j++ {
                userHoldings[currentUserIndex][j] = userHoldings[currentUserIndex][j+1]
            }
            holdingsCount[currentUserIndex]--
            fmt.Println("Crypto berhasil dihapus dari kepemilikan Anda.")
            return
        }
    }
    fmt.Println("Anda tidak memiliki crypto ini.")
}

func simulateTrade() {
    if cryptoCount == 0 {
        fmt.Println("Belum ada crypto yang tersedia.")
        return
    }
    var keyword string
    fmt.Print("Masukkan nama crypto yang ingin ditransaksikan: ")
    fmt.Scan(&keyword)
    cIdx := sequentialSearch(keyword)
    if cIdx == -1 {
        fmt.Println("Crypto tidak ditemukan.")
        return
    }
    var tradeType int
    fmt.Println("1. Beli")
    fmt.Println("2. Jual")
    fmt.Print("Pilih transaksi: ")
    fmt.Scan(&tradeType)

    var amount float64
    fmt.Print("Masukkan jumlah: ")
    fmt.Scan(&amount)
    if amount <= 0 {
        fmt.Println("Jumlah harus positif.")
        return
    }

    switch tradeType {
    case 1: 
        for i := 0; i < holdingsCount[currentUserIndex]; i++ {
            if userHoldings[currentUserIndex][i].CryptoID == cryptoList[cIdx].ID {
                userHoldings[currentUserIndex][i].Amount += amount
                fmt.Println("Transaksi beli berhasil.")
                return
            }
        }
        
        userHoldings[currentUserIndex][holdingsCount[currentUserIndex]] = OwnedCrypto{
            CryptoID: cryptoList[cIdx].ID,
            Amount:   amount,
        }
        holdingsCount[currentUserIndex]++
        fmt.Println("Transaksi beli berhasil.")
    case 2: 
        for i := 0; i < holdingsCount[currentUserIndex]; i++ {
            if userHoldings[currentUserIndex][i].CryptoID == cryptoList[cIdx].ID {
                if userHoldings[currentUserIndex][i].Amount < amount {
                    fmt.Println("Jumlah jual melebihi kepemilikan.")
                    return
                }
                userHoldings[currentUserIndex][i].Amount -= amount
                fmt.Println("Transaksi jual berhasil.")
                if userHoldings[currentUserIndex][i].Amount == 0 {
                    
                    for j := i; j < holdingsCount[currentUserIndex]-1; j++ {
                        userHoldings[currentUserIndex][j] = userHoldings[currentUserIndex][j+1]
                    }
                    holdingsCount[currentUserIndex]--
                }
                return
            }
        }
        fmt.Println("Anda tidak memiliki crypto ini.")
    default:
        fmt.Println("Pilihan transaksi salah.")
    }
}

func searchByPriceAbove10() {
    fmt.Println("Crypto dengan harga > $10:")
    found := false
    for i := 0; i < cryptoCount; i++ {
        if cryptoList[i].Price > 10 {
            printCrypto(cryptoList[i])
            found = true
        }
    }
    if !found {
        fmt.Println("Tidak ditemukan crypto dengan harga di atas $10.")
    }
}