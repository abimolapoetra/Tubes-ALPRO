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