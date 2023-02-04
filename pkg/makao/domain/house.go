package domain

// House details
type House struct {
	ID          string  `json:"id"`
	Active      bool    `json:"active"`
	Number      string  `json:"number"`
	Category    string  `json:"category"`
	Class       string  `json:"class"` // applicable where houses maybe charged differently due to size
	RentValue   float64 `json:"rentValue"`
	State       string  `json:"state"`
	ResidenceID string  `json:"residence_id"`
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
