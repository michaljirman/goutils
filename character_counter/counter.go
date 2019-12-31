package main

import (
	"fmt"
	"sort"
	"unicode"
)


func main()  {
	text := "hello world!123"
	m := make(map[rune]int)
	for _, r := range text {
		if unicode.IsLetter(r){
			m[r]++
		}
	}
	//Answer is now in m.
	// iteration over map is random by design
	//for r, n := range m {
	//	fmt.Printf("%s: %d\n", string(r), n)
	//}

	// Sort by frequency and format output:
	lfs := make(lfList, 0, len(m))
	for l, f := range m {
		lfs = append(lfs, &letterFreq{l, f})
	}
	sort.Sort(lfs)
	for _, lf := range lfs {
		fmt.Printf("   %c (%d)    %7d\n", lf.rune, lf.rune,lf.freq)
	}


	w1 := "bcadA"
	w2 := SortString(w1)

	fmt.Println(w1)
	fmt.Println(w2)
}

type letterFreq struct {
	rune
	freq int
}
type lfList []*letterFreq
func (lfs lfList) Len() int { return len(lfs) }
func (lfs lfList) Less(i, j int) bool {
	return lfs[i].rune < lfs[j].rune
	//switch fd := lfs[i].freq - lfs[j].freq; {
	//case fd < 0:
	//	return false
	//case fd > 0:
	//	return true
	//}
	//return lfs[i].rune < lfs[j].rune
}
func (lfs lfList) Swap(i, j int) {
	lfs[i], lfs[j] = lfs[j], lfs[i]
}




type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}