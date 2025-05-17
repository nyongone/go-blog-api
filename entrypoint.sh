#!/bin/sh

dockerize -wait tcp://$DB_HOST:$DB_PORT -timeout 30s

./migrate

exec "$@"