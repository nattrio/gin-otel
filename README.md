# Gin Opentelemetry Demo

[Tutorial](https://github.com/uptrace/uptrace/tree/master/example/docker)

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

Open Uptrace UI at [:14318](http://localhost:14318)
