package main

import (
	"reflect"
	"testing"
)

var incoming = `But here is an artist. He desires to paint you the dreamiest, shadiest,
quietest, most enchanting bit of romantic landscape in all the valley
of the Saco. What is the chief element he employs? There stand his
trees, each with a hollow trunk, as if a hermit and a crucifix were
within; and here sleeps his meadow, and there sleep his cattle; and up
from yonder cottage goes a sleepy smoke. Deep into distant woodlands
winds a mazy way, reaching to overlapping spurs of mountains bathed in
their hill-side blue. But though the picture lies thus tranced, and
though this pine-tree shakes down its sighs like leaves upon this
shepherd’s head, yet all were vain, unless the shepherd’s eye were
fixed upon the magic stream before him. Go visit the Prairies in June,
when for scores on scores of miles you wade knee-deep among
Tiger-lilies—what is the one charm wanting?—Water—there is not a drop
of water there! Were Niagara but a cataract of sand, would you travel
your thousand miles to see it? Why did the poor poet of Tennessee, upon
suddenly receiving two handfuls of silver, deliberate whether to buy
him a coat, which he sadly needed, or invest his money in a pedestrian
trip to Rockaway Beach? Why is almost every robust healthy boy with a
robust healthy soul in him, at some time or other crazy to go to sea?
Why upon your first voyage as a passenger, did you yourself feel such a
mystical vibration, when first told that you and your ship were now out
of sight of land? Why did the old Persians hold the sea holy? Why did
the Greeks give it a separate deity, and own brother of Jove? Surely
all this is not without meaning. And still deeper the meaning of that
story of Narcissus, who because he could not grasp the tormenting, mild
image he saw in the fountain, plunged into it and was drowned. But that
same image, we ourselves see in all rivers and oceans. It is the image
of the ungraspable phantom of life; and this is the key to it all.

Now, when I say that I am in the habit of going to sea whenever I begin
to grow hazy about the eyes, and begin to be over conscious of my
lungs, I do not mean to have it inferred that I ever go to sea as a
passenger. For to go as a passenger you must needs have a purse, and a
purse is but a rag unless you have something in it. Besides, passengers
get sea-sick—grow quarrelsome—don’t sleep of nights—do not enjoy
themselves much, as a general thing;—no, I never go as a passenger;
nor, though I am something of a salt, do I ever go to sea as a
Commodore, or a Captain, or a Cook. I abandon the glory and distinction
of such offices to those who like them. For my part, I abominate all
honorable respectable toils, trials, and tribulations of every kind
whatsoever. It is quite as much as I can do to take care of myself,
without taking care of ships, barques, brigs, schooners, and what not.
And as for going as cook,—though I confess there is considerable glory
in that, a cook being a sort of officer on ship-board—yet, somehow, I
never fancied broiling fowls;—though once broiled, judiciously
buttered, and judgmatically salted and peppered, there is no one who
will speak more respectfully, not to say reverentially, of a broiled
fowl than I will. It is out of the idolatrous dotings of the old
Egyptians upon broiled ibis and roasted river horse, that you see the
mummies of those creatures in their huge bake-houses the pyramids.
`

var words = []string{"but", "here", "is", "an", "artist", "he", "desires", "to", "paint", "you", "the", "dreamiest", "shadiest", "quietest",
	"most", "enchanting", "bit", "of", "romantic", "landscape", "in", "all", "the", "valley", "of", "the", "saco", "what", "is", "the", "chief",
	"element", "he", "employs", "there", "stand", "his", "trees", "each", "with", "a", "hollow", "trunk", "as", "if", "a", "hermit", "and", "a",
	"crucifix", "were", "within", "and", "here", "sleeps", "his", "meadow", "and", "there", "sleep", "his", "cattle", "and", "up", "from", "yonder",
	"cottage", "goes", "a", "sleepy", "smoke", "deep", "into", "distant", "woodlands", "winds", "a", "mazy", "way", "reaching", "to", "overlapping",
	"spurs", "of", "mountains", "bathed", "in", "their", "hillside", "blue", "but", "though", "the", "picture", "lies", "thus", "tranced", "and",
	"though", "this", "pinetree", "shakes", "down", "its", "sighs", "like", "leaves", "upon", "this", "shepherds", "head", "yet", "all", "were",
	"vain", "unless", "the", "shepherds", "eye", "were", "fixed", "upon", "the", "magic", "stream", "before", "him", "go", "visit", "the", "prairies",
	"in", "june", "when", "for", "scores", "on", "scores", "of", "miles", "you", "wade", "kneedeep", "among", "tigerlilieswhat", "is", "the", "one",
	"charm", "wantingwaterthere", "is", "not", "a", "drop", "of", "water", "there", "were", "niagara", "but", "a", "cataract", "of", "sand", "would",
	"you", "travel", "your", "thousand", "miles", "to", "see", "it", "why", "did", "the", "poor", "poet", "of", "tennessee", "upon", "suddenly",
	"receiving", "two", "handfuls", "of", "silver", "deliberate", "whether", "to", "buy", "him", "a", "coat", "which", "he", "sadly", "needed", "or",
	"invest", "his", "money", "in", "a", "pedestrian", "trip", "to", "rockaway", "beach", "why", "is", "almost", "every", "robust", "healthy", "boy",
	"with", "a", "robust", "healthy", "soul", "in", "him", "at", "some", "time", "or", "other", "crazy", "to", "go", "to", "sea", "why", "upon",
	"your", "first", "voyage", "as", "a", "passenger", "did", "you", "yourself", "feel", "such", "a", "mystical", "vibration", "when", "first", "told",
	"that", "you", "and", "your", "ship", "were", "now", "out", "of", "sight", "of", "land", "why", "did", "the", "old", "persians", "hold", "the",
	"sea", "holy", "why", "did", "the", "greeks", "give", "it", "a", "separate", "deity", "and", "own", "brother", "of", "jove", "surely", "all",
	"this", "is", "not", "without", "meaning", "and", "still", "deeper", "the", "meaning", "of", "that", "story", "of", "narcissus", "who", "because",
	"he", "could", "not", "grasp", "the", "tormenting", "mild", "image", "he", "saw", "in", "the", "fountain", "plunged", "into", "it", "and", "was",
	"drowned", "but", "that", "same", "image", "we", "ourselves", "see", "in", "all", "rivers", "and", "oceans", "it", "is", "the", "image", "of",
	"the", "ungraspable", "phantom", "of", "life", "and", "this", "is", "the", "key", "to", "it", "all", "now", "when", "i", "say", "that", "i", "am",
	"in", "the", "habit", "of", "going", "to", "sea", "whenever", "i", "begin", "to", "grow", "hazy", "about", "the", "eyes", "and", "begin", "to",
	"be", "over", "conscious", "of", "my", "lungs", "i", "do", "not", "mean", "to", "have", "it", "inferred", "that", "i", "ever", "go", "to", "sea",
	"as", "a", "passenger", "for", "to", "go", "as", "a", "passenger", "you", "must", "needs", "have", "a", "purse", "and", "a", "purse", "is", "but",
	"a", "rag", "unless", "you", "have", "something", "in", "it", "besides", "passengers", "get", "seasickgrow", "quarrelsomedont", "sleep", "of",
	"nightsdo", "not", "enjoy", "themselves", "much", "as", "a", "general", "thingno", "i", "never", "go", "as", "a", "passenger", "nor", "though",
	"i", "am", "something", "of", "a", "salt", "do", "i", "ever", "go", "to", "sea", "as", "a", "commodore", "or", "a", "captain", "or", "a", "cook",
	"i", "abandon", "the", "glory", "and", "distinction", "of", "such", "offices", "to", "those", "who", "like", "them", "for", "my", "part", "i",
	"abominate", "all", "honorable", "respectable", "toils", "trials", "and", "tribulations", "of", "every", "kind", "whatsoever", "it", "is", "quite",
	"as", "much", "as", "i", "can", "do", "to", "take", "care", "of", "myself", "without", "taking", "care", "of", "ships", "barques", "brigs",
	"schooners", "and", "what", "not", "and", "as", "for", "going", "as", "cookthough", "i", "confess", "there", "is", "considerable", "glory", "in",
	"that", "a", "cook", "being", "a", "sort", "of", "officer", "on", "shipboardyet", "somehow", "i", "never", "fancied", "broiling", "fowlsthough",
	"once", "broiled", "judiciously", "buttered", "and", "judgmatically", "salted", "and", "peppered", "there", "is", "no", "one", "who", "will",
	"speak", "more", "respectfully", "not", "to", "say", "reverentially", "of", "a", "broiled", "fowl", "than", "i", "will", "it", "is", "out", "of",
	"the", "idolatrous", "dotings", "of", "the", "old", "egyptians", "upon", "broiled", "ibis", "and", "roasted", "river", "horse", "that", "you",
	"see", "the", "mummies", "of", "those", "creatures", "in", "their", "huge", "bakehouses", "the", "pyramids"}

var ngs = map[string]uint32{
	"you yourself feel": 1, "out of sight": 1, "a passenger you": 1, "go as a": 2, "a purse and": 1, "you have something": 1, "not and as": 1,
	"magic stream before": 1, "without meaning and": 1, "all rivers and": 1, "going to sea": 1, "tigerlilieswhat is the": 1,
	"old persians hold": 1, "narcissus who because": 1, "am something of": 1, "care of myself": 1, "what not and": 1, "a sort of": 1,
	"his money in": 1, "deeper the meaning": 1, "to have it": 1, "a salt do": 1, "beach why is": 1, "almost every robust": 1, "not mean to": 1,
	"somehow i never": 1, "vibration when first": 1, "to sea whenever": 1, "and roasted river": 1, "every robust healthy": 1, "sea holy why": 1,
	"still deeper the": 1, "and distinction of": 1, "of going to": 1, "upon your first": 1, "land why did": 1, "the image of": 1,
	"ships barques brigs": 1, "in their hillside": 1, "not a drop": 1, "silver deliberate whether": 1, "to sea why": 1, "see the mummies": 1,
	"dreamiest shadiest quietest": 1, "upon suddenly receiving": 1, "image of the": 1, "fancied broiling fowlsthough": 1, "poet of tennessee": 1,
	"is quite as": 1, "officer on shipboardyet": 1, "do not mean": 1, "you must needs": 1, "it besides passengers": 1, "kind whatsoever it": 1,
	"crucifix were within": 1, "meadow and there": 1, "blue but though": 1, "first told that": 1, "as much as": 1, "i never fancied": 1,
	"a pedestrian trip": 1, "rockaway beach why": 1, "that i am": 1, "reverentially of a": 1, "the habit of": 1, "habit of going": 1,
	"in it besides": 1, "passenger did you": 1, "and oceans it": 1, "now when i": 1, "when i say": 1, "with a hollow": 1,
	"one charm wantingwaterthere": 1, "toils trials and": 1, "of life and": 1, "to grow hazy": 1, "fowl than i": 1, "here is an": 1,
	"of romantic landscape": 1, "and here sleeps": 1, "yonder cottage goes": 1, "but here is": 1, "the prairies in": 1, "can do to": 1,
	"pedestrian trip to": 1, "when first told": 1, "rag unless you": 1, "and tribulations of": 1, "much as a": 1, "a cook i": 1,
	"a hermit and": 1, "scores on scores": 1, "wantingwaterthere is not": 1, "and still deeper": 1, "take care of": 1, "stand his trees": 1,
	"ever go to": 2, "general thingno i": 1, "i can do": 1, "much as i": 1, "artist he desires": 1, "of water there": 1,
	"enjoy themselves much": 1, "for my part": 1, "of officer on": 1, "though the picture": 1, "did you yourself": 1, "of narcissus who": 1,
	"get seasickgrow quarrelsomedont": 1, "saw in the": 1, "cook being a": 1, "broiled judiciously buttered": 1, "sleepy smoke deep": 1,
	"two handfuls of": 1, "persians hold the": 1, "in their huge": 1, "a passenger for": 1, "purse is but": 1, "do to take": 1,
	"taking care of": 1, "whatsoever it is": 1, "for going as": 1, "i will it": 1, "is out of": 1, "go visit the": 1, "soul in him": 1,
	"is the key": 1, "or a cook": 1, "to overlapping spurs": 1, "fixed upon the": 1, "all this is": 1, "passengers get seasickgrow": 1,
	"an artist he": 1, "most enchanting bit": 1, "there stand his": 1, "a sleepy smoke": 1, "i do not": 1, "roasted river horse": 1,
	"is the one": 1, "buy him a": 1, "of land why": 1, "the eyes and": 1, "first voyage as": 1, "unless you have": 1,
	"nightsdo not enjoy": 1, "buttered and judgmatically": 1, "meaning of that": 1, "dotings of the": 1, "shakes down its": 1,
	"niagara but a": 1, "phantom of life": 1, "or a captain": 1, "to take care": 1, "will it is": 1, "sleeps his meadow": 1,
	"scores of miles": 1, "sea why upon": 1, "i say that": 1, "for scores on": 1, "he could not": 1, "the ungraspable phantom": 1,
	"brigs schooners and": 1, "creatures in their": 1, "their huge bakehouses": 1, "bathed in their": 1, "a robust healthy": 1,
	"fountain plunged into": 1, "sea whenever i": 1, "valley of the": 1, "ship were now": 1, "trunk as if": 1, "goes a sleepy": 1,
	"sand would you": 1, "broiled ibis and": 1, "drowned but that": 1, "i begin to": 1, "and begin to": 1, "care of ships": 1,
	"way reaching to": 1, "a coat which": 1, "now out of": 1, "that story of": 1, "huge bakehouses the": 1, "passenger nor though": 1,
	"into it and": 1, "whenever i begin": 1, "begin to grow": 1, "needs have a": 1, "quite as much": 1, "element he employs": 1,
	"like leaves upon": 1, "grow hazy about": 1, "every kind whatsoever": 1, "glory in that": 1, "desires to paint": 1, "a crucifix were": 1,
	"all now when": 1, "those who like": 1, "jove surely all": 1, "the old egyptians": 1, "upon this shepherds": 1, "yet all were": 1,
	"handfuls of silver": 1, "deliberate whether to": 1, "of my lungs": 1, "themselves much as": 1, "captain or a": 1, "of every kind": 1,
	"this shepherds head": 1, "vain unless the": 1, "there were niagara": 1, "whether to buy": 1, "thus tranced and": 1,
	"glory and distinction": 1, "i ever go": 2, "speak more respectfully": 1, "bakehouses the pyramids": 1, "his meadow and": 1,
	"he sadly needed": 1, "the old persians": 1, "begin to be": 1, "purse and a": 1, "respectfully not to": 1, "is an artist": 1,
	"down its sighs": 1, "passenger for to": 1, "for to go": 1, "the shepherds eye": 1, "am in the": 1, "hazy about the": 1,
	"and a purse": 1, "i abandon the": 1, "barques brigs schooners": 1, "such offices to": 1, "and though this": 1, "you travel your": 1,
	"is not without": 1, "grasp the tormenting": 1, "his cattle and": 1, "a captain or": 1, "ibis and roasted": 1, "you wade kneedeep": 1,
	"with a robust": 1, "one who will": 1, "seasickgrow quarrelsomedont sleep": 1, "though i am": 1, "do i ever": 1, "to paint you": 1,
	"mystical vibration when": 1, "brother of jove": 1, "key to it": 1, "deep into distant": 1, "why is almost": 1, "a mystical vibration": 1,
	"there is considerable": 1, "commodore or a": 1, "saco what is": 1, "poor poet of": 1, "robust healthy soul": 1, "of sight of": 1,
	"because he could": 1, "but that same": 1, "as a general": 1, "that you see": 1, "all the valley": 1, "pinetree shakes down": 1,
	"it a separate": 1, "the meaning of": 1, "a cataract of": 1, "a cook being": 1, "once broiled judiciously": 1, "a rag unless": 1,
	"will speak more": 1, "other crazy to": 1, "that you and": 1, "own brother of": 1, "could not grasp": 1, "than i will": 1,
	"healthy boy with": 1, "why upon your": 1, "of nightsdo not": 1, "and as for": 1, "salt do i": 1, "abandon the glory": 1,
	"the glory and": 1, "as cookthough i": 1, "all were vain": 1, "water there were": 1, "thousand miles to": 1, "have a purse": 1,
	"broiled fowl than": 1, "the saco what": 1, "tennessee upon suddenly": 1, "part i abominate": 1, "idolatrous dotings of": 1,
	"kneedeep among tigerlilieswhat": 1, "deity and own": 1, "same image we": 1, "schooners and what": 1, "robust healthy boy": 1,
	"the poor poet": 1, "conscious of my": 1, "shipboardyet somehow i": 1, "old egyptians upon": 1, "he employs there": 1, "trees each with": 1,
	"spurs of mountains": 1, "eye were fixed": 1, "a drop of": 1, "told that you": 1, "a commodore or": 1, "distinction of such": 1,
	"out of the": 1, "those creatures in": 1, "reaching to overlapping": 1, "of mountains bathed": 1, "a passenger did": 1,
	"quarrelsomedont sleep of": 1, "miles you wade": 1, "is almost every": 1, "nor though i": 1, "judgmatically salted and": 1,
	"in all the": 1, "why did the": 3, "confess there is": 1, "river horse that": 1, "must needs have": 1, "respectable toils trials": 1,
	"a broiled fowl": 1, "a hollow trunk": 1, "him go visit": 1, "story of narcissus": 1, "to go as": 1, "thingno i never": 1,
	"the dreamiest shadiest": 1, "hollow trunk as": 1, "to rockaway beach": 1, "eyes and begin": 1, "over conscious of": 1,
	"lies thus tranced": 1, "of that story": 1, "the fountain plunged": 1, "in all rivers": 1, "were vain unless": 1, "your thousand miles": 1,
	"like them for": 1, "trials and tribulations": 1, "sighs like leaves": 1, "did the poor": 1, "the sea holy": 1, "tormenting mild image": 1,
	"besides passengers get": 1, "abominate all honorable": 1, "of miles you": 1, "it inferred that": 1, "it is out": 1, "is the chief": 1,
	"it why did": 1, "trip to rockaway": 1, "image we ourselves": 1, "and your ship": 1, "and own brother": 1, "him at some": 1,
	"holy why did": 1, "but a rag": 1, "have something in": 1, "healthy soul in": 1, "this is not": 1, "sleep of nightsdo": 1, "not to say": 1,
	"into distant woodlands": 1, "but though the": 1, "him a coat": 1, "go to sea": 3, "meaning and still": 1, "something in it": 1,
	"who like them": 1, "quietest most enchanting": 1, "hermit and a": 1, "miles to see": 1, "invest his money": 1, "of ships barques": 1,
	"and what not": 1, "that a cook": 1, "never fancied broiling": 1, "is considerable glory": 1, "you see the": 1, "he desires to": 1,
	"cataract of sand": 1, "hold the sea": 1, "plunged into it": 1, "sleep his cattle": 1, "were fixed upon": 1, "sadly needed or": 1,
	"to be over": 1, "a purse is": 1, "to those who": 1, "broiling fowlsthough once": 1, "chief element he": 1, "smoke deep into": 1,
	"wade kneedeep among": 1, "your first voyage": 1, "surely all this": 1, "who because he": 1, "not without meaning": 1, "as i can": 1,
	"cookthough i confess": 1, "fowlsthough once broiled": 1, "the valley of": 1, "head yet all": 1, "travel your thousand": 1,
	"receiving two handfuls": 1, "at some time": 1, "the greeks give": 1, "not grasp the": 1, "on shipboardyet somehow": 1, "paint you the": 1,
	"romantic landscape in": 1, "woodlands winds a": 1, "charm wantingwaterthere is": 1, "and peppered there": 1, "is not a": 1,
	"yourself feel such": 1, "sight of land": 1, "landscape in all": 1, "see it why": 1, "of those creatures": 1, "shepherds eye were": 1,
	"drop of water": 1, "no one who": 1, "each with a": 1, "if a hermit": 1, "i am in": 1, "money in a": 1, "a separate deity": 1,
	"image he saw": 1, "my lungs i": 1, "what is the": 1, "though this pinetree": 1, "leaves upon this": 1, "unless the shepherds": 1,
	"i abominate all": 1, "greeks give it": 1, "it all now": 1, "of the old": 1, "something of a": 1, "of a salt": 1, "winds a mazy": 1,
	"visit the prairies": 1, "in the habit": 1, "mean to have": 1, "june when for": 1, "inferred that i": 1, "as for going": 1,
	"of the idolatrous": 1, "time or other": 1, "you and your": 1, "your ship were": 1, "did the greeks": 1, "a passenger nor": 1,
	"judiciously buttered and": 1, "say reverentially of": 1, "of a broiled": 1, "here sleeps his": 1, "and there sleep": 1, "the picture lies": 1,
	"on scores of": 1, "in june when": 1, "cook i abandon": 1, "i confess there": 1, "of such offices": 1, "salted and peppered": 1,
	"cattle and up": 1, "or invest his": 1, "is but a": 1, "i am something": 1, "shadiest quietest most": 1, "and up from": 1, "to buy him": 1,
	"is no one": 1, "in a pedestrian": 1, "in him at": 1, "the tormenting mild": 1, "about the eyes": 1, "employs there stand": 1,
	"its sighs like": 1, "suddenly receiving two": 1, "coat which he": 1, "boy with a": 1, "did the old": 1, "it is quite": 1,
	"needed or invest": 1, "that i ever": 1, "sea as a": 2, "all honorable respectable": 1, "mountains bathed in": 1, "among tigerlilieswhat is": 1,
	"a mazy way": 1, "shepherds head yet": 1, "who will speak": 1, "bit of romantic": 1, "we ourselves see": 1, "say that i": 1,
	"being a sort": 1, "cottage goes a": 1, "when for scores": 1, "horse that you": 1, "but a cataract": 1, "is the image": 1,
	"to say reverentially": 1, "this is the": 1, "passenger you must": 1, "my part i": 1, "tribulations of every": 1, "distant woodlands winds": 1,
	"tranced and though": 1, "would you travel": 1, "was drowned but": 1, "more respectfully not": 1, "of the ungraspable": 1, "he saw in": 1,
	"have it inferred": 1, "as if a": 1, "to go to": 1, "were now out": 1, "give it a": 1, "be over conscious": 1, "considerable glory in": 1,
	"sort of officer": 1, "his trees each": 1, "which he sadly": 1, "as a passenger": 4, "mild image he": 1, "feel such a": 1,
	"going as cookthough": 1, "upon broiled ibis": 1, "voyage as a": 1, "ourselves see in": 1, "never go as": 1, "myself without taking": 1,
	"the mummies of": 1, "within and here": 1, "of silver deliberate": 1, "not enjoy themselves": 1, "honorable respectable toils": 1,
	"of myself without": 1, "mummies of those": 1, "were within and": 1, "from yonder cottage": 1, "upon the magic": 1, "and this is": 1,
	"offices to those": 1, "it and was": 1, "rivers and oceans": 1, "the key to": 1, "lungs i do": 1, "of sand would": 1, "to see it": 1,
	"separate deity and": 1, "of jove surely": 1, "i never go": 1, "the idolatrous dotings": 1, "mazy way reaching": 1, "their hillside blue": 1,
	"of tennessee upon": 1, "enchanting bit of": 1, "some time or": 1, "and was drowned": 1, "egyptians upon broiled": 1, "up from yonder": 1,
	"that same image": 1, "see in all": 1, "in that a": 1, "crazy to go": 1, "in the fountain": 1, "it is the": 1, "life and this": 1,
	"to sea as": 2, "as a commodore": 1, "and a crucifix": 1, "were niagara but": 1, "such a mystical": 1, "oceans it is": 1, "to it all": 1,
	"peppered there is": 1, "overlapping spurs of": 1, "the magic stream": 1, "before him go": 1, "the one charm": 1, "this pinetree shakes": 1,
	"without taking care": 1, "picture lies thus": 1, "ungraspable phantom of": 1, "a general thingno": 1, "them for my": 1, "of the saco": 1,
	"hillside blue but": 1, "stream before him": 1, "there is no": 1, "there sleep his": 1, "prairies in june": 1, "and judgmatically salted": 1,
	"you the dreamiest": 1, "the chief element": 1, "or other crazy": 1,
}

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
	out := Preprocess(incoming)
	if !Equalslice(out, words) {
		t.Fatalf("Expected %v, got %v", words, out)
	}
}

func TestNgrams(t *testing.T) {
	ng := ngrams(words, 3)
	if !reflect.DeepEqual(ng, ngs) {
		t.Fatalf("Expected %v, got %v", ngs, ng)
	}
}
