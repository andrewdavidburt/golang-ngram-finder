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

	reg, err := regexp.Compile("[^a-zA-Z0-9' ]+")
	if err != nil {
		log.Fatal(err)
	}
	result := reg.ReplaceAllString(mid, "")

	output := strings.ToLower(result)

	isSpace := func(char rune) bool {
		return unicode.IsSpace(char)
	}
	return strings.FieldsFunc(output, isSpace)
}

func main() {
	var incoming string

	if len(os.Args) <= 1 {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				incoming += scanner.Text()
			}
		} else {
			fmt.Println("This program counts 3-word sequences (trigrams) in a document, and outputs the top 100 in order. ")
			fmt.Println("Please either specify one or more text files as arguments after the program on the command-line, or pipe text in via stdin.")
			os.Exit(0)
		}
	} else {
		for _, file := range os.Args[1:] {
			incoming += string(openFile(file))
		}
	}

	words := Preprocess(string(incoming))

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

	fmt.Println("Rank: 3-Word Sequence, Count")
	fmt.Println("____________________________")

	if len(ss) >= 100 {
		for i := 0; i < 100; i++ {
			fmt.Printf("%d: %s, %d\n", i+1, ss[i].Key, ss[i].Value)
		}
	} else {
		for i := 0; i < len(ss); i++ {
			fmt.Printf("%d: %s, %d\n", i+1, ss[i].Key, ss[i].Value)
		}
	}
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
