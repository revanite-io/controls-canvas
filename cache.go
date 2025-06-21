package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/revanite-io/sci/layer2"
	"gopkg.in/yaml.v3"
)

const cacheDir = "tmp"

// getCacheFilename generates a unique cache filename based on the URLs
func getCacheFilename(urls []string) string {
	// Sort URLs to ensure consistent hashing
	urlStr := strings.Join(urls, "|")
	hash := sha256.Sum256([]byte(urlStr))
	return filepath.Join(cacheDir, "controls-canvas-"+hex.EncodeToString(hash[:8])+".yaml")
}

// ensureCacheDir creates the cache directory if it doesn't exist
func ensureCacheDir() error {
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		return os.MkdirAll(cacheDir, 0755)
	}
	return nil
}

// loadFromCache attempts to load catalog data from cache
func loadFromCache(urls []string) (*layer2.Catalog, error) {
	cacheFile := getCacheFilename(urls)
	data, err := os.ReadFile(cacheFile)
	if err != nil {
		return nil, err
	}

	var catalog layer2.Catalog
	err = yaml.Unmarshal(data, &catalog)
	if err != nil {
		return nil, err
	}

	return &catalog, nil
}

// saveToCache saves catalog data to cache
func saveToCache(urls []string, catalog *layer2.Catalog) error {
	if err := ensureCacheDir(); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	data, err := yaml.Marshal(catalog)
	if err != nil {
		return fmt.Errorf("failed to marshal catalog: %w", err)
	}

	cacheFile := getCacheFilename(urls)
	if err := os.WriteFile(cacheFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write cache file: %w", err)
	}

	return nil
}
