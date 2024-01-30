package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kundanacc20/Offer_Rolledout/db"
)

func main() {
	dsn := "root:Itsmypasword@047@tcp(localhost:3306)/interview_dashboard1"
	database := db.InitDB(dsn)

	defer database.Close()

	r := gin.Default()

	// rest api end point
	r.GET("/candidates_offers_rolledout_accepted", GetCandidatesWithAcceptedOffers)
	// r.GET("/candidates_offer_rolledout_awaited", GetCandidatesWithAwaitedOffers)
	// r.GET("/total_count_candidates_offer_rolledout_accepted", GetAcceptedCandidatesCount)
	// r.GET("/total_count_candidates_offer_rolledout_awaited", GetAwaitedCandidatesCount)

	// Run the server on port 9090
	if err := r.Run(":9090"); err != nil {
		log.Fatal(err)
	}
}
