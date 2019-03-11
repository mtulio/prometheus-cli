package main

import "time"

type Query struct {
	matches []string
	tStart  time.Time
	tEnd    time.Time
}

type CliOptions struct {
	promURL    *string
	query      *string
	queryList  *[]string
	tStartStr  *string
	tStartTime *time.Time
	tEndStr    *string
	tEndTime   *time.Time
	verbose    *bool
}
