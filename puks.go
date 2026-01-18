package DatZ5006

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type graph struct {
	length int
	edges  []edge
}

type edge struct {
	a int
	b int
	w int
}

func parse(reader io.Reader) (*graph, error) {
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
						result.edges = append(result.edges, edge{a: a, b: b, w: value})
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
