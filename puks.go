package DatZ5006

import (
	"bufio"
	"cmp"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type graph struct {
	length int
	edges  []edge
}

type edge struct {
	A int
	B int
	w int
}

type unionFind struct {
	parent []int
}

func (uf *unionFind) find(i int) int {
	if uf.parent[i] == i {
		return i
	}
	uf.parent[i] = uf.find(uf.parent[i])
	return uf.parent[i]
}

func (uf *unionFind) union(a, b int) bool {
	rootA := uf.find(a)
	rootB := uf.find(b)
	if rootA != rootB {
		uf.parent[rootA] = rootB
		return true
	}
	return false
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	for i := range n {
		parent[i] = i
	}
	return &unionFind{parent: parent}
}

func Process(g *graph) (int, []edge) {
	uf := newUnionFind(g.length)

	slices.SortFunc(g.edges, func(a, b edge) int {
		return cmp.Compare(b.w, a.w)
	})

	var feedback []edge
	w := 0
	for _, item := range g.edges {
		if !uf.union(item.A-1, item.B-1) {
			w += item.w
			feedback = append(feedback, item)
		}
	}

	return w, feedback
}

func Parse(reader io.Reader) (*graph, error) {
	result := graph{}
	i := 0
	var a, b int
	sc := bufio.NewScanner(reader)
	for sc.Scan() {
		line := sc.Text()                // Read file line by line
		items := strings.FieldsSeq(line) // Tokenize integers by whitespaces
		for item := range items {
			if value, err := strconv.Atoi(item); err != nil {
				return nil, errors.Errorf("Can not parse %s", item)
			} else {
				switch i % 3 {
				case 0:
					if i == 0 {
						result.length = value
					} else {
						result.edges = append(result.edges, edge{A: a, B: b, w: value})
					}
				case 1:
					a = value
				case 2:
					b = value
				}
				i++
			}
		}
	}
	if err := sc.Err(); err != nil {
		return nil, errors.Wrap(err, "File scan")
	}
	if i%3 != 1 {
		return nil, errors.Errorf("Wrong number of integers (%d) for graph", i)
	}
	return &result, nil
}
