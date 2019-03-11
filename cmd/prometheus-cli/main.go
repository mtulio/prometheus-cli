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

func die(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr, "")
	os.Exit(1)
}

func init() {
	cli.promURL = flag.String("server", getEnvURL(), "Prometheus URL")
	cli.match = flag.String("match", getEnvMatch(), "Prometheus Match sepparated by comma.")
	cli.tStartStr = flag.String("time.start", "", "Prometheus Query Start.")
	cli.tEndStr = flag.String("time.end", "", "Prometheus Query End.")
	cli.verbose = flag.Bool("verbose", false, "Verbose the result.")
	cli.force = flag.Bool("force", false, "Force commad. For delete, remove from disk (Clean Tombstones).")
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
		getSeries()
	case "delete":
		deleteSeries()
	default:
		die("Unknown command '%s'", flag.Arg(0))
	}

}
