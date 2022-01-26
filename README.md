# New Relic Golang Assessment 

## 100 most common three word sequences in text

###  Andrew Burt

This program counts three word sequences (trigrams, a case of n-grams) in an input text, and returns the top 100 ranked sequences with counts.  
  
To run:  
`go run main.go file.txt` to count three word sequences in file.txt  
`go run main.go file1.txt file2.txt file3.txt` to count three word sequences in file1.txt, file2.txt, and file3.txt  
`cat file.txt|go run main.go` also counts three word sequences in file.txt  
If no file is sent via command-line argument or a pipe, a message is returned explaining what the program does, and how to send it text files.  
If a file has fewer than 100 three word sequences, it returns however many sequences are available.  

There are tests that can be run with `go test` for both the pre-processing function and the trigram/sequence detection function.  
  
Possible problem: The sample counts mentioned in the e-mail (the sperm whale - 85, the white whale - 71, of the whale - 67), don't all match my results. I did read rules several times in the e-mail, and tried interpreting them several ways (with all non-letter input ignored/removed including in-word punctuation like hyphens and apostrophes, with only between-word punctuation ignored, etc.), but the closest I got was my first version, which is in the program here. The top three it returns are (the sperm whale - 86, of the whale - 78, the white whale - 71).   I did a manual count of a sequence or two, which seem to match the numbers I came up with, so I must have misunderstood one of the rules, or perhaps I'm using a different version of the moby-dick text (the gutenberg project updates and edits its files occasionally, and as the link went to the HTML version, and I had to browse to the text version, I may have gotten a slightly different version - I included the version I used in the repo). 

If I had more time, I'd have written the app with channels to concurrently process the text files, to handle the issue described in the extra credit "handle 1000 moby dick texts at once while performant" case. If memory consumption was a strong concern, I might also write it so it would only read in a smaller piece of the file at a time, or a line at a time, processing each as it came in and letting it go from memory after processing. It would require a little extra work to retain trigrams that span multiple lines or file-chunks, retaining the last few words of each chunk or line while processing the following one. 
