package main

import "fmt"

const NMAX int = 1000

type waktu struct {
    menit, detik int
}

type lagu struct {
    title, penyanyi string
    durasi          waktu
}

type playList struct {
    lagu [NMAX]lagu
}

var joedoellagu, penyanyi string
var m, d int

func main() {
    var listLagu playList
    var idx = 0
    MengisiPlaylist(&listLagu, &idx)
    OutDariPlaylist(listLagu, idx)
}

func MencariLagu(jdl string, penyani string, T playList, n int) bool {
    for i := 0; i < n; i++ {
        if T.lagu[i].title == jdl && T.lagu[i].penyanyi == penyani {
            return true
        }
    }
    return false
}

func MengisiPlaylist(T *playList, n *int) {
    fmt.Scan(&joedoellagu, &penyanyi)
    for joedoellagu != "#" || penyanyi != "#" {
        fmt.Scan(&m, &d)
        if !MencariLagu(joedoellagu, penyanyi, *T, *n) {
            T.lagu[*n].title = joedoellagu
            T.lagu[*n].penyanyi = penyanyi
            T.lagu[*n].durasi.menit = m
            T.lagu[*n].durasi.detik = d
            *n += 1
        }
        fmt.Scan(&joedoellagu, &penyanyi)
    }
}

func MencariDurasi(T playList, n int) lagu {
    var detik int
    var terlama lagu
    detik = 0
    for i := 0; i < n; i++ {
        if (T.lagu[i].durasi.menit)*60+T.lagu[i].durasi.detik > detik {
            detik = (T.lagu[i].durasi.menit)*60 + T.lagu[i].durasi.detik
            terlama = T.lagu[i]
        }
    }
    return terlama
}

func OutDariPlaylist(T playList, n int) {
    var lama = MencariDurasi(T, n)
    for i := 0; i < n; i++ {
        if T.lagu[i] == lama {
            fmt.Println("*", T.lagu[i].title, T.lagu[i].durasi.menit, "Menit", T.lagu[i].durasi.detik, "Detik")
        } else {
            fmt.Println(T.lagu[i].title)
        }
    }
}
