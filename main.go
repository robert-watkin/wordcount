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

type wordCount struct {
	word  string
	count int
}

func main() {
	var topFlag = flag.Int("top", 0, "Specify the top number of records to return")
	var minFlag = flag.Int("min", 0, "Specify the minimum number of records to return")

	caseSensitiveFlag := flag.Bool("case-sensitive", false, "Case sensitive flag")
	flag.BoolVar(caseSensitiveFlag, "c", false, "Case sensitive flag (shorthand)")

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
		text := scanner.Text()
		if !*caseSensitiveFlag {
			text = strings.ToLower(text)
		}

		words := strings.Split(text, " ")

		for _, word := range words {
			if word == "" {
				continue
			}
			wordCountsMap[word]++
			// fmt.Printf("Found %d %s\n", wordCounts[word], word)
		}
	}

	// check for any errors that may have occurred
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err) // TODO what is diff between Fprintln and Println
	}

	// order based on count
	// create comparison function
	countCmp := func(a, b wordCount) int {
		res := cmp.Compare(b.count, a.count)
		if res == 0 {
			return cmp.Compare(b.word, a.word)
		} else {
			return cmp.Compare(b.count, a.count)
		}
	}

	// copy the map into a slice
	wordCountsSlice := make([]wordCount, 0, len(wordCountsMap))
	for word, count := range wordCountsMap {
		wordCountsSlice = append(wordCountsSlice, wordCount{word, count})
	}

	// sort the slice based on the comparison function
	slices.SortFunc(wordCountsSlice, countCmp)

	if *topFlag != 0 {
		wordCountsSlice = wordCountsSlice[:*topFlag]
	}

	if *minFlag != 0 {
		wordCountsSlice = wordCountsSlice[0:*minFlag]
	}

	for _, x := range wordCountsSlice {
		fmt.Printf("%d %s\n", x.count, x.word)
	}
	os.Exit(1)
}
