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

	////////////////////////////////
	//                            //
	//           CONFIG           //
	//                            //
	////////////////////////////////

	// Read and parse the config.yml file
	config, err := config.ParseConfig("./config/config.yml")
	if err != nil {
		panic(err)
	}

	//////////////////////////////////
	//                              //
	//           DATABASE           //
	//                              //
	//////////////////////////////////

	// Instantiate Database object
	db := new(mysql.MySQL)
	if err := db.Open(config); err != nil {
		panic(err)
	}

	// Prepare database
	if err := db.Init(); err != nil {
		panic(err)
	}

	///////////////////////////////////
	//                               //
	//           WEBSERVER           //
	//                               //
	///////////////////////////////////

	// Instantiate WebServer object
	ws, err := webserver.NewWebServer(db, config)
	if err != nil {
		panic(err)
	}

	// Run the webserver
	err = ws.RunWebServer()
	if err != nil {
		panic(err)
	}
}
