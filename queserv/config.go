package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

type DBConfig struct {
	Name  string
	Title string
	Path  string
}

var (
	theDBConfigs []DBConfig
)

// init theDBConfigs
func init() {
	configFileName := path.Join("./config/", "config.json")
	f, err := os.Open(configFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&(theDBConfigs))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Read theDBConfigs is :", theDBConfigs)
}

// DBItemByName is a sugar for dbconfig.DBConfigs
func DBConfigByName(name string) *DBConfig {
	for i, dbconfig := range theDBConfigs {
		if dbconfig.Name == name {
			return &theDBConfigs[i]
		}
	}
	return nil
}

// OpenDB open database for dbname
func OpenDB(dbname string) (db *sql.DB, err error) {
	dbconfig := DBConfigByName(dbname)
	if dbconfig == nil {
		log.Fatalf("%v not find.\n", dbname)
	}
	return sql.Open("sqlite3", dbconfig.Path)
}
