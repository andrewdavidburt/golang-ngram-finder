package main

import (
	"testing"
)

var incoming = `But here is an artist. He desires to paint you the dreamiest, shadiest,
quietest, most enchanting bit of romantic landscape in all the valley
of the Saco. What is the chief element he employs? There stand his
trees, each with a hollow trunk, as if a hermit and a crucifix were
within; and here sleeps his meadow, and there sleep his cattle; and up
from yonder cottage goes a sleepy smoke. Deep into distant woodlands
winds a mazy way, reaching to overlapping sleeps his meadow`

var words = []string{"but", "here", "is", "an", "artist", "he", "desires", "to", "paint", "you", "the", "dreamiest", "shadiest", "quietest",
	"most", "enchanting", "bit", "of", "romantic", "landscape", "in", "all", "the", "valley", "of", "the", "saco", "what", "is", "the", "chief",
	"element", "he", "employs", "there", "stand", "his", "trees", "each", "with", "a", "hollow", "trunk", "as", "if", "a", "hermit", "and", "a",
	"crucifix", "were", "within", "and", "here", "sleeps", "his", "meadow", "and", "there", "sleep", "his", "cattle", "and", "up", "from", "yonder",
	"cottage", "goes", "a", "sleepy", "smoke", "deep", "into", "distant", "woodlands", "winds", "a", "mazy", "way", "reaching", "to", "overlapping",
	"sleeps", "his", "meadow"}

var ngs = map[string]int{"a crucifix were": 1, "a hermit and": 1, "a hollow trunk": 1, "a mazy way": 1, "a sleepy smoke": 1, "all the valley": 1,
	"an artist he": 1, "and a crucifix": 1, "and here sleeps": 1, "and there sleep": 1, "and up from": 1, "artist he desires": 1, "as if a": 1,
	"bit of romantic": 1, "but here is": 1, "cattle and up": 1, "chief element he": 1, "cottage goes a": 1, "crucifix were within": 1,
	"deep into distant": 1, "desires to paint": 1, "distant woodlands winds": 1, "dreamiest shadiest quietest": 1, "each with a": 1,
	"element he employs": 1, "employs there stand": 1, "enchanting bit of": 1, "from yonder cottage": 1, "goes a sleepy": 1, "he desires to": 1,
	"he employs there": 1, "here is an": 1, "here sleeps his": 1, "hermit and a": 1, "his cattle and": 1, "his meadow and": 1, "his trees each": 1,
	"hollow trunk as": 1, "if a hermit": 1, "in all the": 1, "into distant woodlands": 1, "is an artist": 1, "is the chief": 1, "landscape in all": 1,
	"mazy way reaching": 1, "meadow and there": 1, "most enchanting bit": 1, "of romantic landscape": 1, "of the saco": 1, "paint you the": 1,
	"quietest most enchanting": 1, "reaching to overlapping": 1, "romantic landscape in": 1, "saco what is": 1, "shadiest quietest most": 1,
	"sleep his cattle": 1, "sleeps his meadow": 2, "sleepy smoke deep": 1, "smoke deep into": 1, "stand his trees": 1, "the chief element": 1,
	"the dreamiest shadiest": 1, "the saco what": 1, "the valley of": 1, "there sleep his": 1, "there stand his": 1, "to paint you": 1,
	"trees each with": 1, "trunk as if": 1, "up from yonder": 1, "valley of the": 1, "way reaching to": 1, "were within and": 1, "what is the": 1,
	"winds a mazy": 1, "with a hollow": 1, "within and here": 1, "woodlands winds a": 1, "yonder cottage goes": 1, "you the dreamiest": 1,
	"to overlapping sleeps": 1, "overlapping sleeps his": 1}

// iterates through slice to compare values, assuming they're of the same length
func Equalslice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if s != b[i] {
			return false
		}
	}
	return true
}

func TestPreprocess(t *testing.T) {
	out := preprocess(incoming)
	if !Equalslice(out, words) {
		t.Fatalf("Expected %v, got %v", words, out)
	}
}
