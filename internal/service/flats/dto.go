package flats

type CreateFlatIn struct {
	HouseID int64
	Price   int64
	Rooms   int
	Number  int
}

type CreateFlatOut struct {
	ID      int64
	HouseID int64
	Price   int64
	Rooms   int
	Status  string
	Number  int
}
