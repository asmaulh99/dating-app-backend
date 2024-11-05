package main

import (
	"github.com/asmaulh99/dating-app-backend/pkg/auth"
	"github.com/asmaulh99/dating-app-backend/pkg/configs"
	"github.com/asmaulh99/dating-app-backend/pkg/database"
	"github.com/asmaulh99/dating-app-backend/pkg/helpers"
	"github.com/asmaulh99/dating-app-backend/routes"
	"github.com/asmaulh99/dating-app-backend/services"

	profileRepo "github.com/asmaulh99/dating-app-backend/infrastructure/profile"
)

func main() {
	cfg := configs.GetConfig()

	httpClient := auth.NewClientHTTP(cfg.AuthSecretKey)
	dbConn, err := database.CreateDBConn(cfg.DBDebugMode,
		"mysql",
		&database.DBConnectionProps{
			Hostname: cfg.DBHost,
			Username: cfg.DBUsername,
			Password: cfg.DBPassword,
			DBName:   cfg.DBName,
		},
	)
	if err != nil {
		panic(err)
	}

	profileService := services.NewProfileService(profileRepo.NewMysqlRepository(dbConn))

	router := routes.GetRouters(cfg, httpClient, dbConn, routes.Service{
		ProfileService: profileService,
	})

	helpers.RunServerGraceFully(cfg.RestPort, router)
}
