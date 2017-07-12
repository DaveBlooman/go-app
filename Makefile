export APP_ENV = development
export GEN_DBUSER = postgres
export GEN_DBPASS = postgres
export GEN_DBSSLMODE = disable
export GEN_DBNAME = test
export GEN_DBHOST = localhost
export GEN_DBPORT = 5432
export PGPASSWORD = postgres

tests:
	go test `glide novendor`

install:
	glide install

start:
	go run main.go
