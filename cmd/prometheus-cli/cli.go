package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func getEnvURL() string {
	return os.Getenv("PROMETHEUS_API_URL")
}

func getEnvMatch() string {
	return os.Getenv("PROMETHEUS_API_MATCH")
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	// fmt.Fprintf(os.Stderr, "\t%s [flags] query <expression>\n", os.Args[0])
	// fmt.Fprintf(os.Stderr, "\t%s [flags] query_range <expression> <end_timestamp> <range_seconds> [<step_seconds>]\n", os.Args[0])
	// fmt.Fprintf(os.Stderr, "\t%s [flags] metrics\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%s [flags] series\n", os.Args[0])
	fmt.Printf("\nFlags:\n")
	flag.PrintDefaults()
}

func validadeArgsTimeRange() bool {
	if *cli.tStartStr == "" {
		emsg := fmt.Errorf("Missing -start")
		fmt.Println(emsg)
		return false
	} else {
		t, err1 := time.Parse(time.RFC3339Nano, *cli.tStartStr)
		if err1 != nil {
			emsg := fmt.Errorf("Unable to convert -start")
			fmt.Println(emsg)
			return false
		}
		cli.tStartTime = &t
	}

	if *cli.tEndStr == "" {
		emsg := fmt.Errorf("Missing -end")
		fmt.Println(emsg)
		return false
	} else {
		t, err2 := time.Parse(time.RFC3339Nano, *cli.tEndStr)
		if err2 != nil {
			emsg := fmt.Errorf("Unable to convert -end")
			fmt.Println(emsg)
			return false
		}
		cli.tEndTime = &t
	}

	return true
}

func validateArgsMatch() bool {

	if *cli.match == "" {
		emsg := fmt.Errorf("Missing -match or PROMETHEUS_API_MATCH env")
		fmt.Println(emsg)
		return false
	}
	s := strings.Split(*cli.match, ",")
	cli.matchs = &s

	if ok := validadeArgsTimeRange(); !ok {
		return false
	}

	return true
}
