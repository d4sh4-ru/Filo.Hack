package model

type EventShow struct {
	ShowID      uint      `gorm:"primaryKey;autoIncrement" json:"show_id"`
	EventTypeID uint      `gorm:"not null" json:"event_type_id"`
	ResidentID  uint      `gorm:"not null" json:"resident_id"`
	EventType   EventType `gorm:"foreignKey:EventTypeID" json:"event_type"`
	Resident    Resident  `gorm:"foreignKey:ResidentID" json:"resident"`
}

func (EventShow) TableName() string {
	return "event-show"
}
