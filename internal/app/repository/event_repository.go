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

func (r *EventRepository) Add(event model.Event) {
	r.DB.Create(&event)
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
