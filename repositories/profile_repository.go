package repositories

import (
	"context"

	"github.com/asmaulh99/dating-app-backend/entities"
)

type ProfileRepository interface {
	GetRecomendationProfile(ctx context.Context, query entities.GetProfileQuery) ([]*entities.Profile, error)
}
