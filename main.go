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
	"sync"
	"unicode"
)

type kv struct {
	Key   string
	Value int
}

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

// breaks up file into arbitrarily sized chunks
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

// channeled version of ngram finder
func ngramFinderConcurrent(words []string, size int, ch chan<- map[string]int) {
	allgrams := make(map[string]int)
	offset := size / 2
	max := len(words)
	for i := range words {

		if i < offset || i+size-offset > max { //  don't run ngram finder where it will run off the beginning or end of the collection
			continue
		}
		gram := strings.Join(words[i-offset:i+size-offset], " ") // collect ngram from words in collection of n/size length (to either side of counter)
		allgrams[gram]++                                         // increment map counter for given ngram
	}
	ch <- allgrams
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

func setup(args []string) string {

	var incoming string

	// this checks whether there are command-line arguments. if so, it takes in all files as the corpus to check for trigrams.
	// if not, it checks whether stdin is coming from a pipe or the terminal. if from the terminal, it gives a message describing
	// what the program does and how to use it. if from a pipe (not terminal), it accepts the piped-in file(s) as input to process.

	if len(args) <= 1 {
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
	return incoming
}

func collectSequenceListConcurrent(words []string) []kv {
	var ngPartials []map[string]int
	var ngComplete map[string]int
	var sorted []kv

	// breaks up the file into bite-sized pieces
	wordPartials := breakup(words, 1000, 3)

	// finds ngrams in each separate piece, concurrently
	var wg sync.WaitGroup
	ch := make(chan map[string]int)
	defer close(ch)
	for _, p := range wordPartials {
		wg.Add(1)
		go func() {
			// pre-processed data is sent to look for three-word sequences/trigrams
			ngramFinderConcurrent(p, 3, ch)
			wg.Done()
		}()
		ngPartials = append(ngPartials, <-ch)
	}
	wg.Wait()

	// merges them back together
	ngComplete = mergeMaps(ngPartials...)

	// since maps in golang are inherently unordered, they cannot be sorted. therefore, an index of some sort is required, such as this slice of key-values
	for k, v := range ngComplete {
		sorted = append(sorted, kv{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	return sorted
}

func collectSequenceListSequential(words []string) []kv {
	var sorted []kv

	// pre-processed data is sent to look for three-word sequences/trigrams
	ng := ngramFinder(words, 3)

	// since maps in golang are inherently unordered, they cannot be sorted. therefore, an index of some sort is required, such as this slice of key-values

	for k, v := range ng {
		sorted = append(sorted, kv{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	return sorted
}

func displayOutput(sorted []kv) {
	// top 100 results are output in sorted order. if fewer than 100 results are present, however many are available are output in sorted order.
	fmt.Println("Rank: 3-Word Sequence - Count")
	fmt.Println("____________________________")
	if len(sorted) >= 100 {
		for i := 0; i < 100; i++ {
			fmt.Printf("%d:   %s - %d\n", i+1, sorted[i].Key, sorted[i].Value)
		}
	} else {
		for i := 0; i < len(sorted); i++ {
			fmt.Printf("%d:   %s - %d\n", i+1, sorted[i].Key, sorted[i].Value)
		}
	}
}

func main() {

	// initial command-line processing and accepting incoming data
	incoming := setup(os.Args)

	// incoming data is sent for pre-processing
	words := preprocess(string(incoming))

	// concurrent version of 3-word sequence processing
	fmt.Println("____________________________")
	fmt.Println("First, concurrent version:")
	sortedC := collectSequenceListConcurrent(words)
	displayOutput(sortedC)

	// sequential version of 3-word sequence processing
	fmt.Println("____________________________")
	fmt.Println("Next, sequential version:")
	sortedS := collectSequenceListSequential(words)
	displayOutput(sortedS)

}
