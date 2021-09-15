package main

import (
	"strconv"
	"testing"
)

var singleHashTest = []struct {
	in  int
	out string
}{
	{0, "4108050209~502633748"},
	{1, "2212294583~709660146"},
}

func TestSingleHash(t *testing.T) {
	for _, tt := range singleHashTest {
		t.Run(strconv.Itoa(tt.in), func(t *testing.T) {
			in := make(chan interface{}, 1)
			in <- tt.in
			close(in)

			out := make(chan interface{}, 1)
			SingleHash(in, out)
			result := <-out
			if result != tt.out {
				t.Errorf("got %s instead of %s", result, tt.out)
			}
		})
	}
}

var multiHashTest = []struct {
	in  string
	out string
}{
	{"4108050209~502633748", "29568666068035183841425683795340791879727309630931025356555"},
	{"2212294583~709660146", "4958044192186797981418233587017209679042592862002427381542"},
}

func TestMultiHash(t *testing.T) {
	for _, tt := range multiHashTest {
		t.Run(tt.in, func(t *testing.T) {
			in := make(chan interface{}, 1)
			in <- tt.in
			close(in)

			out := make(chan interface{}, 1)
			MultiHash(in, out)
			result := <-out

			if result != tt.out {
				t.Errorf("got %s instead of %s", result, tt.out)
			}
		})
	}
}
