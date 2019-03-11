# prometheus-cli

Prometheus Command Line Interface to interact with [Prometheus API](https://prometheus.io/docs/prometheus/latest/querying/api/).


## USAGE

[Series Documentation reference](https://prometheus.io/docs/prometheus/latest/querying/api/#querying-metadata) to understand the API usage.

### SERIES Query

* Show all series for an query

```bash
$ PROMETHEUS_API_URL=http://prometheus.internal:9090 \
    ./bin/prometheus-cli \
        -query=go_info \
        -start="2018-06-01T20:10:30.781Z" \
        -end="2018-12-31T00:00:00.000Z" \
        -verbose \
        series
```

### SERIES Removal

* Remove series

```bash
$ PROMETHEUS_API_URL=http://prometheus.internal:9090 \
    ./bin/prometheus-cli \
        -match=go_info \
        -time.start="2018-09-02T00:00:00.000Z" \
        -time.end="2018-09-02T23:59:59.999Z" \
        -verbose \
        -force \
        delete
```

* Delete all series on time range

> WARNING: it will lookup by `__name__` and delete all series from an given range

```bash
$ PROMETHEUS_API_URL=http://prometheus.internal:9090 \
    ./bin/prometheus-cli \
        -time.start="2018-09-02T00:00:00.000Z" \
        -time.end="2018-09-02T23:59:59.999Z" \
        -verbose \
        -query='{__name__}' \
        delete-batch
```

### LABELS Query

* List Label - metric name

```bash
$ PROMETHEUS_API_URL=http://prometheus.internal:9090 \
    ./bin/prometheus-cli \
        -verbose \
        -query=__name__ \
        label
```

* List Label - instance

```bash
$ PROMETHEUS_API_URL=http://prometheus.internal:9090 \
    ./bin/prometheus-cli \
        -query=instance \
        -verbose \
        label
```


## CONTRIBUTING

Feel free to contributing opening an Issue or an PR. =)
