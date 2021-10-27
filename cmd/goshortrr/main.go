/*
                             GoShortrr

	  A fast, simple and powerful URL Shortener built with Go and React.

	            Copyright (c) 2021 Fabian Reinders (fabiancdng)

*/

package main

import (
	"github.com/fabiancdng/GoShortrr/internal/config"
	"github.com/fabiancdng/GoShortrr/internal/database/mysql"
	"github.com/fabiancdng/GoShortrr/internal/utils"
	"github.com/fabiancdng/GoShortrr/internal/webserver"
	"log"
	"time"
)

func main() {
	// Print startup banner
	utils.PrintStartupBanner()

	// Delay startup a little (for example to make sure all needed
	// Docker containers in the stack are online)
	utils.StartupDelay(10 * time.Second)

	////////////////////////////////
	//                            //
	//           CONFIG           //
	//                            //
	////////////////////////////////

	// Try to read and parse the config file 'config.yml'
	// If it doesn't exist, try to read and parse config from env variables
	config, err := config.ParseConfig("./config/config.yml")
	if err != nil {
		panic(err)
	}
	log.Println(config)

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
