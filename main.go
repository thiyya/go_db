package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Result struct {
	LastInsertedId int64 `json:"lastInsertedId"`
	RowsAffected   int64 `json:"rowsAffected"`
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Root")
	})
	http.HandleFunc("/mySql/people", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(getPeopleByMysql())
			break
		case "POST":
			writer.Header().Set("Content-Type", "application/json")
			lastInsertedId, rowsAffected := addPersonByMysql("erhan")
			json.NewEncoder(writer).Encode(Result{LastInsertedId: lastInsertedId, RowsAffected: rowsAffected})
		default:
			fmt.Fprintf(writer, "Not supported.")
		}
	})
	http.HandleFunc("/postgreSql/people", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(getPeopleByPostgreSql())
			break
		case "POST":
			writer.Header().Set("Content-Type", "application/json")
			lastInsertedId, rowsAffected := addPersonByPostgreSql("erhan_postgre")
			json.NewEncoder(writer).Encode(Result{LastInsertedId: lastInsertedId, RowsAffected: rowsAffected})
		default:
			fmt.Fprintf(writer, "Not supported.")
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getPeopleByMysql() []*Person {
	// user:password@tcp(dockerservicename:port)/dbName
	db, err := sql.Open("mysql", "erhan:pass@tcp(mysql_db:3306)/mysqlDB")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	results, err := db.Query("SELECT * FROM Person")
	if err != nil {
		log.Fatal(err.Error())
	}

	var people []*Person
	for results.Next() {
		var p Person
		err = results.Scan(&p.ID, &p.Name)
		if err != nil {
			log.Print(err.Error())
		}
		people = append(people, &p)
	}
	return people
}

func addPersonByMysql(name string) (int64, int64) {
	// user:password@tcp(dockerservicename:port)/dbName
	db, err := sql.Open("mysql", "erhan:pass@tcp(mysql_db:3306)/mysqlDB")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	result, err := db.Exec("INSERT INTO Person (`name`) VALUES ('" + name + "');")
	if err != nil {
		log.Fatal(err.Error())
	}
	lastInsertedId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	return lastInsertedId, rowsAffected
}

func getPeopleByPostgreSql() []*Person {
	connStr := "postgres://postgres:postgres@postgre_db:5432?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	results, err := db.Query("SELECT * FROM Person")
	if err != nil {
		log.Fatal(err.Error())
	}

	var people []*Person
	for results.Next() {
		var p Person
		err = results.Scan(&p.ID, &p.Name)
		if err != nil {
			log.Print(err.Error())
			return []*Person{}
		}
		people = append(people, &p)
	}
	return people
}

func addPersonByPostgreSql(name string) (int64, int64) {
	connStr := "postgres://postgres:postgres@postgre_db:5432?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	result, err := db.Exec("INSERT INTO Person (name) VALUES ('" + name + "');")
	if err != nil {
		log.Fatal(err.Error())
	}
	lastInsertedId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	return lastInsertedId, rowsAffected
}
