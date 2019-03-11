package main

import "time"

type Query struct {
	matches []string
	query   string
	tStart  time.Time
	tEnd    time.Time
}

type CliOptions struct {
	promURL    *string
	query      *string
	match      *string
	matchs     *[]string
	tStartStr  *string
	tStartTime *time.Time
	tEndStr    *string
	tEndTime   *time.Time
	verbose    *bool
	force      *bool
}
