package services

import (
	"github.com/asmaulh99/dating-app-backend/entities"
	"github.com/asmaulh99/dating-app-backend/repositories"
	"github.com/gin-gonic/gin"
)

type IProfileService interface {
	GetRecomendationProfile(ctx *gin.Context, query entities.GetProfileQuery) ([]*entities.Profile, error)
}
type ProfileService struct {
	ProfileRepository repositories.ProfileRepository
}

func NewProfileService(profileRepository repositories.ProfileRepository) IProfileService {
	return &ProfileService{
		ProfileRepository: profileRepository,
	}
}

func (service *ProfileService) GetRecomendationProfile(ctx *gin.Context, query entities.GetProfileQuery) ([]*entities.Profile, error) {
	var data []*entities.Profile

	query.Limit = 10
	data, _ = service.ProfileRepository.GetRecomendationProfile(ctx, query)
	return data, nil
}
