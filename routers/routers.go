package routers

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/kundanacc20/Offer_Rolledout/db"
	"github.com/kundanacc20/Offer_Rolledout/handlers"
)

// SetupRouter sets up the API endpoints
func SetupRouter(database *db.Database) *gin.Engine {
	r := gin.Default()

	// rest api end point
	r.GET("/interview-db/home/offer_rolled_out_accepted", func(c *gin.Context) { handlers.GetCandidatesWithAcceptedOffers(database, c) })
	r.GET("/interview-db/home/offer_rolled_out_awaited", func(c *gin.Context) { handlers.GetCandidatesWithAwaitedOffers(database, c) })
	r.GET("/interview-db/home/offer_rolled_out_accepted_count", func(c *gin.Context) { handlers.GetAcceptedCandidatesCount(database, c) })
	r.GET("/interview-db/home/offer_rolled_out_awaited_count", func(c *gin.Context) { handlers.GetAwaitedCandidatesCount(database, c) })

	r.GET("/interview-db/home/onboarded", func(c *gin.Context) { handlers.GetListOfAllCandidatOnboarded(database, c) })
	r.GET("/interview-db/home/onboarded_count", func(c *gin.Context) { handlers.GetOnboardedCount(database, c) })

	r.GET("/interview-db/home/DM", func(c *gin.Context) { handlers.GetListOfAllCandidatAtDM(database, c) })
	r.GET("/interview-db/home/DM_selected", func(c *gin.Context) { handlers.GetListOfAllCandidatAtDMSelected(database, c) })
	r.GET("/interview-db/home/DM_rejected", func(c *gin.Context) { handlers.GetListOfAllCandidatAtDMRejected(database, c) })
	r.GET("/interview-db/home/DM_pending", func(c *gin.Context) { handlers.GetListOfAllCandidatAtDMPending(database, c) })
	r.GET("/interview-db/home/DM_pending_count", func(c *gin.Context) { handlers.GetPendingDMCount(database, c) })
	r.GET("/interview-db/home/DM_selected_count", func(c *gin.Context) { handlers.GetSelectedDMCount(database, c) })
	r.GET("/interview-db/home/DM_rejected_count", func(c *gin.Context) { handlers.GetRejectedDMCount(database, c) })

	r.GET("/interview-db/home/L1", func(c *gin.Context) { handlers.GetListOfAllCandidatAtL1(database, c) })
	r.GET("/interview-db/home/L1_pending", func(c *gin.Context) { handlers.GetListOfAllCandidatAtL1Pending(database, c) })
	r.GET("/interview-db/home/L1_pending_count", func(c *gin.Context) { handlers.GetPendingL1Count(database, c) })
	r.GET("/interview-db/home/L1_selected", func(c *gin.Context) { handlers.GetListOfAllCandidatAtL1Selected(database, c) })
	r.GET("/interview-db/home/L1_rejected", func(c *gin.Context) { handlers.GetListOfAllCandidatAtL1Rejected(database, c) })
	r.GET("/interview-db/home/L1_count_selected", func(c *gin.Context) { handlers.GetSelectedL1Count(database, c) })
	r.GET("/interview-db/home/L1_count_rejected", func(c *gin.Context) { handlers.GetRejectedL1Count(database, c) })
	r.GET("/interview-db/home/L2", func(c *gin.Context) { handlers.GetListOfAllCandidatAtL2(database, c) })
	r.GET("/interview-db/home/L2_pending", func(c *gin.Context) { handlers.GetListOfAllCandidatAtL2Pending(database, c) })
	r.GET("/interview-db/home/L2_pending_count", func(c *gin.Context) { handlers.GetPendingL2Count(database, c) })
	r.GET("/interview-db/home/L2_selected", func(c *gin.Context) { handlers.GetListOfAllCandidatAtL2Selected(database, c) })
	r.GET("/interview-db/home/L2_rejected", func(c *gin.Context) { handlers.GetListOfAllCandidatAtL2Rejected(database, c) })
	r.GET("/interview-db/home/L2_count_selected", func(c *gin.Context) { handlers.GetSelectedL2Count(database, c) })
	r.GET("/interview-db/home/L2_count_rejected", func(c *gin.Context) { handlers.GetRejectedL2Count(database, c) })

	r.GET("/interview-db/home/resume_shortlisted", func(c *gin.Context) { handlers.GetListOfAllCandidatResumeShortListed(database, c) })
	r.GET("/interview-db/home/resume_rejected", func(c *gin.Context) { handlers.GetListOfAllCandidatResumeRejected(database, c) })

	r.POST("/interview-db/home/admin/candidates", func(c *gin.Context) { handlers.AddCandidateToResume(database, c) })
	r.GET("/interview-db/home/admin/candidates", func(c *gin.Context) { handlers.GetListOfAllCandidate(database, c) })
	r.PUT("/interview-db/home/admin/candidates/:id", func(c *gin.Context) { handlers.UpdateCandidate(database, c) })
	r.POST("/interview-db/home/admin/interviewstatus", func(c *gin.Context) { handlers.AddCandidateToInteviewStatusTable(database, c) })
	r.PUT("/interview-db/home/admin/inteviewstatus/:id", func(c *gin.Context) { handlers.UpdateCandidateAtInterviewStatus(database, c) })

	r.POST("/interview-db/register", func(c *gin.Context) { handlers.SignupHandler(database, c) })

	// Set the HTML template directory
	absolutePath, err := filepath.Abs("views/*")
	if err != nil {
		log.Fatal(err)
	}
	r.LoadHTMLGlob(absolutePath)

	// Register the login page
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// Register the user login handler
	r.POST("/login", func(c *gin.Context) {
		handlers.UserLoginHandler(database, c)
	})

	return r
}
