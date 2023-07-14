package config

import "path/filepath"

type Config struct {
	StorageFile string
}

func GetDefault() Config {
	storageFile, _ := filepath.Abs("./storage/gopher.json")

	return Config{
		StorageFile: storageFile,
	}
}
