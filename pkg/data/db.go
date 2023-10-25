package data

import (
	"database/sql"
	"log"

	"github.com/livghit/go-htmx/pkg/config"
	_ "github.com/mattn/go-sqlite3"
)

// database implementation

type Database struct {
	engine *sql.DB
}

// Gets all the needed settings from the envFile
func (db *Database) Connect(envFile config.Env) *Database {
	// Initialisate the Database engine based on the engine from the env
	switch envFile.DBENGINE {
	case "sqlite":
		connection, err := sql.Open("sqlite3", "name")
		if err != nil {
			log.Fatalf("Error happened while connecting to the db : %s", err)
		}
		db.engine = connection
		break
	case "mysql":
		break
	case "postgress":
		break
	default:
		log.Fatal("Database not set inside the env file")
	}

	// in moment not set will be moved upstair in the switch statement
	return &Database{}
}
