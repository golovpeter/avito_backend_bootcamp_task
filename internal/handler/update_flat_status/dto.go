package update_flat_status

type UpdateFlatStatusIn struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type UpdateFlatStatusOut struct {
	ID      int64  `json:"id"`
	HouseID int64  `json:"house_id"`
	Price   int64  `json:"price"`
	Rooms   int    `json:"rooms"`
	Number  int    `json:"number"`
	Status  string `json:"status"`
}
