# Golang N-gram Finder
# Variation: AWS Lambda REST API version w/CircleCI CI/CD pipeline
## (`rest-lambda` branch)

## 100 most common three word sequences in text

###  Andrew Burt

This program counts three word sequences (trigrams, a case of n-grams) in an input text, and returns the top 100 ranked sequences with counts.  

The Readme on the main branch has all primary notes for the program. The Readme on this branch only contains a copy of the Addendum related to the AWS Lambda REST API version.

#### Addendum 
In order to make sure I understood the basics of CI/CD pipelines before the panel interviews, I adapted the program to be an AWS lambda function, and then learned the basics of CircleCI to set up a very basic CI/CD pipeline. The source for this is sitting on the `rest-lambda` branch, found here:  

`https://github.com/andrewdavidburt/nr-assessment/tree/rest-lambda`  

Currently, every time I push to this branch, CircleCI tests and builds the code, then if everything passes it deploys it as an AWS lambda function behind an API Gateway.  
The deployed instance can be accessed in either of two ways:  
Simply but clunkily, you can send in words via the `text` arg on the uri:  

`curl https://vnp5rwu1u8.execute-api.us-west-2.amazonaws.com/staging?text=foobar+foobar+foobar+hello+hello+hello+hi+foobar+foobar+foobar`  

Slightly more elegantly, you can send in a percent-encoded URI via the `uri` arg on the uri:  

`curl https://vnp5rwu1u8.execute-api.us-west-2.amazonaws.com/staging?uri=https%3A%2F%2Fwww.gutenberg.org%2Ffiles%2F2701%2F2701-0.txt`  

This example specifically calls the Project Gutenberg source for the Moby Dick called for in the original code challenge.  
The results are returned as json.  
