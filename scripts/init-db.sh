#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL

	CREATE TABLE IF NOT EXISTS planets (
	        id integer NOT NULL PRIMARY KEY,
	        name varchar(100) NOT NULL,
	        climate varchar(100) NOT NULL,
	        terrain varchar(100) NOT NULL,
	        created_at timestamptz NOT NULL DEFAULT now(),
	        updated_at timestamptz NOT NULL DEFAULT now(),
	        UNIQUE (name) 
	);
	CREATE TABLE IF NOT EXISTS films (
	        id integer NOT NULL,
	        planet_id integer REFERENCES planets (id) ON DELETE CASCADE,
	        title varchar(64) NOT NULL,
	        director varchar(150) NOT NULL,
	        release_date timestamptz NOT NULL DEFAULT now(),
	        created_at timestamptz NOT NULL DEFAULT now(),
	        updated_at timestamptz NOT NULL DEFAULT now()
	);
EOSQL
