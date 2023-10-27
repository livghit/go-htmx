package data

import (
	"database/sql"
	"log"
	"os"

	"github.com/livghit/go-htmx/pkg/config"
	"github.com/livghit/go-htmx/pkg/data/models"
	_ "github.com/mattn/go-sqlite3"
)

// database implementation

type Database struct {
	Engine *sql.DB
}

var DB *sql.DB

func InitDatabase() *Database {
	return &Database{}
}

// Gets all the needed settings from the envFile
func (db *Database) Connect(envFile config.Env) {
	// Initialisate the Database engine based on the engine from the env
	switch envFile.DBENGINE {
	case "sqlite":
		_, error := os.Stat(envFile.DBNAME + ".db")
		if os.IsNotExist(error) {
			//dose not exist
			os.Create(envFile.DBNAME + ".db")
		}
		//file exists
		connection, err := sql.Open("sqlite3", envFile.DBNAME+".db")
		if err != nil {
			log.Fatalf("Error happened while connecting to the db : %s", err)
		}
		log.Default().Print("Connection done ")
		// mapping the connection to the engine
		db.Engine = connection
		DB = db.Engine
		break
	case "mysql":
		log.Fatal("MYSQL engine is not implemented yet")
		break
	case "postgress":
		log.Fatal("Postgress engine is not implemented yet")
		break
	default:
		log.Fatal("Database not set inside the env file")
	}
}

// Find another better way to make migrations
func (db *Database) MakeMigration() {

	log.Default().Println("Making migration...")
	// using the global DB var to access the Prepare Method on the DB
	statement, err := DB.Prepare(models.MigrateProductTable()) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Default().Println("Migration done")
}
