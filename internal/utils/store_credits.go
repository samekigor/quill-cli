package utils

import (
	"encoding/base64"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	creditsFilePath = "/etc/quill/"
	creditsFileName = "credentials.yml"
)

var (
	creditsFullPath = fmt.Sprintf("%s%s", creditsFilePath, creditsFileName)
)

type RegistryEntry struct {
	Registry string `yaml:"registry"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (r *RegistryEntry) EncodePassword() {
	r.Password = base64.StdEncoding.EncodeToString([]byte(r.Password))
}

func saveToYAML(entries []RegistryEntry) error {
	data, err := yaml.Marshal(entries)
	if err != nil {
		return err
	}
	return os.WriteFile(creditsFullPath, data, 0755)
}
func loadFromYAML() ([]RegistryEntry, error) {
	data, err := os.ReadFile(creditsFullPath)
	if err != nil {
		return nil, err
	}

	var entries []RegistryEntry
	err = yaml.Unmarshal(data, &entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func GetRegistryEntry(entries []RegistryEntry, registryName string) *RegistryEntry {
	for _, entry := range entries {
		if entry.Registry == registryName {
			return &entry
		}
	}
	return nil
}
func (r *RegistryEntry) AddRegistryEntry() error {
	entries, err := loadFromYAML()
	if err != nil {
		return err
	}

	for i, entry := range entries {
		if entry.Registry == r.Registry {
			entries[i] = *r
			return saveToYAML(entries)
		}
	}

	entries = append(entries, *r)
	return saveToYAML(entries)
}

func (r *RegistryEntry) RemoveRegistryEntry() error {
	entries, err := loadFromYAML()
	if err != nil {
		return err
	}

	var updatedEntries []RegistryEntry
	for _, entry := range entries {
		if entry.Registry != r.Registry {
			updatedEntries = append(updatedEntries, entry)
		}
	}
	entries = updatedEntries

	return saveToYAML(entries)
}
