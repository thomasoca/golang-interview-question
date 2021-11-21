package main

import (
	"database/sql"
	"fmt"
	"time"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func logDb(searchWord string, page int) {
	db, err := connect()
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into log_db values (?, ?, ?)", currentTime, searchWord, page)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")

}
