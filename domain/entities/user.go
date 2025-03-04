package entities

import "time"

type User struct {
	ID       uint   `json:"primaryKey"`
	Name     string `json:"not null"`
	Email    string `json:"unique;not null"`
	Password string `json:"not null"`
	RoleID   uint   `json:"not null"`
	APIKey   string `json:"unique;not null"`
	Role     Role   `json:"foreignKey:RoleID"`
}

type UserToken struct {
	ID           uint      `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}
