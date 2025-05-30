package model

type Resident struct {
	ResidentID      uint        `gorm:"primaryKey;autoIncrement" json:"resident_id"`
	FirstName       string      `gorm:"size:50;not null" json:"first_name"`
	LastName        string      `gorm:"size:50;not null" json:"last_name"`
	MiddleName      *string     `gorm:"size:50" json:"middle_name,omitempty"`
	Age             *int        `json:"age,omitempty"`
	HouseNumber     string      `gorm:"size:10;not null" json:"house_number"`
	Entrance        *int        `json:"entrance,omitempty"`
	Apartment       *int        `json:"apartment,omitempty"`
	Interests       []Interest  `gorm:"many2many:resident_interest;" json:"interests,omitempty"`
	OrganizedEvents []Event     `gorm:"foreignKey:OrganizerResidentID" json:"organized_events,omitempty"`
	Shows           []EventShow `gorm:"foreignKey:ResidentID" json:"shows,omitempty"`
	Participations  []Event     `gorm:"many2many:event_participation;" json:"participations,omitempty"`
}

func (Resident) TableName() string {
	return "resident"
}
