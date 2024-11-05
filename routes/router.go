package routes

import (
	"github.com/asmaulh99/dating-app-backend/handlers"
	"github.com/asmaulh99/dating-app-backend/pkg/auth"
	"github.com/asmaulh99/dating-app-backend/pkg/configs"
	"github.com/asmaulh99/dating-app-backend/pkg/cors"
	"github.com/asmaulh99/dating-app-backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct {
	ProfileService services.IProfileService
}

func GetRouters(cfg *configs.Config, authCl auth.Authenticator, dbCl *gorm.DB, uc Service) *gin.Engine {
	r := gin.Default()
	if cfg.Env == configs.EnvProduction {
		gin.SetMode(gin.ReleaseMode)
	}
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "PONG")
	})

	r.Use(cors.GetCORSConfig())
	apiV1 := r.Group("api/v1")
	profileHandler := handlers.NewProfileHandler(uc.ProfileService)
	profileV1Group := apiV1.Group(
		"/user/:userID/profile",
		// auth.Authorize(authCl),
	)
	profileV1Group.GET("/recomendation", profileHandler.GetRecomendationProfile)

	return r
}
