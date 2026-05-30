package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	// count tests held in array of structs
	countTests := []struct {
		name          string
		in            string
		want          map[string]int
		caseSensitive bool
	}{
		// test cases defined
		{"test 1", "foo bar foo", map[string]int{"foo": 2, "bar": 1}, false},
		{"test 2", "the The the", map[string]int{"the": 3}, false},
		{"test 3", "the The the", map[string]int{"the": 2, "The": 1}, true},
	}

	// loop through the test cases
	for _, tt := range countTests {
		// run the test
		t.Run(tt.in, func(t *testing.T) {
			// call the Count() function and handle any error returns or unexpected returns
			got, err := Count(strings.NewReader(tt.in), tt.caseSensitive)
			if err != nil {
				t.Fatalf("count() returned error: %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}
