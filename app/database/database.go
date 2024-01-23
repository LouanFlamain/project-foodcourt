package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func InitDb() *sql.DB  {
	conf := mysql.Config{
		User: "root",
		Passwd: "root",
		Net: "tcp",
		Addr: "database:3306",
		DBName: "foodcourt_db",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", conf.FormatDSN())

	if(err != nil){
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