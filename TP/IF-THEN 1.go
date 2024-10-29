package main

import (
    "fmt"
)

func main() {
    var nilai int

    // Meminta pengguna memasukkan nilai ujian
    fmt.Print("Masukkan nilai ujian siswa: ")
    fmt.Scan(&nilai)

    // Menentukan kelulusan
    if nilai >= 85 {
        fmt.Println("Lulus")
    } else {
        fmt.Println("Tidak Lulus")
    }
}