package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadYAML(path string, cfg *PostgresConfig) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		if e := file.Close(); err == nil {
			err = e
		}
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return err
	}

	return nil
}
