package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "root"
	password = "root"
	dbname   = "db"
)

func ExecuteQuery(query string, resultStruct interface{}) {
	psqlInfo := "host=" + host + " port=" + port + " user=" + user +
		" password=" + password + " dbname=" + dbname + " sslmode=disable"

	// Открытие соединения
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Пример выполнения запроса
	rows, err := db.Queryx(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if resultStruct != nil {
		for rows.Next() {
			if err := rows.StructScan(resultStruct); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func ExecuteSelectQueryMultipleResults(query string, multipleResultStruct interface{}) {
	psqlInfo := "host=" + host + " port=" + port + " user=" + user +
		" password=" + password + " dbname=" + dbname + " sslmode=disable"

	// Открытие соединения
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db.Select(multipleResultStruct, query)
}
