package domain

import "time"

type User struct {
	ID        string `json:"id"`
  Name      string `json:"name"`
  Email     string `json:"email"`
  Avater    string `json:"avater"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
