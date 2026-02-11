package models

type Progress struct {
	ID         int  `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     int  `gorm:"not null;index" json:"user_id"`
	BuildingID int  `gorm:"not null;index" json:"building_id"`
	Flag       bool `gorm:"default:false" json:"flag"`
}
