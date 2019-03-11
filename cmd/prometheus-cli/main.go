package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

var (
	promAPI v1.API
	cli     CliOptions
)

const (
	totalSecMonth = 60 * 60 * 24 * 30
	totalSecWeek  = 60 * 60 * 24 * 7
	totalSecDay   = 60 * 60 * 24
	totalSecHour  = 60 * 60
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\t%s [flags] query <expression>\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%s [flags] query_range <expression> <end_timestamp> <range_seconds> [<step_seconds>]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%s [flags] metrics\n", os.Args[0])
	fmt.Printf("\nFlags:\n")
	flag.PrintDefaults()
}

func die(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr, "")
	os.Exit(1)
}

func init() {

	cli.promURL = flag.String("prometheus.url", getEnvURL(), "Prometheus URL")
	cli.query = flag.String("query", "", "Prometheus Query.")
	cli.tStartStr = flag.String("start", "", "Prometheus Query Start.")
	cli.tEndStr = flag.String("end", "", "Prometheus Query End.")
	cli.verbose = flag.Bool("verbose", false, "Verbose the result.")
	flag.Usage = usage
	flag.Parse()
}

func main() {

	if cli.promURL == nil {
		emsg := fmt.Errorf("Missing -prometheus.url or PROMETHEUS_API_URL env")
		fmt.Println(emsg)
		usage()
		os.Exit(1)
	}

	cfg := api.Config{
		Address: *cli.promURL,
	}
	client, err1 := api.NewClient(cfg)
	if err1 != nil {
		fmt.Print("Unable to create the client")
		os.Exit(1)
	}

	promAPI = v1.NewAPI(client)

	switch flag.Arg(0) {
	case "series":
		GetSeries()
	default:
		die("Unknown command '%s'", flag.Arg(0))
	}

}
