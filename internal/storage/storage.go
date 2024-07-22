package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

type Storage struct {
	filename string
	mu       sync.Mutex
}

type Entry struct {
	Service  string `json:"service"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func New(filename string) *Storage {
	return &Storage{filename: filename}
}

func (s *Storage) Save(entries []Entry) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := json.Marshal(entries)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(s.filename, data, 0600)
}

func (s *Storage) Load() ([]Entry, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := ioutil.ReadFile(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Entry{}, nil
		}
		return nil, err
	}

	var entries []Entry
	err = json.Unmarshal(data, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
