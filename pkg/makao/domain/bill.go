package domain

// ADDITIONAL BILLS
type Bill struct {
	ID      string  `gorm:"column:id"`
	Active  bool    `gorm:"column:active"`
	Type    string  `gorm:"column:type"`
	Amount  float64 `gorm:"column:amount"`
	Penalty float64 `gorm:"column:penalty"`
	UserID  string  `gorm:"column:user_id"`
}

// Bill action (Staff)
// - add
// - update
// - delete
// - query

// // Bill action (tenant)
// - view
