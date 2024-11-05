package profile

import (
	"context"

	"github.com/asmaulh99/dating-app-backend/entities"
	"github.com/asmaulh99/dating-app-backend/repositories"
	"gorm.io/gorm"
)

type mySqlRepository struct {
	dbConn *gorm.DB
}

func NewMysqlRepository(dbConn *gorm.DB) repositories.ProfileRepository {
	return &mySqlRepository{
		dbConn: dbConn,
	}
}

func (repo *mySqlRepository) GetRecomendationProfile(ctx context.Context, query entities.GetProfileQuery) ([]*entities.Profile, error) {
	var profileModel []*ProfileModel

	rawQuery := `
	SELECT 
		p.*, 
		COUNT(li2.REFERENCE_INTEREST_ID) AS COMMON_INTEREST
	FROM list_interest li2
	JOIN profile p ON li2.PROFILE_ID = p.ID
	WHERE li2.PROFILE_ID != ? 
	AND li2.REFERENCE_INTEREST_ID IN (
		SELECT li1.REFERENCE_INTEREST_ID
		FROM list_interest li1
		WHERE li1.PROFILE_ID = ?
	)
	GROUP BY li2.PROFILE_ID
	ORDER BY COMMON_INTEREST DESC
	LIMIT ?
	`

	stmt := repo.dbConn.Model(ProfileModel{}).Raw(rawQuery, query.ProfileID, query.ProfileID, query.Limit)
	if err := stmt.Find(&profileModel).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, err
		}
		return nil, err
	}

	return ListProfile(profileModel).ToEntities(), nil
}
