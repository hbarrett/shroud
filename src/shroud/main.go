package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"time"
)

var VERSION string
var CODE string
var CODENAME string

//create a DB handle
var db *sql.DB
var keyStr string
var background string
var foreground string
var formbackground string
func main() {
	VERSION = "0.0"
	CODENAME = "peek"
	
	var configf = ReadConfig() //this is in config.go
	keyStr = configf.Key
	background = configf.Background
	foreground = configf.Foreground
	formbackground = configf.FormBackground
	//var LogFile string
	LogFile := configf.LogDir + "shroud.log"

	dbslice := []string{}
	var err error
	dbuser := configf.DBUsername
	dbpassword := configf.DBPassword
	dbname := configf.DBName
	dsn := dbuser + ":" + dbpassword + "@/" + dbname
	db, err = sql.Open("mysql", dsn) // this does not really open a new connection
	if err != nil {
		log.Fatalf("Error on initializing snort database connection: %s", err.Error())
	}
	//max connections allowed to the snortdb
	db.SetMaxIdleConns(100)
	err = db.Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
	}
	var databases string

	query := "SHOW DATABASES;"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&databases)
		if err != nil {
			log.Fatal(err)
		}
		dbslice = Append(dbslice, databases)
	}
	slicehasdb := contains(dbslice, dbname)
	//	fmt.Println(slicehasdb)
	if slicehasdb == false {
		log.Fatalf("Database %s does not exist.", dbname)
	}
	query = "SHOW TABLES;"
	rows, err = db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var tables string
	for rows.Next() {
		err := rows.Scan(&tables)
		if err != nil {
			log.Fatal(err)
		}
		dbslice = Append(dbslice, tables)
	}
	slicehastable := contains(dbslice, "secrets")
	//      fmt.Println(slicehasdb)
	if slicehastable == false {
		log.Println("Table secrets does not exist....creating it.")
		createtable := "CREATE TABLE secrets (secret VARCHAR(1000), token VARCHAR(128), date DATE, views TINYINT(2));"
		rows, err = db.Query(createtable)
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err := os.OpenFile(LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Cant open log file")
	}
	defer f.Close()

	log.SetOutput(f)

	ticker := time.NewTicker(time.Hour * 1)
	go func() {
		for t := range ticker.C {
			cleanup()
			_ = t
		}
	}()

	listensocket := configf.IP + ":" + configf.Port
	router := NewRouter()
	log.Println("shroud running on " + listensocket)
	log.Fatal(http.ListenAndServe(listensocket, router))
}

func Append(slice []string, items ...string) []string {
	for _, item := range items {
		slice = Extend(slice, item)
	}
	return slice
}

func Extend(slice []string, element string) []string {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]string, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func cleanup() {

	t := time.Now()
	t.Format(time.RFC3339)
	t0 := t.Format("2006-01-02")
	stmt, err := db.Prepare("DELETE FROM secrets WHERE date < ?")
	logErr(err)
	result, err := stmt.Exec(t0)
	logErr(err)
	affect, err := result.RowsAffected()
	logErr(err)
	log.Println(affect)

}

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
