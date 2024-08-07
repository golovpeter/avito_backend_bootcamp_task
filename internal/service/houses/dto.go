package houses

import "time"

type CreateHouseIn struct {
	Address   string
	Year      int
	Developer string
}

type CreateHouseOut struct {
	ID        int64
	Address   string
	Year      int
	Developer string
	CreatedAt time.Time
	UpdatedAt time.Time
}
