# logprocess
## logprocess is a tool to validate, store and get reports from logs.

## Dependencies
- Docker
- Docker Compose
- Go v1.16

## Build

```sh
make
```

## Execute
### Set environment variables
```sh
set -a
. ./env.env
set +a
```

### Run
#### Insert the logs:
```sh
./bin/processor -f 0 < logs.txt
```
#### Get a consumer report:
```sh
./bin/processor -f 1 <consumerID>
```

#### Get a service report:
```sh
./bin/processor -f 2 <serviceID>
```

#### Get the average services latencies report:
```sh
./bin/processor -f 3
```
## Clean (remove the binary and the container)
```sh
make clean
```
## Remove database data
```sh
make destroy
```
