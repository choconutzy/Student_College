# Mahasiswa API
This app build using gin, and MySQL
## Database Migrations
Before run the migration, ensure that you have been installed [golang-migrate cli](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md) and created database in your local
```
username:password@protocol(address)/dbname?param=value
```
### How to Run Migration Up
```
migrate -path src/database/migration/ -database "mysql://username:password@tcp(localhost:3306)/database_name?sslmode=disable" -verbose up
```
### How to Run Migration Down
```
migrate -path src/database/migration/ -database "mysql://username:password@tcp(localhost:3306)/database_name?sslmode=disable" -verbose down
```

## How to Run 

```
nodemon --exec go run main.go --signal SIGTERM
```

## How to Run with nodemon

```
nodemon --exec go run main.go --signal SIGTERM
```