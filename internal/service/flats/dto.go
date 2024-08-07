package flats

type CreateFlatIn struct {
	HouseID int64
	Price   int64
	Rooms   int
	Number  int
}

type FlatData struct {
	ID      int64
	HouseID int64
	Price   int64
	Rooms   int
	Status  string
	Number  int
}

type UpdateFlatIn struct {
	ID     int64
	Status string
}
