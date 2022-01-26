package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

func Preprocess(s string) []string {

	mid := strings.ReplaceAll(s, "\n", " ")
	// mid = strings.ReplaceAll(s, "-", " ")
	// mid = strings.ReplaceAll(s, "—", " ")

	// mid = strings.ReplaceAll(mid, "-", "")
	// mid = strings.ReplaceAll(mid, "—", "")
	// mid = strings.ReplaceAll(mid, "_", "")
	// mid = strings.ReplaceAll(mid, "'", "")
	// mid = strings.ReplaceAll(mid, ",", "")
	// mid = strings.ReplaceAll(mid, ".", "")
	// mid = strings.ReplaceAll(mid, ";", "")
	// mid = strings.ReplaceAll(mid, "?", "")

	reg, err := regexp.Compile("[^a-zA-Z0-9' ]+")
	if err != nil {
		log.Fatal(err)
	}
	result := reg.ReplaceAllString(mid, "")

	output := strings.ToLower(result)
	// fmt.Println(output)

	notALetter := func(char rune) bool {
		return unicode.IsSpace(char)
	}
	return strings.FieldsFunc(output, notALetter)
}

func main() {
	var incoming string

	if len(os.Args) <= 1 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		incoming = string(scanner.Bytes())
	} else {
		for _, file := range os.Args[1:] {
			incoming += string(openFile(file))
		}
	}

	words := Preprocess(string(incoming))

	fmt.Println(words)

	ng := ngrams(words, 3)

	type kv struct {
		Key   string
		Value uint32
	}

	var ss []kv
	for k, v := range ng {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	fmt.Println("1: the sperm whale, 85")
	fmt.Println("2: the white whale, 71")
	fmt.Println("3: of the whale, 67")
	fmt.Println("------------------------")
	fmt.Println("Rank: 3-Word Sequence, Count")
	fmt.Println("____________________________")
	for i := 0; i <= 9; i++ {
		fmt.Printf("%d: %s, %d\n", i+1, ss[i].Key, ss[i].Value)
	}
	fmt.Println("_____")

	// for k, v := range ng {
	// 	fmt.Printf("\"%v:\": %v, ", k, v)
	// }

}
func openFile(filename string) []byte {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicf("couldn't read file: %s", err)
	}
	return data
}

func ngrams(words []string, size int) (count map[string]uint32) {

	count = make(map[string]uint32)
	offset := int(float64(size / 2))

	max := len(words)
	for i := range words {
		if i < offset || i+size-offset > max {
			continue
		}
		gram := strings.Join(words[i-offset:i+size-offset], " ")
		count[gram]++
	}

	return count
}
