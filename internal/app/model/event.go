package model

import "time"

type Event struct {
	EventID             uint       `gorm:"primaryKey;autoIncrement" json:"event_id"`
	EventName           string     `gorm:"size:100;not null" json:"event_name"`
	EventDate           time.Time  `gorm:"type:date;not null" json:"event_date"`
	Address             string     `gorm:"size:200;not null" json:"address"`
	EventTypeID         uint       `gorm:"not null" json:"event_type_id"`
	OrganizerResidentID uint       `gorm:"not null" json:"organizer_resident_id"`
	EventType           EventType  `gorm:"foreignKey:EventTypeID" json:"event_type"`
	Organizer           Resident   `gorm:"foreignKey:OrganizerResidentID" json:"organizer"`
	Participants        []Resident `gorm:"many2many:event_participation;" json:"participants,omitempty"`
}

func (Event) TableName() string {
	return "event"
}
