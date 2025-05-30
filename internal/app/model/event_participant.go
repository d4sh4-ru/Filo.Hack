package model

type EventParticipation struct {
	ResidentID uint     `gorm:"primaryKey" json:"resident_id"`
	EventID    uint     `gorm:"primaryKey" json:"event_id"`
	Resident   Resident `gorm:"foreignKey:ResidentID" json:"resident"`
	Event      Event    `gorm:"foreignKey:EventID" json:"event"`
}

func (EventParticipation) TableName() string {
	return "event-participation"
}
