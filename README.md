# Golang N-gram Finder
## 100 most common three word sequences in text
## Variation: AWS Lambda REST API version w/CircleCI CI/CD pipeline

###  Andrew Burt
### (`rest-lambda` branch)

This program counts three word sequences (trigrams, a case of n-grams) in an input text, and returns the top 100 ranked sequences with counts.  There is an alternate version, on the `rest-lambda` branch, that run as an AWS lambda function, and is auto-deployed using CircleCI (described in the addendum).

#### Usage
In order to make sure I understood the basics of CI/CD pipelines, I adapted the program to be an AWS lambda function, and then learned the basics of CircleCI to set up a very basic CI/CD pipeline. The source for this is sitting on the `rest-lambda` branch, found here:  

`https://github.com/andrewdavidburt/nr-assessment/tree/rest-lambda`  

The original command-line version is on the main branch, here:

`https://github.com/andrewdavidburt/nr-assessment`

Currently, every time I push to this branch, CircleCI tests and builds the code, then if everything passes it deploys it as an AWS lambda function behind an API Gateway.  
The deployed instance can be accessed in either of two ways:  
Simply but clunkily, you can send in words via the `text` arg on the uri:  

`curl https://vnp5rwu1u8.execute-api.us-west-2.amazonaws.com/staging?text=foobar+foobar+foobar+hello+hello+hello+hi+foobar+foobar+foobar`  

Slightly more elegantly, you can send in a percent-encoded URI via the `uri` arg on the uri:  

`curl https://vnp5rwu1u8.execute-api.us-west-2.amazonaws.com/staging?uri=https%3A%2F%2Fwww.gutenberg.org%2Ffiles%2F2701%2F2701-0.txt`  

This example specifically calls the Project Gutenberg source for the Moby Dick called for in the original code challenge.  
The results are returned as json.  

#### Notes

I wrote two versions of the 3-word sequence/ngram finder. One runs in sequence, the other in parallel. I'd originally wrote the sequential one on Tuesday, when I initially turned in the project, and then when I was told I had more time to work on it, I wrote a concurrent version. It breaks the file up into small chunks, then searches them all in parallel. It pads the chunks slightly in order to retain sequences that span chunks. I was hopeful that this would also speed up the application, but on testing the relative speeds, I found that it was very slightly slower (generally a few tenths of a second, depending on the size of the input file). As such, I wasn't sure I should delete the initial sequential version, being slightly faster, though I wanted to leave my work from the concurrent version. I broke out the functions specific to the concurrent implementation into a second file (concurrent.go), and a separate test file (concurrent_test.go), and retained the files both implementations had in common, as well as the initial version, in main.go .  When the program is run, it currently only runs the concurrent version, but the slightly faster sequential version is easily available by removing the "remark" slashes before the sequential implementation call in func main.  

If I had more time, and if memory consumption was a strong concern, I might  write it so it would only read in a smaller piece of the file at a time, or a line at a time, processing each as it came in and letting it go from memory after processing. I'd also explore other approaches to detecting three-word sequences. This approach is fast, but in a massive text, the memory usage could become prohibitive.  
