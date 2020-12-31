package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	dbUser    string
	dbPswd    string
	dbHost    string
	dbPort    string
	dbName    string
	dbDialect string
}

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func configString(c *Config) string {
	result := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.dbUser, c.dbPswd, c.dbHost, c.dbPort, c.dbName)
	return result
}

func GetDBConfig() (string, string) {
	conf := &Config{}

	flag.StringVar(&conf.dbUser, "dbuser", os.Getenv("DB_USER"), "DB user name")
	flag.StringVar(&conf.dbPswd, "dbpswd", os.Getenv("DB_PASSWORD"), "DB pass")
	flag.StringVar(&conf.dbPort, "dbport", os.Getenv("DB_PORT"), "DB port")
	flag.StringVar(&conf.dbHost, "dbhost", os.Getenv("DB_HOST"), "DB host")
	flag.StringVar(&conf.dbName, "dbname", os.Getenv("DB_NAME"), "DB name")
	flag.StringVar(&conf.dbDialect, "dbdialect", os.Getenv("DB_DIALECT"), "DB dialect")

	flag.Parse()
	str := configString(conf)
	return conf.dbDialect, str
}
