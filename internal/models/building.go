package models

type Building struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longtitude  float64 `json:"longtitude"`
	VideoURL    string  `json:"video_url"`
}
