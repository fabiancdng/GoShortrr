/*
                             GoShortrr

	  A fast, simple and powerful URL Shortener built with Go and React.

	            Copyright (c) 2021 Fabian Reinders (fabiancdng)

*/

package main

import (
	"github.com/fabiancdng/GoShortrr/internal/config"
	"github.com/fabiancdng/GoShortrr/internal/database/mysql"
	"github.com/fabiancdng/GoShortrr/internal/webserver"
)

func main() {
	//////////////////////
	//      CONFIG      //
	//////////////////////

	config, err := config.ParseConfig("./config/config.yml")
	if err != nil {
		panic(err)
	}

	///////////////////////
	//      DATABASE     //
	///////////////////////

	db := new(mysql.MySQL)
	if err := db.Open(config); err != nil {
		panic(err)
	}

	// Makes sure all tables exist in database
	if err := db.Init(); err != nil {
		panic(err)
	}

	//////////////////////
	//     WEBSERVER    //
	//////////////////////

	// Creates WebServer
	ws, err := webserver.NewWebServer(db, config)
	if err != nil {
		panic(err)
	}

	// Runs WebServer
	err = ws.RunWebServer()
	if err != nil {
		panic(err)
	}
}
