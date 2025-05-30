package model

type Interest struct {
	InterestID uint       `gorm:"primaryKey;autoIncrement" json:"interest_id"`
	Name       string     `gorm:"size:100;not null" json:"name"`
	Residents  []Resident `gorm:"many2many:resident_interest;" json:"residents,omitempty"`
}

func (Interest) TableName() string {
	return "interest"
}
