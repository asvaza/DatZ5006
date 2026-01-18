package main

import (
	"bufio"
	"fmt"
	"os"

	datz "github.com/asvaza/DatZ5006"
)

func main() {
	if len(os.Args) > 1 {
		file := os.Args[1]

		source, err := os.Open(file)
		if err != nil {
			fmt.Printf("Error - %s\n", err.Error())
			return
		}
		defer source.Close()

		g, err := datz.Parse(source)
		if err != nil {
			fmt.Printf("Error - %s\n", err.Error())
			return
		}

		w, edges := datz.Process(g)

		f, err := os.Create("puks.txt")
		if err != nil {
			fmt.Printf("Error - %s\n", err.Error())
			return
		}
		defer f.Close()

		out := bufio.NewWriter(f)
		fmt.Fprintf(out, "%d\t%d\n", len(edges), w)
		for i, item := range edges {
			fmt.Fprintf(out, "%d\t%d\t\t", item.A, item.B)
			if i%5 == 4 {
				fmt.Fprintf(out, "\n")
			}
		}
		fmt.Fprintf(out, "\n")
		out.Flush()

	} else {
		fmt.Println("Error - file argument required")
		fmt.Println("Example")
		fmt.Println("puks.exe sample_input_2025_1.txt")
	}
}
