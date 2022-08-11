package models

type Item struct {
	ChrtID      int64  `json:"chrt_id" faker:"boundary_start=100, boundary_end=10000"`
	TrackNumber string `json:"track_number" faker:"len=20"`
	Price       int64  `json:"price" faker:"boundary_start=100, boundary_end=10000"`
	Rid         string `json:"rid" faker:"len=20"`
	Name        string `json:"name" faker:"first_name"`
	Sale        int64  `json:"sale" faker:"boundary_start=0, boundary_end=100"`
	Size        string `json:"size" faker:"oneof: 0"`
	TotalPrice  int64  `json:"total_price" faker:"boundary_start=50, boundary_end=10000"`
	NmID        int64  `json:"nm_id" faker:"boundary_start=1000, boundary_end=1000000"`
	Brand       string `json:"brand" faker:"word"`
	Status      int64  `json:"status" faker:"boundary_start=0, boundary_end=500"`
}
