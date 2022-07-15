package storage

type Config struct {
	//строка підключення до БД
	DatabaseURI string `toml:"database_uri"`
}

func NewConfig() *Config {
	return &Config{}
}
