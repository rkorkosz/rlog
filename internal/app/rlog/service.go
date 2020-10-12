package rlog

import (
	"io"
)

type Storage interface {
	Get(string) (*Entry, error)
	Put(Entry) error
	Delete(string) error
	List() ([]Entry, error)
}

type Renderer func(io.Writer, interface{}) error

type Service struct {
	storage  Storage
	renderer Renderer
	out      io.Writer
}

func NewService(out io.Writer, s Storage, r Renderer) *Service {
	return &Service{
		storage:  s,
		renderer: r,
		out:      out,
	}
}

func (s *Service) List() error {
	entries, err := s.storage.List()
	if err != nil {
		return err
	}
	return s.renderer(s.out, &entries)
}

func (s *Service) Get(slug string) error {
	entry, err := s.storage.Get(slug)
	if err != nil {
		return err
	}
	return s.renderer(s.out, entry)
}

func (s *Service) Delete(slug string) error {
	return s.storage.Delete(slug)
}

func (s *Service) Put(e Entry) error {
	err := s.storage.Put(e)
	if err != nil {
		return err
	}
	return s.renderer(s.out, &e)
}
