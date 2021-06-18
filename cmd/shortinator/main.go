/*
                             GoShortrr

	A fast, simple and powerful URL Shortener built with Go and React.

				Copyright (c) 2021 Fabian R. (fabiancdng)

*/

package main

import (
	"github.com/fabiancdng/GoShortrr/internal/database/mysql"
	"github.com/fabiancdng/GoShortrr/internal/webserver"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	///////////////////////
	//      DATABASE     //
	///////////////////////

	db := new(mysql.MySQL)
	if err := db.Open(); err != nil {
		panic(err)
	}

	// Make sure all tables exist in database
	if err := db.Init(); err != nil {
		panic(err)
	}

	//////////////////////
	//     WEBSERVER    //
	//////////////////////

	// Create WebServer
	ws, err := webserver.NewWebServer(db)
	if err != nil {
		panic(err)
	}

	// Run WebServer
	err = ws.RunWebServer()
	if err != nil {
		panic(err)
	}
}
