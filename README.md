# Gin Opentelemetry Demo

## Run

Up service

```shell
# docker compose
make up

# migrate db
make pg-up

# run go app
make app
```

Down service

```shell
make down
# or
# make down-v
```

## Monitoring

Open Uptrace UI at [localhost:14318](http://localhost:14318)

## Note

- [OpenTelemetry](https://uptrace.dev/opentelemetry/) is an open source observability framework hosted by Cloud Native Computing Foundation.
- [Uptrace](https://uptrace.dev/get/opentelemetry-apm.html) is OpenTelemetry APM (Application Performance Monitoring).
- [Uptrace Go](https://uptrace.dev/get/opentelemetry-go.html) is a thin wrapper over [opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go) that configures OpenTelemetry SDK to export data to Uptrace.
- [Gin OpenTelemetry instrumentation](https://uptrace.dev/get/instrument/opentelemetry-gin.html) allows developers to easily add observability to their Gin applications, providing insights into application performance, behavior, and usage patterns.
- [Uptrace: Docker Example](https://github.com/uptrace/uptrace/tree/master/example/docker)
- [Tracing](https://uptrace.dev/opentelemetry/distributed-tracing.html) and [Span](https://uptrace.dev/opentelemetry/distributed-tracing.html#spans).
