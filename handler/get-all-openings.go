package handler

import (
	"net/http"

	"github.com/alexandrejuniorc/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Get Openings
// @Description Get all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} GetAllOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func GetAllOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-opening", openings)
}
