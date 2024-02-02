	package database

	import (
		"database/sql"
		"fmt"
		"foodcourt/app/config"
		"log"

		"github.com/go-sql-driver/mysql"
	)

	func InitDb(config *config.Config) *sql.DB {
		conf := mysql.Config{
			User:                 config.DBUsername,
			Passwd:               config.DBPassword,
			Net:                  "tcp",
			Addr:                 config.DBHost,
			DBName:               config.DBName,
			AllowNativePasswords: true,
		}
		db, err := sql.Open("mysql", conf.FormatDSN())

		if err != nil {
			log.Fatal(err)
			fmt.Println("database connexion failed")
		}

		if err = db.Ping(); err != nil {
			log.Fatal(err)
			fmt.Println("database connexion failed")
		}

		fmt.Println("database connexion successful")

		return db
	}
