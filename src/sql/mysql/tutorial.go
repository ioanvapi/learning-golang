package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func main() {
	db, err := sql.Open("mysql", "root:passwd@tcp(45.77.24.104:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// db ping
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	{
		// db query
		var (
			id   int
			name string
		)

		rows, err := db.Query("select id, name from users where id = ?", 1)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			log.Info(id, name)
		}

		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}
	}

	{
		// db prepare
		var (
			id   int
			name string
		)

		stmt, err := db.Prepare("select id, name from users where id = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		rows, err := stmt.Query(1)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			log.Info(id, name)
		}

		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}
	}

	{
		// single row query on db
		var name string
		err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Info(name)
	}

	{
		// single row query on prepare
		stmt, err := db.Prepare("select name from users where id = ?")
		if err != nil {
			log.Fatal(err)
		}
		var name string
		err = stmt.QueryRow(1).Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Info(name)
	}

	{
		// Statements that Modify Data
		stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
		if err != nil {
			log.Fatal(err)
		}
		res, err := stmt.Exec("Dolly")
		if err != nil {
			log.Fatal(err)
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("ID = %d, affected = %d\n", lastId, rowCnt)
	}
}
