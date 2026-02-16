package models

type Building struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string  `gorm:"not null" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Picture string `gormL"type:text" json:"picture"`
	Latitude    float64 `gorm:"type:float8" json:"latitude"`
	Longitude   float64 `gorm:"type:float8" json:"longitude"`
	VideoURL    string  `gorm:"type:text" json:"video_url"`
}
