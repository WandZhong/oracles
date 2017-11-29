#!/bin/sh

exec ./tranche-manager -pg-user "$POSTGRES_USER" -pg-pwd "$POSTGRES_PASSWORD" -pg-db "$POSTGRES_DATABASE" -pg-addr "$POSTGRES_HOST:$POSTGRES_PORT"
