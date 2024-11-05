package profile

import "github.com/asmaulh99/dating-app-backend/entities"

type ProfileModel struct {
	ID             uint   `gorm:"column:ID"`
	Name           string `gorm:"column:NAME"`
	Bio            string `gorm:"column:BIO"`
	PictureUrl     string `gorm:"column:PICTURE_URL"`
	Age            uint   `gorm:"column:AGE"`
	Gender         string `gorm:"column:GENDER"`
	CommonInterest *uint  `gorm:"column:COMMON_INTEREST;->"`
}

func (model *ProfileModel) ToEntity() *entities.Profile {
	return &entities.Profile{
		ID:         model.ID,
		Name:       model.Name,
		Bio:        model.Bio,
		PictureUrl: model.PictureUrl,
		Age:        model.Age,
		Gender:     model.Gender,
		// CommonInterest: model.CommonInterest,
	}
}

type ListProfile []*ProfileModel

func (models ListProfile) ToEntities() []*entities.Profile {
	var entities = make([]*entities.Profile, len(models))
	for i, v := range models {
		entities[i] = v.ToEntity()
	}
	return entities
}
