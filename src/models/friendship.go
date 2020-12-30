package models

// Friendship represents the type of friendship two users have.
type Friendship struct {
	ID       uint   `json:"id" gorm:"TINYINT; primaryKey"`
	Name     string `json:"name" gorm:"TEXT NOT NULL"`
	Modified string `json:"modified" gorm:"TIMESTAMP"`
}

// CreateFriendshipInput represents the type of friendship two users have.
type CreateFriendshipInput struct {
	Name string `json:"cost" binding:"required"`
}
