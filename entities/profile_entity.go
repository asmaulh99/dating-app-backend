package entities

type GetProfileQuery struct {
	GenderInterest string
	MinAge         int
	MaxAge         int
	ProfileID      int
	Limit          int
}

type Profile struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Bio        string `json:"bio"`
	PictureUrl string `json:"pictureUrl"`
	Age        uint   `json:"age"`
	Gender     string `json:"gender"`
	// CommonInterest *uint  `json:"commonInterset"`
}
