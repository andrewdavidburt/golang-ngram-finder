package main

import (
	"sort"
	"strings"
	"sync"
)

func collectSequenceListConcurrent(words []string) []kv {
	var ngPartials []map[string]int
	var ngComplete map[string]int
	var sorted []kv

	// breaks up the file into bite-sized pieces
	wordPartials := breakup(words, 100, 3)

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

	// since maps in golang are inherently unordered, they cannot be sorted.
	// therefore, an index of some sort is required, such as this slice of key-values
	for k, v := range ngComplete {
		sorted = append(sorted, kv{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	return sorted
}

// channeled version of ngram finder
func ngramFinderConcurrent(words []string, size int, ch chan<- map[string]int) {
	allgrams := make(map[string]int)
	offset := size / 2
	max := len(words)
	for i := range words {
		//  don't run ngram finder where it will run off the beginning or end of the collection
		if i < offset || i+size-offset > max {
			continue
		}
		// collect ngram from words in collection of n/size length (to either side of counter)
		gram := strings.Join(words[i-offset:i+size-offset], " ")
		// increment map counter for given ngram
		allgrams[gram]++
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

// breaks up file into arbitrarily sized chunks
func breakup(complete []string, chunkSize int, ngramLength int) [][]string {
	var partials [][]string
	for i := 0; i < len(complete); i += chunkSize {
		// retention of n-grams/trigrams spanning two chunks by adding n-1 (2) to chunk size
		end := i + chunkSize + ngramLength - 1
		if end > len(complete) {
			end = len(complete)
		}
		partials = append(partials, complete[i:end])
	}
	return partials
}
