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

type DataBaseItem struct {
	Name  string
	Title string
	Path  string
}

type DataBaseConfig struct {
	DataBaseItems []DataBaseItem
}

var dbconfig *DataBaseConfig

func ReadConfig(configfilename string) *DataBaseConfig {
	f, err := os.Open(configfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var dbconfig DataBaseConfig
	err = json.NewDecoder(f).Decode(&dbconfig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dbconfig)
	return &dbconfig
}

func init() {
	dbconfig = ReadConfig(path.Join("./config/", "config.json"))
}

// DBConfig return global DataBaseConfig
func DBConfig() *DataBaseConfig {
	return dbconfig
}

func GetDatabasePath(name string) string {
	for _, base := range DBConfig().DataBaseItems {
		if base.Name == name {
			return base.Path
		}
	}
	return ""
}

// DBItemByName is a sugar for dbconfig.DataBaseItems
func DBItemByName(name string) *DataBaseItem {
	for _, base := range DBConfig().DataBaseItems {
		if base.Name == name {
			return &base
		}
	}
	return nil
}

// OpenDB open database for dbname
func OpenDB(dbpath string) (db *sql.DB, err error) {
	return sql.Open("sqlite3", dbpath)
}
