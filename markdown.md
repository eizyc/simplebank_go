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
docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres16:5432/simple_bank?sslmode=disable" simplebank:last 
```