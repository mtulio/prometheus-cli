package main

import "fmt"

func showResults(rtype string, q *Query, total int) {

	fmt.Printf("\n=====\n")
	fmt.Printf("Query results for command [%s]:\n", rtype)
	fmt.Println("\t Server\t\t : ", *cli.promURL)
	if len(q.matches) > 0 {
		fmt.Println("\t Matches\t : ", q.matches)
	}
	if len(q.query) > 0 {
		fmt.Println("\t Query\t\t : ", q.query)
	}
	if q.tStart.Unix() > 0 {
		fmt.Println("\t Time Start\t : ", q.tStart)
	}
	if q.tStart.Unix() > 0 {
		fmt.Println("\t Time End\t : ", q.tEnd)
	}
	fmt.Println("\t Total Results\t : ", total)
	fmt.Printf("=====\n")

}
