GO-JET TUTORIAL.

+ Create new go project
+ Install 'go-jet'
+ Generate models

## Install go-jet

### Install binaries

```shell
go install github.com/go-jet/jet/v2/cmd/jet@latest
```

### Install module

```shell
$ go get -u github.com/go-jet/jet/v2
```

## Generate models
```shell
jet -source=postgres -dsn="user=postgres password=postgres host=localhost port=15432 dbname=pet_store sslmode=disable" -schema=public -path=./internal/generated
```