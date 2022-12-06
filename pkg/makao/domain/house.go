package domain

// House details
type House struct {
	ID          string  `gorm:"column:id;primary_key"`
	Active      bool    `gorm:"column:active"`
	HouseNumber string  `gorm:"column:house_number"`
	Category    string  `gorm:"column:category"`
	Class       string  `gorm:"column:class"` // applicable where houses maybe charged differently due to size
	RentValue   float64 `gorm:"column:rent_value"`
}

// House actions (admin)
// - add
// - update
// - query
// - delete

// - search house by number

// house search response
// - id
// - name
// - tenant id
// - house type
