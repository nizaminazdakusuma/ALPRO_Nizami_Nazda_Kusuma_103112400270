package main

import (
    "fmt"
)

func main() {
    var nilai int

    fmt.Print("Masukkan nilai ujian siswa: ")
    fmt.Scan(&nilai)

    if nilai >= 85 {
        fmt.Println("Lulus")
    } else {
        fmt.Println("Tidak Lulus")
    }
}