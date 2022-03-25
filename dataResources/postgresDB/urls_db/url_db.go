package urls_db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	Client *sql.DB
)

const (
	driverName = "DRIVER"
	dbUser     = "USER"
	dbPassword = "PASSWORD"
	dbPort     = "PORT"
	dbName     = "DBNAME"
	dbSSLmode  = "SSLMODE"
)

var (
	driver   string
	user     string
	password string
	port     string
	name     string
	sslmode  string
)

func loadEnv() {
	godotenv.Load(".env")
	driver = os.Getenv(driverName)
	user = os.Getenv(dbUser)
	password = os.Getenv(dbPassword)
	port = os.Getenv(dbPort)
	name = os.Getenv(dbName)
	sslmode = os.Getenv(dbSSLmode)
}

func init() {
	loadEnv()
	dataSourceName := fmt.Sprintf("user=%s port=%s dbname=%s password=%s sslmode=%s", user, port, name, password, sslmode)
	fmt.Println(dataSourceName)
	var err error
	Client, err = sql.Open(driver, dataSourceName)
	if err != nil {
		fmt.Println("Database Connectivity Failed")
		panic(err)
	}
	pingErr := Client.Ping()
	if pingErr != nil {
		fmt.Println("ping to database failed")
		panic(pingErr)
	}
	fmt.Println("Database Connected Successfully")

}
