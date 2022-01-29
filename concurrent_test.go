package main

import (
	"reflect"
	"sync"
	"testing"
)

var wordsOrigin = []string{"but", "here", "is", "an", "artist", "he", "desires", "to", "paint", "you",
	"the", "dreamiest", "shadiest", "quietest", "most", "enchanting", "bit", "of", "romantic", "landscape",
	"in", "all", "the", "valley", "of", "the", "saco", "what", "is", "the", "chief", "element", "he", "employs", "there",
	"stand", "his", "trees", "each", "with", "a", "hollow", "trunk", "as", "if", "a", "hermit", "but", "here", "is",
	"but", "here", "is", "hello", "there", "hi", "there", "is", "hello", "there"}

var words1 = []string{"but", "here", "is", "an", "artist", "he", "desires", "to", "paint", "you", "the", "dreamiest"}
var words2 = []string{"the", "dreamiest", "shadiest", "quietest", "most", "enchanting", "bit", "of", "romantic", "landscape", "in", "all"}
var words3 = []string{"in", "all", "the", "valley", "of", "the", "saco", "what", "is", "the", "chief", "element"}
var words4 = []string{"chief", "element", "he", "employs", "there", "stand", "his", "trees", "each", "with", "a", "hollow"}
var words5 = []string{"a", "hollow", "trunk", "as", "if", "a", "hermit", "but", "here", "is", "but", "here"}
var words6 = []string{"but", "here", "is", "hello", "there", "hi", "there", "is", "hello", "there"}

var wordPartials = [][]string{
	words1, words2, words3, words4, words5, words6,
}

var ngPartial1 = map[string]int{
	"an artist he": 1, "artist he desires": 1, "but here is": 1, "desires to paint": 1, "he desires to": 1, "here is an": 1, "is an artist": 1,
	"paint you the": 1, "to paint you": 1, "you the dreamiest": 1,
}
var ngPartial2 = map[string]int{
	"bit of romantic": 1, "dreamiest shadiest quietest": 1, "enchanting bit of": 1, "landscape in all": 1, "most enchanting bit": 1,
	"of romantic landscape": 1, "quietest most enchanting": 1, "romantic landscape in": 1, "shadiest quietest most": 1, "the dreamiest shadiest": 1,
}
var ngPartial3 = map[string]int{
	"all the valley": 1, "in all the": 1, "is the chief": 1, "of the saco": 1, "saco what is": 1, "the chief element": 1, "the saco what": 1,
	"the valley of": 1, "valley of the": 1, "what is the": 1,
}
var ngPartial4 = map[string]int{
	"chief element he": 1, "each with a": 1, "element he employs": 1, "employs there stand": 1, "he employs there": 1, "his trees each": 1,
	"stand his trees": 1, "there stand his": 1, "trees each with": 1, "with a hollow": 1,
}
var ngPartial5 = map[string]int{
	"a hermit but": 1, "a hollow trunk": 1, "as if a": 1, "but here is": 1, "here is but": 1, "hermit but here": 1, "hollow trunk as": 1,
	"if a hermit": 1, "is but here": 1, "trunk as if": 1,
}
var ngPartial6 = map[string]int{
	"but here is": 1, "hello there hi": 1, "here is hello": 1, "hi there is": 1, "is hello there": 2, "there hi there": 1, "there is hello": 1,
}
var ngPartials = []map[string]int{
	ngPartial1, ngPartial2, ngPartial3, ngPartial4, ngPartial5, ngPartial6,
}

var ngComplete = map[string]int{
	"an artist he": 1, "artist he desires": 1, "but here is": 3, "desires to paint": 1, "he desires to": 1, "here is an": 1, "is an artist": 1,
	"paint you the": 1, "to paint you": 1, "you the dreamiest": 1, "bit of romantic": 1, "dreamiest shadiest quietest": 1, "enchanting bit of": 1,
	"landscape in all": 1, "most enchanting bit": 1, "of romantic landscape": 1, "quietest most enchanting": 1, "romantic landscape in": 1,
	"shadiest quietest most": 1, "the dreamiest shadiest": 1, "all the valley": 1, "in all the": 1, "is the chief": 1, "of the saco": 1,
	"saco what is": 1, "the chief element": 1, "the saco what": 1, "the valley of": 1, "valley of the": 1, "what is the": 1, "chief element he": 1,
	"each with a": 1, "element he employs": 1, "employs there stand": 1, "he employs there": 1, "his trees each": 1, "stand his trees": 1,
	"there stand his": 1, "trees each with": 1, "with a hollow": 1, "a hermit but": 1, "a hollow trunk": 1, "as if a": 1, "here is but": 1,
	"hermit but here": 1, "hollow trunk as": 1, "if a hermit": 1, "is but here": 1, "trunk as if": 1, "hello there hi": 1, "here is hello": 1,
	"hi there is": 1, "is hello there": 2, "there hi there": 1, "there is hello": 1,
}

func TestNgramFinderConcurrent(t *testing.T) {
	var ngPartials2 []map[string]int
	// var ngComplete map[string]int

	var wg sync.WaitGroup
	ch := make(chan map[string]int)
	defer close(ch)
	for _, p := range wordPartials {
		wg.Add(1)
		go func() {
			ngramFinderConcurrent(p, 3, ch)
			wg.Done()
		}()
		ngPartials2 = append(ngPartials2, <-ch)
	}
	wg.Wait()

	if !reflect.DeepEqual(ngPartials, ngPartials2) {
		t.Fatalf("Expected %v, got %v", ngPartials, ngPartials2)
	}
}

func TestBreakup(t *testing.T) {
	wordPartials2 := breakup(wordsOrigin, 10, 3)

	if !reflect.DeepEqual(wordPartials, wordPartials2) {
		t.Fatalf("Expected %v, got %v", wordPartials, wordPartials2)
	}
}

func TestMergeMaps(t *testing.T) {
	ngComplete2 := mergeMaps(ngPartials...)
	if !reflect.DeepEqual(ngComplete, ngComplete2) {
		t.Fatalf("Expected %v, got %v", ngComplete, ngComplete2)
	}
}
