package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errorLogger = log.New(os.Stderr, "error: ", log.Llongfile)

type Output struct {
	Group    string `json:"group"`
	Rank     int    `json:"rank"`
	Sequence string `json:"sequence"`
	Count    int    `json:"count"`
}

func formResponse(sorted []kv) ([]*Output, error) {
	var out []*Output
	if len(sorted) >= 10 {
		for i := 0; i < 10; i++ {
			out = append(out, &Output{
				Group:    "group",
				Rank:     i + 1,
				Sequence: sorted[i].Key,
				Count:    sorted[i].Value,
			})
		}
	} else {
		for i := 0; i < len(sorted); i++ {
			out = append(out, &Output{
				Group:    "group",
				Rank:     i + 1,
				Sequence: sorted[i].Key,
				Count:    sorted[i].Value,
			})

		}
	}
	return out, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

// for sorting map
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
// func openFile(filename string) []byte {
// 	data, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		log.Panicf("couldn't read file: %s", err)
// 	}
// 	return data
// }

// checks for arbitrary-length sequences
func ngramFinder(words []string, size int) (allgrams map[string]int) {
	allgrams = make(map[string]int)
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
	return allgrams
}

// func setup(args []string) string {

// 	var incoming string

// 	// this checks whether there are command-line arguments.
// 	// if so, it takes in all files as the corpus to check for sequences.
// 	// if there are none, it checks whether stdin is coming from a pipe or the terminal.
// 	// if stdin is from the terminal, it gives the user a message describing what the program does and how to use it.
// 	// if stdin is from a pipe (not terminal), it accepts the piped-in file(s) as input to process.

// 	if len(args) <= 1 {
// 		stat, _ := os.Stdin.Stat()
// 		if (stat.Mode() & os.ModeCharDevice) == 0 {
// 			reader := bufio.NewReader(os.Stdin)
// 			var holding []rune

// 			for {
// 				input, _, err := reader.ReadRune()
// 				if err != nil && err == io.EOF {
// 					break
// 				}
// 				holding = append(holding, input)
// 			}
// 			incoming = string(holding)

// 		} else {
// 			fmt.Println("This program counts 3-word sequences (trigrams) in a document, and outputs the top 100 in order. ")
// 			fmt.Println("Please either specify one or more text files as arguments on the command-line, or pipe text in via stdin.")
// 			fmt.Println("usage examples (if run from source without build)")
// 			fmt.Println("go run . moby-dick.txt")
// 			fmt.Println("cat moby-dick.txt|go run .")
// 			os.Exit(0)
// 		}
// 	} else {
// 		for _, file := range os.Args[1:] {
// 			incoming += string(openFile(file))
// 		}

// 	}
// 	return incoming
// }

// func collectSequenceListSequential(words []string) []kv {
// 	var sorted []kv

// 	// pre-processed data is sent to look for three-word sequences/trigrams
// 	ng := ngramFinder(words, 3)

// 	// since maps in golang are inherently unordered, they cannot be sorted.
// 	// therefore, an index of some sort is required, such as this slice of key-values
// 	for k, v := range ng {
// 		sorted = append(sorted, kv{k, v})
// 	}
// 	sort.Slice(sorted, func(i, j int) bool {
// 		return sorted[i].Value > sorted[j].Value
// 	})

// 	return sorted
// }

// func displayOutput(sorted []kv) {
// 	// top 100 results are output in sorted order. if fewer than 100 results are present,
// 	// however many are available are output in sorted order.
// 	fmt.Println("Rank: 3-Word Sequence - Count")
// 	fmt.Println("____________________________")
// 	if len(sorted) >= 100 {
// 		for i := 0; i < 100; i++ {
// 			fmt.Printf("%d:   %s - %d\n", i+1, sorted[i].Key, sorted[i].Value)
// 		}
// 	} else {
// 		for i := 0; i < len(sorted); i++ {
// 			fmt.Printf("%d:   %s - %d\n", i+1, sorted[i].Key, sorted[i].Value)
// 		}
// 	}
// }

func callout(uri string) ([]byte, error) {
	log.Println("test4")
	client := &http.Client{}

	_, err := url.ParseRequestURI(uri)
	if err == nil {
		return nil, errors.New(fmt.Sprint(http.StatusBadRequest))
	}
	log.Println("test5")
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	log.Println("test6")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	log.Println("test7")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func manager(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var words []string
	client := &http.Client{}
	log.Println("test1")
	if _, ok := req.QueryStringParameters["uri"]; ok {
		log.Println("test2")
		// body, err := callout(val)
		req, err := http.NewRequest("GET", "https://www.gutenberg.org/files/2701/2701-0.txt", nil)
		resp, err := client.Do(req)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		log.Println("test3")
		if err != nil {
			return serverError(err)
		}
		words = preprocess(string(body))
	} else {
		words = preprocess(req.QueryStringParameters["text"])
	}
	// incoming := setup(req.QueryStringParameters["text"])

	sortedC := collectSequenceListConcurrent(words)
	// displayOutput(sortedC)

	// sortedS := collectSequenceListSequential(words)
	// displayOutput(sortedS)

	// body, err := callout(req.QueryStringParameters["ip"])
	// if err != nil {
	// 	return serverError(err)
	// }

	out, err := formResponse(sortedC)
	if err != nil {
		return serverError(err)
	}

	jsout, err := json.Marshal(out)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(jsout),
	}, nil

}

func main() {
	lambda.Start(manager)
}
