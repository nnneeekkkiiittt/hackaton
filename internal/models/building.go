package models

type Building struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string  `gorm:"not null" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Latitude    float64 `gorm:"type:float8" json:"latitude"`
	Longitude   float64 `gorm:"type:float8" json:"longitude"`
	VideoURL    string  `gorm:"type:text" json:"video_url"`
	Routes      []Route `gorm:"many2many:route_buildings;" json:"routes,omitempty"`
}
