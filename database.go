package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

//Created just so the sql database opens once and also
//inserts myTime and location into the database once

type LocalDatabase struct {
	db                *sql.DB
	stmtInsertMessage *sql.Stmt
}

func (me *LocalDatabase) init(filepath string) {
	var err error
	me.db, err = sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Opened db")
	}
	if me.db == nil {
		panic("db nil")
	}
	me.stmtInsertMessage, err = me.db.Prepare(`
	INSERT INTO chats (
		Name,
		Time,
		Message
		) VALUES (?,?,?)
		`)
	if err != nil {
		panic(err)
	}

}
