package domain

import "time"

type User struct {
	ID         string // UUID
	Email      string
	Password   string // Store hashed passwords
	Name       string
	ProfilePic string // URL to profile picture
	Bio        string
}

type PasswordResetRequest struct {
	ID        string // UUID
	UserID    string // Foreign key to User
	Token     string // Reset token
	ExpiresAt time.Time
}
