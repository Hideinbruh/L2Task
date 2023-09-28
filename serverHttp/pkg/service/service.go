package service

import (
	"awesomeProject2/serverHttp"
	"awesomeProject2/serverHttp/pkg/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{repo: repository}
}

func (s *Service) CreateEvent(event *serverHttp.Event) (int, error) {
	return s.repo.CreateEvent(event)
}

func (s *Service) UpdateEvent(event *serverHttp.Event) error {
	return s.repo.UpdateEvent(event)
}

func (s *Service) DeleteEvent(eventId int) error {
	return s.repo.DeleteEvent(eventId)
}

func (s *Service) EventsForDay(event *serverHttp.Event) ([]serverHttp.Event, error) {
	return s.repo.EventsForDay(event)
}

func (s *Service) EventsForWeek(event *serverHttp.Event) ([]serverHttp.Event, error) {
	return s.repo.EventsForWeek(event)
}

func (s *Service) EventsForMonth(event *serverHttp.Event) ([]serverHttp.Event, error) {
	return s.repo.EventsForMonth(event)
}
