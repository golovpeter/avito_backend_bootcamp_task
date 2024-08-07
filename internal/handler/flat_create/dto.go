package flat_create

type CreateFlatIn struct {
	HouseID int64 `json:"house_id"`
	Price   int64 `json:"price"`
	Rooms   int   `json:"rooms"`
	Number  int   `json:"number"`
}

type CreateFlatOut struct {
	ID      int64  `json:"id"`
	HouseID int64  `json:"house_id"`
	Price   int64  `json:"price"`
	Rooms   int    `json:"rooms"`
	Status  string `json:"status"`
	Number  int    `json:"number"`
}
