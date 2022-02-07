package main

import (
	"encoding/json"
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

//for aws lambda error returns
var errorLogger = log.New(os.Stderr, "error: ", log.Llongfile)

//for json output of sorted triads
type Output struct {
	Group    string `json:"group"`
	Rank     int    `json:"rank"`
	Sequence string `json:"sequence"`
	Count    int    `json:"count"`
}

// for sorting map
type kv struct {
	Key   string
	Value int
}

//create json responses
func formResponse(sorted []kv) ([]*Output, error) {
	var out []*Output
	if len(sorted) >= 100 {
		for i := 0; i < 100; i++ {
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

//create aws lambda server errors
func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

//preprocessing of text input
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

//outgoing call to get a text document specified by uri arg
func callout(uri string) ([]byte, error) {
	client := &http.Client{}
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "text/plain")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func manager(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var words []string
	if val, ok := req.QueryStringParameters["uri"]; ok {
		body, err := callout(val)
		if err != nil {
			return serverError(err)
		}
		words = preprocess(string(body))
	} else {
		words = preprocess(req.QueryStringParameters["text"])
	}

	sortedC := collectSequenceListConcurrent(words)

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
