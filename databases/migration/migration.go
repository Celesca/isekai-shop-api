// Run the schema migration for the database
// Run through function main
package main

import (
	"fmt"

	"github.com/Celesca/isekai-shop-api/config"
	"github.com/Celesca/isekai-shop-api/databases"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	fmt.Println(db.ConnectionGetting())
}
