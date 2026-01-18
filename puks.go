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
	var err error
	stream := make(chan int)
	go func() {
		defer close(stream)
		err = tokenize(reader, stream)
	}()
	if err != nil {
		return nil, err
	}

	result := graph{}
	i := 0
	var a, b int
	for item := range stream {
		switch i % 3 {
		case 0:
			if i == 0 {
				result.length = item
			} else {
				result.edges = append(result.edges, edge{a: a, b: b, w: item})
			}
		case 1:
			a = item
		case 2:
			b = item
		}
		i++
	}
	if i%3 != 1 {
		return nil, errors.Errorf("Wrong number of elements for graph %d", i)
	}
	return &result, nil
}

func tokenize(reader io.Reader, stream chan int) error {
	sc := bufio.NewScanner(reader)
	for sc.Scan() { // Tokenize every line over whitespaces
		line := sc.Text()
		items := strings.FieldsSeq(line)
		for item := range items {
			if value, err := strconv.Atoi(item); err != nil {
				return errors.Errorf("Can not parse %s", item)
			} else {
				stream <- value
			}
		}
	}
	if err := sc.Err(); err != nil {
		return errors.Wrap(err, "File scan")
	}

	return nil
}
