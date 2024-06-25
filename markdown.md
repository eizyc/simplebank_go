### Tools/Package
[dbdiagram](https://dbdiagram.io/d/Simple-Bank-666750ce6bc9d447b153e02f)
[golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
[sqlc](https://github.com/sqlc-dev/sqlc)


### Use makefile to creare postgresql container and database
``` shell
make postgres
make createdb
```

#### Use migrate to create databse schema
``` shell
migrate -path db/migration/ -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
```
#### install golang-migrate
```
brew install golang-migrate
```

#### How to make a new version database migration aka. the database need to change some schema
```
make new_migration
```

#### How to build the docker and build docker network includes postgresql and go container
```
\\ build image, in the Dockerfile folder
docker build -t simplebank:last .
\\ create docker network
docker network create bank-network
\\ connect network to postgresql container
docker network connect bank-network postgres16
\\ container 
docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres:5432/simple_bank?sslmode=disable" simplebank:last 
```

#### How to run docker-compose
```
docker compose up
```

#### Modiy file permissions for start.sh wait-for.sh for Control startup and shutdown order in Compose
To make the api container run after by the postgres container, we have to use [`depends_on`](https://docs.docker.com/compose/startup-order/) and with `condition: service_healthy`
By Legacy versions, besides the depends_on field, we need some script like [wait-for](https://github.com/mrako/wait-for), new version docker compose supported inside. Anyway, If we want to use some scipts, don't miss change scipts permissions.
```
chomd +x fileName
```


#### Create a new db migration
```
make new_migration name=<migration_name>
```

#### Create dbdocs and db_schema
##### dbdocs need
```
npm i dbdocs -g
```
and `make db_docs` 
##### db_schema need
```
npm install -g @dbml/cli
```
and `make db_schema` 

#### GRPC
##### Install
[Quick start](https://grpc.io/docs/languages/go/quickstart/)
```
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
##### gRPC clients
[gRPC clients evans](https://github.com/ktr0731/evans)
```
brew tap ktr0731/evans
brew install evans
```
Request
```
evans --host localhost --port 9090 -r repl
```
