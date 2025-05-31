package repository

import (
	"Filo.Hack/internal/app/model"
	"gorm.io/gorm"
)

type EventRepositoryI interface {
	Add(event model.Event)
	GetAll() []model.Event
	GetAllByDate(date string) []model.Event
	GetAllByCategory(category string) []model.Event
	GetAllByAddress(address string) []model.Event
}

type EventRepository struct {
	DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{DB: db}
}

func (r *EventRepository) Add(event *model.Event) error {
	return r.DB.Create(event).Error
}

func (r *EventRepository) GetAll() []model.Event {
	var events []model.Event
	r.DB.Find(&events)
	return events
}

func (r *EventRepository) GetAllByDate(date string) []model.Event {
	var events []model.Event
	r.DB.Where("date = ?", date).Find(&events)
	return events
}

func (r *EventRepository) GetAllByCategory(category string) []model.Event {
	var events []model.Event
	r.DB.Where("category = ?", category).Find(&events)
	return events
}

func (r *EventRepository) GetAllByAddress(address string) []model.Event {
	var events []model.Event
	r.DB.Where("address = ?", address).Find(&events)
	return events
}

func (r *EventRepository) GetOrCreateInterest(name string) (*model.Interest, error) {
	var interest model.Interest
	result := r.DB.Where(model.Interest{Name: name}).FirstOrCreate(&interest)

	if result.Error != nil {
		return nil, result.Error
	}

	return &interest, nil
}

func (r *EventRepository) GetAllInterests() (*[]model.Interest, error) {
	var interests []model.Interest
	result := r.DB.Find(&interests)
	if result.Error != nil {
		return nil, result.Error
	}
	return &interests, nil
}
