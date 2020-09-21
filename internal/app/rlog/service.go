package rlog

import (
	"io"
	"log"

	"github.com/spf13/viper"
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

func NewService(out io.Writer) *Service {
	s, err := NewBoltStorage(viper.GetString("dbpath"), "entries")
	if err != nil {
		log.Fatal(err)
	}
	return &Service{
		storage:  s,
		renderer: RenderJSON,
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
