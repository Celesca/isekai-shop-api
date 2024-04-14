package main

import (
	"github.com/Celesca/isekai-shop-api/config"
	"github.com/Celesca/isekai-shop-api/databases"
	"github.com/Celesca/isekai-shop-api/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db.ConnectionGetting())

	server.Start()

}
