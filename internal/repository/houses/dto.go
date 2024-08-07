package houses

import "time"

type InsertNewHouseIn struct {
	Address   string
	Year      int
	Developer string
}

type InsertNewHouseOut struct {
	ID        int64     `db:"id"`
	Address   string    `db:"address"`
	Year      int       `db:"year"`
	Developer string    `db:"developer"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
