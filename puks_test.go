package DatZ5006

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizerString(t *testing.T) {
	AssertTokenizerText(t, []int{1, 2, 3, 4}, "1 2 3 4", "")
	AssertTokenizerText(t, []int{1, 2, 3, -4}, "1\n2\n3\n-4", "")
	AssertTokenizerText(t, []int{1, 2, 3, 4}, "1\t2\n3\r4", "")
	AssertTokenizerText(t, []int{1, 2, 3, 4}, "1\r\n2\n\r3 4", "")
	AssertTokenizerText(t, nil, "1 2 a 4", "Can not parse a")
}

func TestTokenizerFile(t *testing.T) {
	AssertTokenizerFile(t, 121, []int{10, 1, 2, 69}, []int{6, 9, 10, -90}, "test/sample_input_2025_1.txt")
	AssertTokenizerFile(t, 1801, []int{100, 1, 2, 15}, []int{10, 88, 100, -52}, "test/sample_input_2025_2.txt")
	AssertTokenizerFile(t, 1801, []int{100, 1, 2, 15}, []int{10, 88, 100, -52}, "test/sample_input_2025_2a.txt")
	AssertTokenizerFile(t, 9001, []int{400, 1, 2, 5}, []int{29, 352, 400, -88}, "test/sample_input_2025_3.txt")
	AssertTokenizerFile(t, 9001, []int{400, 1, 2, 5}, []int{29, 352, 400, -88}, "test/sample_input_2025_3a.txt")
}

func TestParseFile(t *testing.T) {
	source, err := os.Open("test/sample_input_2025_1.txt")
	if err != nil {
		t.Error(err)
	}
	defer source.Close()
	graph, err := parse(source)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 10, graph.length)
	assert.Equal(t, 40, len(graph.edges))
	assert.Equal(t, edge{1, 2, 69}, graph.edges[0])
	assert.Equal(t, edge{9, 10, -90}, graph.edges[len(graph.edges)-1])
}

func AssertTokenizerFile(t *testing.T, length int, start []int, end []int, file string) {
	source, err := os.Open(file)
	if err != nil {
		t.Error(err)
	}
	defer source.Close()

	stream := make(chan int)
	go func() {
		if err := tokenize(source, stream); err != nil {
			t.Error(err)
		}
		close(stream)
	}()

	actual := []int{}
	for item := range stream {
		actual = append(actual, item)
	}
	assert.Equal(t, length, len(actual))
	assert.Equal(t, start, actual[:4])
	assert.Equal(t, end, actual[len(actual)-4:])
}

func AssertTokenizerText(t *testing.T, expected []int, text string, failure string) {
	actual := []int{}
	stream := make(chan int)
	go func() {
		for item := range stream {
			actual = append(actual, item)
		}
		assert.Equal(t, expected, actual)
	}()

	if err := tokenize(strings.NewReader(text), stream); err != nil {
		if failure == "" {
			t.Error(err)
		} else {
			assert.Equal(t, failure, err.Error())
		}
	}
	close(stream)
}
