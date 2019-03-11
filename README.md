# prometheus-cli

Prometheus Command Line Interface to interact with [Prometheus API](https://prometheus.io/docs/prometheus/latest/querying/api/).


## Series

[Series Documentation reference](https://prometheus.io/docs/prometheus/latest/querying/api/#querying-metadata)

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

* Removing series

```
PROMETHEUS_API_URL=http://prometheus.internal:9090 \
    ./bin/prometheus-cli \
        -match=go_info \
        -time.start="2018-09-02T00:00:00.000Z" \
        -time.end="2018-09-02T23:59:59.999Z" z
        -verbose \
        delete
0 {__name__="go_info", instance="redis-master-01", job="node", version="go1.9.1"}
1 {__name__="go_info", instance="redis-master-02", job="node", version="go1.9.1"}
2 {__name__="go_info", instance="redis-master-03", job="node", version="go1.9.1"}
3 {__name__="go_info", instance="redis-master-04", job="node", version="go1.9.1"}
4 {__name__="go_info", instance="redis-master-05", job="node", version="go1.9.1"}
5 {__name__="go_info", instance="redis-slave1-01", job="node", version="go1.9.1"}
6 {__name__="go_info", instance="redis-slave1-02", job="node", version="go1.9.1"}
7 {__name__="go_info", instance="redis-slave1-03", job="node", version="go1.9.1"}
8 {__name__="go_info", instance="redis-slave1-04", job="node", version="go1.9.1"}
9 {__name__="go_info", instance="redis-slave1-05", job="node", version="go1.9.1"}
10 {__name__="go_info", instance="redis-slave2-01", job="node", version="go1.9.1"}
11 {__name__="go_info", instance="redis-slave2-02", job="node", version="go1.9.1"}
12 {__name__="go_info", instance="redis-slave2-03", job="node", version="go1.9.1"}
13 {__name__="go_info", instance="redis-slave2-04", job="node", version="go1.9.1"}
14 {__name__="go_info", instance="redis-slave2-05", job="node", version="go1.9.1"}
15 {__name__="go_info", instance="server-ondemand-153", job="node", version="go1.9.1"}
16 {__name__="go_info", instance="server-ondemand-244", job="node", version="go1.9.1"}
18 {__name__="go_info", instance="prometheus", job="node", version="go1.9.1"}
19 {__name__="go_info", instance="prometheus.internal:9091", job="pushgateway", version="go1.10.3"}

=====
Query results for command [series]:
	 Server		 :  http://prometheus.internal:9090
	 Matches	 :  [go_info]
	 Time Start	 :  2018-09-02 00:00:00 +0000 UTC
	 Time End	 :  2018-09-02 23:59:59.999 +0000 UTC
	 Total Results	 :  20
=====
The above series will permantinent deleted.
Do you want to continue? y
Success!
```
