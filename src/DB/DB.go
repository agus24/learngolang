package DB

import (
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"log"
)

func Init() (*sql.DB) {
	cfg := mysql.Config{
		User: "root",
		Passwd: "",
		Net: "tcp",
		Addr: "localhost:3306",
		DBName: "recordings",
	}

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())

	// defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected.")

	return db
}

