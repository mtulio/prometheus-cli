package main

import "fmt"

func showResults(rtype string, q *Query, total int) {

	fmt.Printf("\n=====\n")
	fmt.Printf("Query results for command [%s]:\n", rtype)
	fmt.Println("\t Server\t\t : ", *cli.promURL)
	fmt.Println("\t Matches\t : ", q.matches)
	fmt.Println("\t Time Start\t : ", q.tStart)
	fmt.Println("\t Time End\t : ", q.tEnd)
	fmt.Println("\t Total Results\t : ", total)
	fmt.Printf("=====\n")

}
