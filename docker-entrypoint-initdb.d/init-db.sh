#!/usr/bin/env bash

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER gammamobility;
    CREATE DATABASE gammamobility_db ENCODING UTF8;
    GRANT ALL PRIVILEGES ON DATABASE gammamobility_db TO gammamobility;
    ALTER USER gammamobility WITH PASSWORD 'gammamobility';
    ALTER USER gammamobility WITH SUPERUSER;
EOSQL