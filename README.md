# prometheus-cli

Prometheus Command Line Interface to interact with [Prometheus API](https://prometheus.io/docs/prometheus/latest/querying/api/).


## Series

* Show all series for an query

```bash

PROMETHEUS_API_URL=http://prometheus.internal:9090 \
    ./bin/prometheus-cli \
        -query=go_info \
        -start="2018-06-01T20:10:30.781Z" \
        -end="2018-12-31T00:00:00.000Z" \
        -verbose \
        series

```

> Sample answer:

```log
[...]
2864 {__name__="go_info", instance="prometheus.internal:9091", job="pushgateway", version="go1.11.1"}
2865 {__name__="go_info", instance="stress-master", job="node", version="go1.9.1"}
2866 {__name__="go_info", job="node", version="go1.9.1"}
2867 {__name__="go_info", job="node", version="go1.9.2"}

=====
Query results for command [series]:
	 Server		 :  http://prometheus.internal:9090
	 Matches	 :  [go_info]
	 Time Start	 :  2018-06-01 20:10:30.781 +0000 UTC
	 Time End	 :  2018-12-31 00:00:00 +0000 UTC
	 Total Results	 :  2868
=====

```