package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func getEnvURL() string {
	return os.Getenv("PROMETHEUS_API_URL")
}

func getEnvQuery() string {
	return os.Getenv("PROMETHEUS_API_QUERY")
}

func validateArgsQuery() bool {

	if *cli.query == "" {
		emsg := fmt.Errorf("Missing -query or PROMETHEUS_API_QUERY env")
		fmt.Println(emsg)
		return false
	}
	s := strings.Split(*cli.query, ",")
	cli.queryList = &s

	if *cli.tStartStr == "" {
		emsg := fmt.Errorf("Missing -start")
		fmt.Println(emsg)
		return false
	}
	t, err1 := time.Parse(time.RFC3339Nano, *cli.tStartStr)
	if err1 != nil {
		emsg := fmt.Errorf("Unable to convert -start")
		fmt.Println(emsg)
		return false
	}
	cli.tStartTime = &t

	if *cli.tEndStr == "" {
		emsg := fmt.Errorf("Missing -end")
		fmt.Println(emsg)
		return false
	}
	t, err2 := time.Parse(time.RFC3339Nano, *cli.tEndStr)
	if err2 != nil {
		emsg := fmt.Errorf("Unable to convert -end")
		fmt.Println(emsg)
		return false
	}
	cli.tEndTime = &t

	return true
}
