package main

import (
	"bufio"
	"cmp"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	// var iFlag = flag.String("i", "", "Input a file name")
	flag.Parse()
	var filename = flag.Arg(0)

	// note:
	// both os.Stdin and the return of os.Open are io.Reader
	var src io.Reader

	if flag.NArg() == 0 {
		fmt.Println("No filename provided. Please provide input now:")

		src = os.Stdin
	} else {
		// note:
		// read a file and return []byte + error
		// data, err := os.ReadFile(filename)
		// if err = nil {
		// 	log.Fatal(err)
		// }

		// open a file
		data, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		src = data
		defer data.Close() // data closed at the end of main
	}

	wordCountsMap := make(map[string]int)

	// create a scanner to loop through each line
	scanner := bufio.NewScanner(src)
	for scanner.Scan() { // Scan() returns the string up to \n
		words := strings.Split(scanner.Text(), " ")

		for _, word := range words {
			if word == "" {
				continue
			}
			wordCountsMap[word]++
			// fmt.Printf("Found %d %s\n", wordCounts[word], word)
		}
	}
	fmt.Println("finished counting...")

	// check for any errors that may have occurred
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err) // TODO what is diff between Fprintln and Println
	}

	// count results
	countCmp = func(a, b int) int {
		return cmp.Compare(a, b)
	}

	wordCountsSlice := make([]struct {
		word  string
		count int
	}, len(wordCountsMap))
	
	for word, count := range wordCountsMap {
		wordCountsSlice = append(wordCountsSlice, {word, count})
	}

	slices.Sort(wordCountsSlice)

	fmt.Println("exiting...")
	os.Exit(1)
}
