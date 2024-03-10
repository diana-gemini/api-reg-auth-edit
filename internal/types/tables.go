package types

import (
	"database/sql"
)

func CreateTables(db *sql.DB) error {
	stmt := `	
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(255),
			email VARCHAR(255) UNIQUE NOT NULL,
			password CHAR(60),
			mobile VARCHAR(100) ,
			birthdate TEXT,
			token TEXT,
			expires DATETIME
		);	
			
	`
	// new:=`gfdg fdfs dg `

	_, err := db.Exec(stmt)

	return err
}
