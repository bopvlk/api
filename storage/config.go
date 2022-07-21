package storage

type ConfigDB struct {
	//строка підключення до БД
	DatabaseURI string `toml:"database_uri"`
}

func NewConfigDB() *ConfigDB {
	return &ConfigDB{}
}
