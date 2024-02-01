package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kundanacc20/Offer_Rolledout/db"
	"github.com/kundanacc20/Offer_Rolledout/handlers"
)

func main() {
	dsn := "root:Itsmypasword@047@tcp(localhost:3306)/interview_dashboard1"
	database := db.InitDB(dsn)

	defer database.Close()

	r := gin.Default()

	// rest api end point
	r.GET("/candidates_offers_rolledout_accepted", func(c *gin.Context) { handlers.GetCandidatesWithAcceptedOffers(database, c) })
	r.GET("/candidates_offers_rolledout_awaited", func(c *gin.Context) { handlers.GetCandidatesWithAwaitedOffers(database, c) })
	r.GET("/count_candidates_offers_rolledout_accepted", func(c *gin.Context) { handlers.GetAcceptedCandidatesCount(database, c) })
	r.GET("/count_candidates_offers_rolledout_awaited", func(c *gin.Context) { handlers.GetAwaitedCandidatesCount(database, c) })

	// Run the server on port 9090
	if err := r.Run(":9090"); err != nil {
		log.Fatal(err)
	}
}
