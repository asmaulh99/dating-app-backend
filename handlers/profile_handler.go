package handlers

import (
	"net/http"
	"strconv"

	"github.com/asmaulh99/dating-app-backend/entities"
	"github.com/asmaulh99/dating-app-backend/pkg/helpers"
	"github.com/asmaulh99/dating-app-backend/services"
	"github.com/gin-gonic/gin"
)

type IProfileHandler interface {
	GetRecomendationProfile(ctx *gin.Context)
}

type ProfileHandler struct {
	ProfileService services.IProfileService
}

func NewProfileHandler(profileService services.IProfileService) IProfileHandler {
	return &ProfileHandler{ProfileService: profileService}
}

func (h *ProfileHandler) GetRecomendationProfile(ctx *gin.Context) {
	var fetchQuery entities.GetProfileQuery

	if userIDStr := ctx.Param("userID"); userIDStr != "" {
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			ctx.String(http.StatusBadRequest, "invalid userID")
			return
		}
		fetchQuery.ProfileID = userID
	}

	report, err := h.ProfileService.GetRecomendationProfile(ctx, fetchQuery)
	if err != nil {
		helpers.SendJSONErrorResponse(ctx, err)
		return
	}
	if report == nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"recommendationProfile": map[string]interface{}{},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"recommendationProfile": report,
	})
}
