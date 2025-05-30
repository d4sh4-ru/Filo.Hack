package model

type ResidentInterest struct {
	ResidentID uint     `gorm:"primaryKey" json:"resident_id"`
	InterestID uint     `gorm:"primaryKey" json:"interest_id"`
	Resident   Resident `gorm:"foreignKey:ResidentID" json:"resident"`
	Interest   Interest `gorm:"foreignKey:InterestID" json:"interest"`
}

func (ResidentInterest) TableName() string {
	return "resident-interest"
}
