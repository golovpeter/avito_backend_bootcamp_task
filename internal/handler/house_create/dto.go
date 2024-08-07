package house_create

import "time"

type CreateHouseIn struct {
	Address   string `json:"address"`
	Year      int    `json:"year"`
	Developer string `json:"developer"`
}

type CreateHouseOut struct {
	ID        int64     `json:"id"`
	Address   string    `json:"address"`
	Year      int       `json:"year"`
	Developer string    `json:"developer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
