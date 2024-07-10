package config

import (
	"net/http"
	"os"
	"strings"
  "embed"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

type Env struct {
	/**
	* define the fields inside you're env
	* in order to have them ready to use
	* ex. Env.DB_NAME and so on
	 */
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Error(".env file was not found. Are you sure you have one ?")
	}
}

/**
* This will take the database type definde in the env and also all
* the needed inputs for it ex pw , connection string and so on
 */
func ConnectToDatabase() {
	engine := os.Getenv("DB_ENGINE")
	switch strings.ToLower(engine) {
	case "mysql":
		log.Info("Connecting to database via " + engine)
	case "postgress":
		log.Info("Connecting to database via " + engine)
	case "sqlite3":
		log.Info("Connecting to database via " + engine)
	case "":
		log.Error("The DB_ENGINE inside the env File may not be set !")

	}
}

func StaticFiles(path string) {
	fs := http.FileServer(http.Dir("./" + path))
	http.Handle("/static/", http.StripPrefix("/"+path+"/", fs))
}

func EmbededFiles(path string) embed.FS{
  // here the files that need to be embeded
  var embeded embed.FS
  return embeded
}
