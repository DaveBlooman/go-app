#!/bin/bash

export APP_ENV=development
export GEN_DBUSER=postgres
export GEN_DBPASS=postgres
export GEN_DBSSLMODE=disable
export GEN_DBNAME=test
export GEN_DBHOST=db
export GEN_DBPORT=5432
export PGPASSWORD=postgres

set -e

host="db"

until psql -h "$host" -U "postgres" -c '\l'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
migrate -url postgres://postgres@db/test?sslmode=disable -path /go/src/github.com/DaveBlooman/go-app/migrations up


./myapp
