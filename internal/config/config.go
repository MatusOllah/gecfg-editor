package config

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io"
	"log/slog"
	"maps"
	"os"
	"path/filepath"
	"sync"
)

var ErrNotFound error = errors.New("key not found")

type MapFunc func(string, interface{}) interface{}

type Config struct {
	dataLock sync.RWMutex
	data     map[string]interface{}

	file    *os.File
	buf     *bytes.Buffer
	encoder *gob.Encoder
	decoder *gob.Decoder
}

// Register simply calls gob.Register. If you are encoding a non-primitive type (like a struct or map) that implements something you should use this.
func Register(value interface{}) {
	gob.Register(value)
}

// New creates / opens and decodes a new Config.
func New(path string) (*Config, error) {
	if exists(path) {
		cfg, err := Open(path)
		if err != nil {
			return nil, err
		}

		return cfg, nil
	} else {
		cfg, err := Create(path)
		if err != nil {
			return nil, err
		}

		return cfg, nil
	}
}

// Create creates a new Config.
func Create(path string) (*Config, error) {
	cfg := new(Config)

	// creates file
	// vytvori subor
	os.Mkdir(filepath.Dir(path), os.ModePerm)

	file, _ := os.Create(path)

	cfg.data = make(map[string]interface{})
	cfg.file = file
	cfg.buf = new(bytes.Buffer)
	cfg.encoder = gob.NewEncoder(cfg.buf)
	cfg.decoder = gob.NewDecoder(cfg.buf)

	return cfg, nil
}

// Open opens a existing Config.
func Open(path string) (*Config, error) {
	cfg := new(Config)

	// opens file
	// otvori subor
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	cfg.file = file
	cfg.buf = new(bytes.Buffer)
	cfg.encoder = gob.NewEncoder(cfg.buf)
	cfg.decoder = gob.NewDecoder(cfg.buf)

	// decodes raw data
	// dekoduje surove data
	_, err = io.Copy(cfg.buf, cfg.file)
	if err != nil {
		return nil, err
	}

	cfg.dataLock.Lock()
	if err := cfg.decoder.Decode(&cfg.data); err != nil {
		return nil, err
	}
	cfg.dataLock.Unlock()

	return cfg, nil
}

// Data returns a copy of the map.
func (cfg *Config) Data() map[string]interface{} {
	cfg.dataLock.RLock()
	data := cfg.data
	cfg.dataLock.RUnlock()
	return data
}

// SetData overwrites the map.
func (cfg *Config) SetData(m map[string]interface{}) {
	cfg.dataLock.RLock()
	cfg.data = m
	cfg.dataLock.RUnlock()
}

// Append appends m to the map.
func (cfg *Config) Append(m map[string]interface{}) {
	cfg.dataLock.Lock()
	maps.Copy(cfg.data, m)
	cfg.dataLock.Unlock()
}

// Get gets a value from the map.
func (cfg *Config) Get(key string) (interface{}, error) {
	cfg.dataLock.RLock()
	value, ok := cfg.data[key]
	if !ok {
		return nil, ErrNotFound
	}
	cfg.dataLock.RUnlock()
	return value, nil
}

// MustGet simply calls Get and returns nil if an error occured.
func (cfg *Config) MustGet(key string) interface{} {
	value, err := cfg.Get(key)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	return value
}

// GetWithFallback gets a value from the map and returns the given fallback if not found.
func (cfg *Config) GetWithFallback(key string, fallback interface{}) interface{} {
	cfg.dataLock.RLock()
	value, ok := cfg.data[key]
	if !ok {
		slog.Warn("key not found", "key", key, "fallback", fallback)
		return fallback
	}
	cfg.dataLock.RUnlock()
	return value
}

// Set sets key to value.
func (cfg *Config) Set(key string, value interface{}) {
	cfg.dataLock.Lock()
	cfg.data[key] = value
	cfg.dataLock.Unlock()
}

// Delete deletes key from the map.
func (cfg *Config) Delete(key string) {
	cfg.dataLock.Lock()
	delete(cfg.data, key)
	cfg.dataLock.Unlock()
}

// Toggle toggles a bool value.
func (cfg *Config) Toggle(key string) {
	cfg.dataLock.Lock()
	cfg.data[key] = !cfg.data[key].(bool)
	cfg.dataLock.Unlock()
}

// Exists checks if key exists.
func (cfg *Config) Exists(key string) bool {
	cfg.dataLock.RLock()
	_, ok := cfg.data[key]
	cfg.dataLock.RUnlock()
	return ok
}

// Map iterates over the map and applies the MapFunc to every item.
func (cfg *Config) Map(fn MapFunc) {
	cfg.dataLock.Lock()
	for k, v := range cfg.data {
		cfg.data[k] = fn(k, v)
	}
	cfg.dataLock.Unlock()
}

// Wipe wipes (clears) the map.
func (cfg *Config) Wipe() {
	cfg.dataLock.Lock()
	clear(cfg.data)
	cfg.dataLock.Unlock()
}

// Flush gob encodes and writes data to the file.
func (cfg *Config) Flush() error {
	cfg.dataLock.RLock()
	if err := cfg.encoder.Encode(cfg.data); err != nil {
		return err
	}
	cfg.dataLock.RUnlock()

	_, err := cfg.file.WriteAt(cfg.buf.Bytes(), 0)
	if err != nil {
		return err
	}

	cfg.file.Sync()

	return nil
}

// Close simply calls (*os.File).Close and closes the file.
func (cfg *Config) Close() error {
	return cfg.file.Close()
}
