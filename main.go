package main

import (
	"log"

	"github.com/kundanacc20/Offer_Rolledout/db"
	"github.com/kundanacc20/Offer_Rolledout/routers"
)

func main() {
	dsn := "root:Itsmypasword@047@tcp(localhost:3306)/interview_dashboard1"
	database := db.InitDB(dsn)

	defer database.Conn.Close()

	// Call SetupRouter from routers package
	r := routers.SetupRouter(database)

	// Run the server on port 9090
	if err := r.Run(":9090"); err != nil {
		log.Fatal(err)
	}
}
