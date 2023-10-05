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


### UNIT 1: Simple queries 

+ Select pets. 
+ Select a pet by ID.
+ Select only available pets. 
+ Insert new pet
+ Update pet status


### UNIT 2: Joins and unions

+ Select pets with categories
+ Select pet's tags and category

### Unit 3: Insert and Update rows

+ Create new order
+ Update order status
+ Batch insert
+ Deleting orders
+ 

### Unit4: UPSERT queries
+ Insert pet ON Conflict UPDATE