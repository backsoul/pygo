package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type event struct {
	Name string
}

type message struct {
	Message string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func connectDB() (*sql.DB, error) {
	mysqlUser := getEnv("MYSQL_USER", "your_mysql_username")
	mysqlPassword := getEnv("MYSQL_PASSWORD", "your_mysql_password")
	mysqlHost := getEnv("MYSQL_HOST", "mysql")
	mysqlPort := getEnv("MYSQL_PORT", "3306")
	mysqlDatabase := getEnv("MYSQL_DATABASE", "your_mysql_database")

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)
	return sql.Open("mysql", mysqlInfo)
}

func executeSQL(db *sql.DB, sqlStatement string, args ...interface{}) error {
	_, err := db.Exec(sqlStatement, args...)
	return err
}

func handleDatabaseOperation(w http.ResponseWriter, r *http.Request, sqlStatement string, args ...interface{}) {
	db, err := connectDB()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err := executeSQL(db, sqlStatement, args...); err != nil {
		http.Error(w, "Error executing SQL statement", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	var response message
	response.Message = "Success"
	json.NewEncoder(w).Encode(response)
}

// CREATE
func create(w http.ResponseWriter, r *http.Request) {
	var createEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &createEvent)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO data (name) VALUES (?)`
	handleDatabaseOperation(w, r, sqlStatement, createEvent.Name)
}

// READ
func read(w http.ResponseWriter, r *http.Request) {
	var response message
	name := "None"
	sqlStatement := `SELECT * FROM data LIMIT 1`

	db, err := connectDB()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	err = db.QueryRow(sqlStatement).Scan(&name)
	if err != nil {
		http.Error(w, "Error executing SQL statement", http.StatusInternalServerError)
		return
	}

	response.Message = name
	json.NewEncoder(w).Encode(response)
}

// UPDATE
func update(w http.ResponseWriter, r *http.Request) {
	var createEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &createEvent)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	sqlStatement := `UPDATE data SET name = ?`
	handleDatabaseOperation(w, r, sqlStatement, createEvent.Name)
}

// DELETE
func delete(w http.ResponseWriter, r *http.Request) {
	var createEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &createEvent)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	sqlStatement := `DELETE FROM data WHERE name = ?`
	handleDatabaseOperation(w, r, sqlStatement, createEvent.Name)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/create", create).Methods("POST")
	router.HandleFunc("/read", read).Methods("GET")
	router.HandleFunc("/update", update).Methods("PUT")
	router.HandleFunc("/delete", delete).Methods("DELETE")

	port := getEnv("APP_PORT", "8080")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
