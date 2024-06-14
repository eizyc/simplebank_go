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