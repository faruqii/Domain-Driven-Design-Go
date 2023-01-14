package entities

type Movie struct {
	ID          string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseDate string `json:"release_date"`
	Rating      int    `json:"rating"`
}
