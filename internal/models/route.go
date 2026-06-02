package models

type Route struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `gorm:"not null;uniqueIndex" json:"name"`
	Description string     `gorm:"type:text" json:"description"`
	Buildings   []Building `gorm:"many2many:route_buildings;" json:"buildings,omitempty"`
}
