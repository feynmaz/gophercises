package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/feynmaz/gophercises/exercise_3_adventure/models"
)

var (
	ErrOpenFile   = errors.New("failed to open storage file")
	ErrReadFile   = errors.New("failed to read storage file")
	ErrUnmarshall = errors.New("failed to unmarshall storage file content")

	ErrMarshall = errors.New("failed to marshall storage")
)

type Storage struct {
	Stories map[string]models.Story
}

func NewStorageFromFile(filepath string) (*Storage, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrOpenFile, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrReadFile, err)
	}

	var stories map[string]models.Story
	if err := json.Unmarshal(data, &stories); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrUnmarshall, err)
	}

	return &Storage{
		Stories: stories,
	}, nil
}

func (s *Storage) PrintContent() error {
	byteContent, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		return fmt.Errorf("%w: %w", ErrMarshall, err)
	}
	fmt.Println(string(byteContent))
	return nil
}

func (s *Storage) GetInitial() (*models.Story, error) {
	return s.GetById("intro")
}

func (s *Storage) GetById(storyID string) (*models.Story, error) {
	story, ok := s.Stories[storyID]
	if !ok {
		return &models.Story{}, nil
	}
	return &story, nil
}
