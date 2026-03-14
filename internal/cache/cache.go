// Package cache
package cache

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nicoki2004/d2-internal/internal/logger"
)

type Manager interface {
	Load(filename string, target any) error
	Save(filename string, data any) error
	GetChacheFileName() string
}

type FileCache struct{}

func NewFileCache() *FileCache {
	return &FileCache{}
}

// Load - try to load cache from a file.
func (fc *FileCache) Load(filename string, target any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cache file not found: %w", err)
	}

	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("error unmarshaling cache: %w", err)
	}

	return nil
}

// Save - Saves data into a file, filename
func (fc *FileCache) Save(filename string, data any) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		logger.GetLogger().Error("Error serializando caché: %v", err)
		return err
	}

	err = os.WriteFile(filename, jsonData, 0o644)
	if err != nil {
		logger.GetLogger().Error("Error writing cache file: %v", err)
		return err
	}

	return nil
}

// GetChacheFileName ...
func GetChacheFileName() string {
	if path := os.Getenv("CACHE_FILE_NAME"); path != "" {
		return path
	}
	return "profile_cache.json"
}
