package models

// SpokenLanguage represents a language such as English, Spanish, etc.
type SpokenLanguage struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Name         string `json:"name" gorm:"not null"`
	Abbreviation string `json:"abbreviation" gorm:"not null"`
	Creation     string `json:"creation"`
}

// CreateSpokenLanguageInput type with automatic ID.
type CreateSpokenLanguageInput struct {
	Name         string `json:"name" binding:"required"`
	Abbreviation string `json:"abbreviation" binding:"required"`
}
