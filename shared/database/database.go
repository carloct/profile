package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sqlx.DB
)

type Databases struct {
	Type   string
	MySQL  MySQLInfo
	SQLite SQLiteInfo
}

type MySQLInfo struct {
	Username  string
	Password  string
	Name      string
	Hostname  string
	Port      int
	Parameter string
}

type SQLiteInfo struct {
	Parameter string
}

func DSN(ci MySQLInfo) string {
	return ci.Username +
		":" +
		ci.Password +
		"@tcp(" +
		ci.Hostname +
		":" +
		fmt.Sprintf("%d", ci.Port) +
		")/" +
		ci.Name + ci.Parameter
}

func Connect(d Databases) {
	var err error

	switch d.Type {
	case "MySQL":
		// Connect to MySQL
		if DB, err = sqlx.Connect("mysql", DSN(d.MySQL)); err != nil {
			log.Println("SQL Driver Error", err)
		}

		// Check if is alive
		if err = DB.Ping(); err != nil {
			log.Println("Database Error", err)
		}
	case "SQLite":
		// Connect to SQLite
		if DB, err = sqlx.Connect("sqlite3", d.SQLite.Parameter); err != nil {
			log.Println("SQL Driver Error", err)
		}

		// Check if is alive
		if err = DB.Ping(); err != nil {
			log.Println("Database Error", err)
		}
	default:
		log.Println("No registered database in config")
	}
}
