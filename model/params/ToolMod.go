package params

import "time"

type UUID struct {
	ID        string     `gorm:"primary_key"`
	CreatedAt time.Time  `sortBy:"created_at"`
	UpdatedAt time.Time  `sortBy:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
}
