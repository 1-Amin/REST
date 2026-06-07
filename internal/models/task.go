package models

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey"     json:"id"`
	Title       string    `gorm:"size:200"       json:"title"`
	Description string    `gorm:"type:text"      json:"description"`
	Done        bool      `gorm:"default:false"  json:"done"`
	OwnerID     uint      `gorm:"index"          json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTaskDTO struct {
	Title       string `json:"title"       binding:"required,min=3,max=200"`
	Description string `json:"description" binding:"max=2000"`
}

type UpdateTaskDTO struct {
	Title       *string `json:"title,omitempty"       binding:"omitempty,min=3,max=200"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=2000"`
	Done        *bool   `json:"done,omitempty"`
}
