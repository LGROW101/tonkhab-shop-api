package main

import (
	"fmt"

	"github.com/LGROW101/tonkhab-shop-api/config"
	"github.com/LGROW101/tonkhab-shop-api/databases"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	fmt.Println(db.Connect())
}
