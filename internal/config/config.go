package config

type PostgresConfig struct {
	host     string `yaml:"host"`
	port     int    `yaml:"port"`
	db       string `yaml:"db"`
	username string `yaml:"username"`
	password string `yaml:"password"`
}
