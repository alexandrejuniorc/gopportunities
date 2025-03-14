package handler

import (
	"net/http"

	"github.com/alexandrejuniorc/gopportunities/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOpeningHandler(ctx *gin.Context) {
	// CreateOpeningRequest struct is defined in handler/request.go
	request := CreateOpeningRequest{}

	// BindJSON is a method from gin.Context
	ctx.BindJSON(&request)

	// Validate is a method from CreateOpeningRequest struct defined in handler/request.go
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Model:    gorm.Model{},
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	// Create is a method from gorm.DB
	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
