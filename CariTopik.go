package main

import "fmt"

type tag struct {
	topik  string
	banyak int
}

type tabTopic [99]string

type tabTag [99]tag

func main() {
	var h_tag tabTag
	var tag int
	tweetTag := tabTopic{
		"libur",
		"makan",
		"lebaran",
		"pricon",
		"hari_raya",
		"libur",
		"libur",
		"libur",
		"libur",
		"libur",
		"libur",
		"makan",
		"lebaran",
		"pricon",
		"libur",
		"makan",
		"libur",
		"makan",
		"lebaran",
		"pricon",
		"hari_raya",
		"lebaran",
		"pricon",
		"hari_raya",
	}
	MengisiiTag(&h_tag, &tag, tweetTag, 22)
	fmt.Println("Topik Trending di Twitter yaitu:", MencariTrendingTopik(h_tag, tag))
}

func MencariTopik(T tabTag, topik string, n int) int {
	for i := 0; i < n; i++ {
		if T[i].topik == topik {
			return i
		}
	}
	return -1
}

func MengisiiTag(T *tabTag, n_tag *int, hash tabTopic, n_hash int) {
	var idx int
	for i := 0; i < n_hash; i++ {
		idx = MencariTopik(*T, hash[i], *n_tag)
		if idx != -1 {
			T[idx].banyak += 1
		} else {
			T[*n_tag].topik = hash[i]
			T[*n_tag].banyak = 1
			*n_tag += 1
		}
	}
}

func MencariTrendingTopik(T tabTag, n int) string {
	var banyak, id int
	banyak = 0
	for i := 0; i < n; i++ {
		if T[i].banyak > banyak {
			banyak = T[i].banyak
			id = i
		}
	}
	return T[id].topik
}
