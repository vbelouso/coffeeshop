# Just for fun coffee shop project

## Getting Started

### Prerequisites

`Go` version with module support.

`docker-compose` must be [installed](https://docs.docker.com/compose/install/)

### Starting project

Start `db` service with the following command:

```shell
docker-compose up -d
```

Start API server

```shell
go run .
```

Check API server

```shell
curl localhost:8080/customers/1
```

```shell
curl localhost:8080/customers
```

```shell
curl localhost:8080/orders/1
```

```shell
curl localhost:8080/orders
```

## Notes

Dig into UUID usage instead of Integer for PGSQL
DB Transactions?
ConnectionPool

## Improvements

Reduce code repetition for errors(make constants/vars)
Save only structs inside models.go