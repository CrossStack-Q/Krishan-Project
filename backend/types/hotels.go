package types

type Hotel struct {
	HotelID  int      `json:"hotel_id"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Rating   float64  `json:"rating"`
	Landmark string   `json:"landmark"`
	Images   []string `json:"images"`
	Reviews  []Review `json:"reviews"`
}
type Review struct {
	UserName string  `json:"userName"`
	Content  string  `json:"content"`
	Rating   float64 `json:"rating"`
}
