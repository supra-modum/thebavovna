
## Start server and log into Postgres as the user named postgres

```zsh
brew services start postgresql@14
```

```zsh
psql -U postgres
```

## Run migrations:
https://github.com/go-pg/migrations

```
cd migrations/
go run *.go up
```

To create executable file for migrations:

```
cd migrations/
go build -o migrations *.go
```
To run the executable file:

```
./migrations init
./migrations up
```

Read more: https://dev.to/matijakrajnik/migrations-8fn