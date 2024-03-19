package handlers

import "database/sql"

var db *sql.DB

func SetupHandlers(dbObj *sql.DB) {
	db = dbObj
}
