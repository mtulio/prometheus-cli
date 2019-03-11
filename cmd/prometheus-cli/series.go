package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/context"
)

func getSeries() {
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

	series, err := promAPI.Series(context.Background(),
		q.matches, q.tStart, q.tEnd)

	if err != nil {
		fmt.Println(err)
		return
	}

	if *cli.verbose {
		for i, v := range series {
			fmt.Println(i, v)
		}
	}

	showResults("series", &q, len(series))
}

func deleteSeries() {

	getSeries()

	fmt.Println("The above series will permantinent deleted.")
	fmt.Printf("Do you want to continue? ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	pass := true
	if input == "y" || input == "yes" {
		pass = true
	} else {
		pass = false
	}

	if !pass {
		return
	}

	q := Query{
		matches: *cli.matchs,
		tStart:  *cli.tStartTime,
		tEnd:    *cli.tEndTime,
	}

	if err := promAPI.DeleteSeries(context.Background(),
		q.matches, q.tStart, q.tEnd); err != nil {
		fmt.Println(err)
		return
	}

	if *cli.force {
		fmt.Println("Cleaning Tombestones.")
		if err := promAPI.CleanTombstones(context.Background()); err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Success!")
}
