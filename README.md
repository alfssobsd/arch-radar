# Development

## Setup environment
### Start database
```
docker-compose start database
```
### Install GRPC Client
https://github.com/uw-labs/bloomrpc


## Generate grpc servers from proto
```
docker-compose run protobufs
```

## Generate models
```
genna model -c postgres://archradar_user:archradar_pass@localhost:5432/archradar_db?sslmode=disable -o backend-service/dataproviders/pg_provider/model/model.go -t public.* -f -g 9
```
