package psql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base model contains defines common fields across tables
type Base struct {
	CreatedAt time.Time `gorm:"column:created;not null"`
	UpdatedAt time.Time `gorm:"column:updated;not null"`
}

// SMSData is the model for the sms_data table
type SMSData struct {
	Base

	ID          string     `gorm:"column:id;primary_key"`
	LinkID      string     `gorm:"column:link_id"`
	Date        *time.Time `gorm:"column:date"`
	Sender      string     `gorm:"column:sender"`
	Text        string     `gorm:"column:text"`
	Recipient   string     `gorm:"column:recipient"`
	NetworkCode string     `gorm:"column:network_code"`
}

// BeforeCreate is a hook that is called before creating a record
func (s *SMSData) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New().String()
	return
}

// TableName returns the name of the table
func (SMSData) TableName() string {
	return "sms_data"
}
