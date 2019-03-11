package main

import (
	"fmt"
	"os"

	"golang.org/x/net/context"
)

func GetSeries() {
	// defer die("GetSeries()", nil)

	if !validateArgsMatch() {
		fmt.Println("Invalid query arguments")
		usage()
		os.Exit(1)
	}

	q := Query{
		matches: *cli.matchs,
		tStart:  *cli.tStartTime,
		tEnd:    *cli.tEndTime,
	}

	apiS, err := promAPI.Series(context.Background(),
		q.matches, q.tStart, q.tEnd)

	if err != nil {
		fmt.Println(err)
	}

	if *cli.verbose {
		for i, v := range apiS {
			fmt.Println(i, v)
		}
	}

	showResults("series", &q, len(apiS))
}
