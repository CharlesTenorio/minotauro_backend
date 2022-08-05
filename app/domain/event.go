package domain

import (
	"errors"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Event struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	EventDate   time.Time `json:"Data do evento"`
	StatusEvent string    `json:"status do evento"`
	Active      bool      `json:"activate"`
	ImgEvent    string    `json:"img"`
	CreateAt    time.Time `json:"data"`
}

type EventRepository interface {
	GetById(id string) (Event, error)
	Create(Event Event) (Event, error)
	Update(Event Event) (Event, error)
	Delete(id string) (string, error)
	FindAll() ([]Event, error)
}

func NewEvent() *Event {
	event := Event{
		Id:       uuid.NewV4().String(),
		CreateAt: time.Now(),
		Active:   true,
	}

	return &event
}

func (e *Event) validade() error {
	if e.Name == "" {
		return errors.New("O Nome o brigatório")
	}

	if e.EventDate.IsZero() {
		return errors.New("Data do evento é o brigatória")
	}

	return nil

}

func (event *Event) formatSpace() {
	event.Name = strings.TrimSpace(event.Name)

}
