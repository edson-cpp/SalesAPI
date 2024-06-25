package domain

import (
	"errors"
	"github.com/google/uuid"
)
var (
	ErrInvalidSpotNumber     = errors.New("Invalid spot number")
	ErrSpotNotFound          = errors.New("Spot not found")
	ErrSpotAlreadyReserved   = errors.New("Spot already reserved")
	ErrSpotNameTwoCharacters = errors.New("Spot name must be at least 2 characters long")
	ErrSpotNameRequired		 = errors.New("Spot name is required")
	ErrSpotStartWithLetter	 = errors.New("Spot name must start with a letter")
	ErrSpotEndWithNumber	 = errors.New("Spot name must end with a number")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	Id       string
	EventId  string
	Name     string
	Status   SpotStatus
	TicketId string
}

func NewSpot(event *Event, name strin) (*Spot, error) {
	spot := &Spot{
		Id: uuid.New().String(),
		EventId: event.Id,
		Name: name,
		Status: SpotStatusAvailable,
	}

	if spot.Validate() != nil {
		return nil, spot.Validate()
	}

	return spot, nil
}

func (s *Spot) Validate() error (
	if s.Name == "" {
		return ErrSpotNameRequired
	}

	if len(s.Name) < 2 {
		return ErrSpotNameTwoCharacters
	}

	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotStartWithLetter
	}

	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotEndWithNumber
	}

	return nil
)

func (s *Spot) Reserve(ticketId string) error {
	if s.Status == SpotStatusSold {
		return ErrSpotAlreadyReserved
	}
	s.Status = SpotStatusSold
	s.TicketId = ticketId
	return nil
}