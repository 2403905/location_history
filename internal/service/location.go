package service

import (
	"github.com/2403905/location_history/internal/model"
)

type LocationRepo interface {
	AddLocation(orderId string, location model.Location)
	GetLocation(orderId string) ([]model.Location, error)
	DeleteLocation(orderId string) error
}

type Location struct {
	repo LocationRepo
}

func NewLocation(repo LocationRepo) Location {
	return Location{
		repo: repo,
	}
}

func (s *Location) AppendLocation(orderId string, location model.Location) {
	s.repo.AddLocation(orderId, location)
}

func (s *Location) GetLocation(orderId string, max int) ([]model.Location, error) {
	locationList, err := s.repo.GetLocation(orderId)
	if err != nil {
		return nil, err
	}
	if max == 0 {
		max = len(locationList)
	}
	res := make([]model.Location, 0, max)
	for i := len(locationList); i > max; i-- {
		res = append(res, locationList[i])
	}
	return res, err
}

func (s *Location) DeleteLocation(orderId string) error {
	return s.repo.DeleteLocation(orderId)
}
