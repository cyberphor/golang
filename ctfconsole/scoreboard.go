package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func CreateScoreboard() {
    scoreboard, err := sql.Open("sqlite3","./scoreboard.db")
    if err != nil { 
        panic(err)
    }
    statement := `CREATE TABLE IF NOT EXISTS players (
        id INT AUTO_INCREMENT NOT NULL, 
        user VARCHAR(64) NOT NULL, 
        pass VARCHAR(255) NOT NULL, 
        PRIMARY KEY ('id')
    );`
    query, err := scoreboard.Prepare(statement)
    if err != nil { 
        panic(err)
    }
    query.Exec()
    scoreboard.Close()
}