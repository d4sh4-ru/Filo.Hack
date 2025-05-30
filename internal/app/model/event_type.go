package model

type EventType struct {
	EventTypeID uint        `gorm:"primaryKey;autoIncrement" json:"event_type_id"`
	Name        string      `gorm:"size:100;not null" json:"name"`
	Events      []Event     `gorm:"foreignKey:EventTypeID" json:"events,omitempty"`
	Shows       []EventShow `gorm:"foreignKey:EventTypeID" json:"shows,omitempty"`
}

func (EventType) TableName() string {
	return "event-type"
}
