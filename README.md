# New Relic Golang Assessment 

## 100 most common three word sequences in text

###  Andrew Burt

This program counts three word sequences (trigrams, a case of n-grams) in an input text, and returns the top 100 ranked sequences with counts.

To run: 
`go run main.go file.txt` to count three word sequences in file.txt
`go run main.go file1.txt file2.txt file3.txt` to count three word sequences in file1.txt, file2.txt, and file3.txt
`cat file.txt|go run main.go` also counts three word sequences in file.txt
If no file is sent via command-line argument or a pipe, a message is returned explaining what the program does, and how to send it text files.

There are tests that can be run with `go test` for both the pre-processing function and the trigram/sequence detection function. 

Possible problem: The sample counts mentioned in the e-mail (the sperm whale - 85, the white whale - 71, of the whale - 67), don't all match my results. I did read rules several times in the e-mail, and tried interpreting them several ways (with all non-letter input ignored/removed including in-word punctuation like hyphens and apostrophes, with only between-word punctuation ignored, etc.), but the closest I got was my first version, which is in the program here. The top three it returns are (the sperm whale - 86, of the whale - 78, the white whale - 71). 
