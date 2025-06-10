package main

import (
	database "api-go/database"
	routers "api-go/routers"
	"log"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Falha ao conectar ao banco", err)
	}
	defer db.Close()

	r := routers.Routers(db)

	r.Run(":3030")
}
