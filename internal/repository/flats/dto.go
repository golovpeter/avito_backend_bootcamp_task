package flats

type InsertNewFlatIn struct {
	HouseID int64
	Price   int64
	Rooms   int
	Number  int
}

type FlatData struct {
	ID      int64  `db:"id"`
	HouseID int64  `db:"house_id"`
	Price   int64  `db:"price"`
	Rooms   int    `db:"rooms"`
	Status  string `db:"status"`
	Number  int    `db:"number"`
}

type UpdateFlatIn struct {
	ID     int64  `db:"id"`
	Status string `db:"status"`
}
