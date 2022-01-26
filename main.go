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
	// replace newlines with spaces as separators
	mid := strings.ReplaceAll(s, "\n", " ")

	// regex replace to keep only letters, numbers, apostrophes, and spaces
	reg, err := regexp.Compile("[^a-zA-Z0-9' ]+")
	if err != nil {
		log.Fatal(err)
	}
	result := reg.ReplaceAllString(mid, "")

	// force all to lower-case
	output := strings.ToLower(result)

	// break string into slice of strings (words) based on space character
	isSpace := func(char rune) bool {
		return unicode.IsSpace(char)
	}
	return strings.FieldsFunc(output, isSpace)
}

func main() {
	var incoming string

	// this checks whether there are command-line arguments. if so, it takes in all files as the corpus to check for trigrams.
	// if not, it checks whether stdin is coming from a pipe or the terminal. if from the terminal, it gives a message describing
	// what the program does and how to use it. if from a pipe (not terminal), it accepts the piped-in file(s) as input to process.
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

	// incoming data is sent for pre-processing
	words := Preprocess(string(incoming))

	// pre-processed data is sent to look for three-word sequences/trigrams
	ng := ngrams(words, 3)

	// since maps in golang are inherently unordered, they cannot be sorted. therefore, an index of some sort is required, such as this slice of key-values
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

	// top 100 results are output in sorted order. if fewer than 100 results are present, however many are available are output in sorted order.
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

// just reads a file
func openFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicf("couldn't read file: %s", err)
	}
	return data
}

// checks for arbitrary-length sequences
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
