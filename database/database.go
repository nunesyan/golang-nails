package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := "postgresql://postgres.yfejmegmzycacystolhy:yannunesatzler@aws-0-sa-east-1.pooler.supabase.com:6543/postgres"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return db, err
	}

	return db, nil
}
