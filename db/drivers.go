package db

// Import database drivers to register them with database/sql.
// Add additional driver imports here as needed.
import (
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver — registers "sqlite3"
)
