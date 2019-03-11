package main

import (
	"fmt"
	"os"

	"golang.org/x/net/context"
)

func getLabel() {

	if !validateArgsQuery() {
		fmt.Println("Invalid Query Labels arguments")
		usage()
		os.Exit(1)
	}

	q := Query{
		query: *cli.query,
	}

	labels, err := promAPI.LabelValues(context.Background(),
		q.query)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *cli.verbose {
		for i, v := range labels {
			fmt.Println(i, v)
		}
	}

	showResults("label", &q, len(labels))
}
