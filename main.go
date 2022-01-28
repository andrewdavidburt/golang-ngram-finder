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

func preprocess(s string) []string {
	// replace newlines with spaces as separators
	mid := strings.ReplaceAll(s, "\n", " ")

	// regex replace to keep only letters, numbers, and spaces
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
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

func breakup(complete []string, chunkSize int, ngramLength int) [][]string {
	var partials [][]string
	for i := 0; i < len(complete); i += chunkSize {
		end := i + chunkSize + ngramLength - 1 // retention of ngrams spanning chunks
		if end > len(complete) {
			end = len(complete)
		}
		partials = append(partials, complete[i:end])
	}
	return partials
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
func ngramFinder(words []string, size int) (allgrams map[string]int) {
	allgrams = make(map[string]int)
	offset := size / 2
	max := len(words)
	for i := range words {
		if i < offset || i+size-offset > max { //  don't run ngram finder where it will run off the beginning or end of the collection
			continue
		}
		gram := strings.Join(words[i-offset:i+size-offset], " ") // collect ngram from words in collection of n/size length (to either side of counter)
		allgrams[gram]++                                         // increment map counter for given ngram
	}
	return allgrams
}

func mergeMaps(maps ...map[string]int) map[string]int {
	result := make(map[string]int)
	for _, m := range maps {
		for k, v := range m {
			if _, ok := result[k]; ok {
				result[k] = result[k] + v
			} else {
				result[k] = v
			}
		}
	}
	return result
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
	words := preprocess(string(incoming))

	type kv struct {
		Key   string
		Value int
	}

	////////////////////////////////////

	// breaks up the file into bite-sized pieces
	var ngbroken []map[string]int
	var ngcomplete map[string]int

	var ss3 []kv

	partials := breakup(words, 1024, 3)

	// finds ngrams in each separate piece
	for _, p := range partials {
		partgram := ngramFinder(p, 3)
		ngbroken = append(ngbroken, partgram)

	}

	// merges them back together
	ngcomplete = mergeMaps(ngbroken...)

	for k, v := range ngcomplete {
		ss3 = append(ss3, kv{k, v})
	}

	sort.Slice(ss3, func(i, j int) bool {
		return ss3[i].Value > ss3[j].Value
	})

	for i := 0; i < 100; i++ {
		fmt.Printf("%d:   %s - %d\n", i+1, ss3[i].Key, ss3[i].Value)
	}

	////////////////////////////////////

	// pre-processed data is sent to look for three-word sequences/trigrams
	ng := ngramFinder(words, 3)

	// since maps in golang are inherently unordered, they cannot be sorted. therefore, an index of some sort is required, such as this slice of key-values

	var ss []kv
	for k, v := range ng {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	fmt.Println("Rank: 3-Word Sequence - Count")
	fmt.Println("____________________________")

	// top 100 results are output in sorted order. if fewer than 100 results are present, however many are available are output in sorted order.
	if len(ss) >= 100 {
		for i := 0; i < 100; i++ {
			fmt.Printf("%d:   %s - %d\n", i+1, ss[i].Key, ss[i].Value)
		}
	} else {
		for i := 0; i < len(ss); i++ {
			fmt.Printf("%d:   %s - %d\n", i+1, ss[i].Key, ss[i].Value)
		}
	}

}
