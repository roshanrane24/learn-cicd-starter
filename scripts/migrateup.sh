#!/bin/bash

if [ -f .env ]; then
    source .env
fi

cd sql/schema
#goose turso $DATABASE_URL up
GOOSE_DRIVER=turso GOOSE_DBSTRING=$DATABSE_URL goose up
