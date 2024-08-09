package get_flats

type GetFlatsIn struct {
	HouseID int64 `json:"house_id"`
}

type GetFlatsOut struct {
	Flats []*FlatData `json:"flats"`
}

type FlatData struct {
	ID      int64  `json:"id"`
	HouseID int64  `json:"house_id"`
	Price   int64  `json:"price"`
	Rooms   int    `json:"rooms"`
	Number  int    `json:"number"`
	Status  string `json:"status"`
}
