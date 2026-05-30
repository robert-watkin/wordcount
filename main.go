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
	var reader io.Reader

	if flag.NArg() == 0 {
		reader = os.Stdin
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
		reader = data
		defer func() {
			err := data.Close() // data closed at the end of main
			if err != nil {
				log.Fatalf("Failed to close data input: %v", err)
			}
		}()
	}

	wordCountsMap, err := Count(reader, *caseSensitiveFlag)
	// call to count with error handling
	if err != nil {
		log.Fatalf("failed to count input : %v", err)
		// Or: fmt.Fprintln(os.Stderr, "reading standard input:", err); os.Exit(1)
	}

	// order based on count
	// create comparison function
	countCmp := func(a, b wordCount) int {
		return cmp.Or(cmp.Compare(b.count, a.count), cmp.Compare(a.word, b.word))
	}

	// copy the map into a slice
	wordCountsSlice := make([]wordCount, 0, len(wordCountsMap))
	for word, count := range wordCountsMap {
		if *minFlag <= count {
			wordCountsSlice = append(wordCountsSlice, wordCount{word, count})
		}
	}

	// sort the slice based on the comparison function
	slices.SortFunc(wordCountsSlice, countCmp)

	if *topFlag != 0 && *topFlag <= len(wordCountsSlice) {
		wordCountsSlice = wordCountsSlice[:*topFlag]
	}

	for _, x := range wordCountsSlice {
		fmt.Printf("%d %s\n", x.count, x.word)
	}
}

func Count(r io.Reader, caseSensitiveFlag bool) (map[string]int, error) {
	wordCountsMap := make(map[string]int)

	// create a scanner to loop through each line
	scanner := bufio.NewScanner(r)
	for scanner.Scan() { // Scan() returns the string up to \n
		text := scanner.Text()
		if !caseSensitiveFlag {
			text = strings.ToLower(text)
		}

		words := strings.Fields(text)

		for _, word := range words {
			wordCountsMap[word]++
		}
	}

	return wordCountsMap, scanner.Err()
}
