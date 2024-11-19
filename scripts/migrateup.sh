#!/bin/bash

if [ -f .env ]; then
    source .env
fi

cd sql/schema
#goose turso $DATABASE_URL up
GOOSE_DRIVER=turso GOOSE_DBSTRING="libsql://notely-db-roshanrane24.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MzIwMDk2NzcsImlkIjoiZDVmNGRiMjYtMmUyMi00ODVjLTgxZWQtMDk5M2RhZmVmZDg2In0.9sll1uRsuH0ECA4RiWACkUuleul_ZI9SHdwUvgyaMTdhbucX1NLteIZUaSqOaKu_4advdhRAKrljN48EJUOTDA" goose up
