package model

import "time"

type Event struct {
	EventID             uint       `gorm:"primaryKey;autoIncrement" json:"event_id,omitempty"`
	EventName           string     `gorm:"size:100;not null" json:"event_name"`
	EventDate           time.Time  `gorm:"type:date;not null" json:"event_date"`
	Address             string     `gorm:"size:200;not null" json:"address"`
	EventTypeName       string     `gorm:"not null" json:"event_type_name,omitempty"`
	OrganizerResidentID uint       `gorm:"not null" json:"organizer_resident_id,omitempty"`
	Organizer           Resident   `gorm:"foreignKey:OrganizerResidentID" json:"organizer"`
	Participants        []Resident `gorm:"many2many:event_participation;" json:"participants,omitempty"`
}

func (Event) TableName() string {
	return "event"
}
