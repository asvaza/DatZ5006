package DatZ5006

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseString(t *testing.T) {
	AssertParseText(t, 1, 1, edge{2, 3, 4}, edge{2, 3, 4}, "1 2 3 4", "")
	AssertParseText(t, 1, 1, edge{2, 3, -4}, edge{2, 3, -4}, "1\n2\n3\n-4", "")
	AssertParseText(t, 1, 1, edge{2, 3, 4}, edge{2, 3, 4}, "1\t2\n3\r4", "")
	AssertParseText(t, 1, 1, edge{2, 3, 4}, edge{2, 3, 4}, "1\r\n2\n\r3 4", "")
	AssertParseText(t, 0, 0, edge{}, edge{}, "1 2 a 4", "Can not parse a")
	AssertParseText(t, 0, 0, edge{}, edge{}, "0", "")
	AssertParseText(t, 1, 1, edge{}, edge{}, "1 2", "Wrong number of integers (2) for graph")
	AssertParseText(t, 1, 1, edge{}, edge{}, "1 2 3", "Wrong number of integers (3) for graph")
}

func TestParseFile(t *testing.T) {
	AssertParseFile(t, 10, 40, edge{1, 2, 69}, edge{9, 10, -90}, "test/sample_input_2025_1.txt")
	AssertParseFile(t, 100, 600, edge{1, 2, 15}, edge{88, 100, -52}, "test/sample_input_2025_2.txt")
	AssertParseFile(t, 100, 600, edge{1, 2, 15}, edge{88, 100, -52}, "test/sample_input_2025_2a.txt")
	AssertParseFile(t, 400, 3000, edge{1, 2, 5}, edge{352, 400, -88}, "test/sample_input_2025_3.txt")
	AssertParseFile(t, 400, 3000, edge{1, 2, 5}, edge{352, 400, -88}, "test/sample_input_2025_3a.txt")
}

func TestProcess10(t *testing.T) {
	AssertProcess(t, 682, 40-9, "test/sample_input_2025_1.txt")
	AssertProcess(t, 2180, 600-99, "test/sample_input_2025_2.txt")
	AssertProcess(t, 19674, 3000-399, "test/sample_input_2025_3.txt")
	AssertProcess(t, 2176, 600-99, "test/sample_input_2025_2a.txt")
	AssertProcess(t, 19663, 3000-399, "test/sample_input_2025_3a.txt")
}

func AssertProcess(t *testing.T, w, k int, file string) {
	source, err := os.Open(file)
	if err != nil {
		t.Error(err)
	}
	defer source.Close()

	g, err := Parse(source)
	if err != nil {
		t.Error(err)
		return
	}

	w, edges := Process(g)
	assert.Equal(t, w, w)
	assert.Equal(t, k, len(edges))
}

func AssertParseFile(t *testing.T, vertices, edges int, first, last edge, file string) {
	source, err := os.Open(file)
	if err != nil {
		t.Error(err)
	}
	defer source.Close()

	graph, err := Parse(source)
	if err != nil {
		t.Error(err)
		return
	}
	require.NotNil(t, graph)
	assert.Equal(t, vertices, graph.length)
	assert.Equal(t, edges, len(graph.edges))
	assert.Equal(t, first, graph.edges[0])
	assert.Equal(t, last, graph.edges[len(graph.edges)-1])
}

func AssertParseText(t *testing.T, vertices, edges int, first, last edge, text string, failure string) {
	graph, err := Parse(strings.NewReader(text))
	if err != nil {
		if failure == "" {
			t.Error(err)
		} else {
			assert.Equal(t, failure, err.Error())
		}
		return
	}
	require.NotNil(t, graph)
	assert.Equal(t, vertices, graph.length)
	assert.Equal(t, edges, len(graph.edges))
	if len(graph.edges) > 0 {
		assert.Equal(t, first, graph.edges[0])
		assert.Equal(t, last, graph.edges[len(graph.edges)-1])
	}
}
