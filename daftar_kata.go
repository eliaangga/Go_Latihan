package main

import "fmt"

type tabStr [99]string

func main() {
	var t tabStr
	var teks, w string
	var katakata int
	var post, idn int
	fmt.Scan(&teks, &w)
	MengisiiArray(teks, &t, &katakata)
	MencariPosisi(t, katakata+1, w, &post, &idn)
	if post == -1 && idn == -1 {
		fmt.Println("-1")
	} else {
		fmt.Println(post+1, idn)
	}
}

func MengisiiArray(s string, daftar *tabStr, k *int) {
	var katakata string
	katakata = ""
	for i := 0; i < len(s); i++ {
		if s[i] == '_' {
			katakata = ""
			*k += 1
		} else {
			katakata = katakata + string(s[i])
			daftar[*k] = katakata
		}
	}
}

func MencariIndex(w string, k string) int {
	var x int
	x = len(w)
	for i := 0; i < len(k); i++ {
		if len(k) >= i+x {
			if w == k[i:i+x] {
				return i
			}
		}
	}
	return -1
}

func MencariPosisi(daftar tabStr, n int, w string, posisi, index *int) {
	var idn int
	*posisi = -1
	*index = -1
	for i := 0; i < n; i++ {
		idn = MencariIndex(w, daftar[i])
		if idn != -1 {
			*posisi = i
			*index = idn
			return
		}
	}
}
