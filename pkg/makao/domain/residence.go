package domain

type Residence struct {
	ID                 string `gorm:"column:id"`
	Active             bool   `gorm:"column:active"`
	Name               string `gorm:"column:name"`
	RegistrationNumber string `gorm:"column:registration_number"`
	Location           string `gorm:"column:location"`
	LivingRoomsCount   int    `gorm:"column:living_rooms_count"`
	Owner              string `gorm:"column:owner"`
}

// Proceed with Residence

// Residence home page
