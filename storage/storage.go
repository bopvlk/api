package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//instance of storage
type Storage struct {
	config *Config
	//DataBase file discriptor
	db *sql.DB
	//subfield for repo interfasing (model user)
	userRepository *UserRepository
	//subfield for repo interfasing (model article)
	articleRepository *ArticleRepository
}

//storage constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

//open conection method
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Database conections created successfuly")
	return nil
}

func (storage *Storage) Close() {
	storage.db.Close()
}

//public repo for Article
func (s *Storage) Article() *ArticleRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}
	s.articleRepository = &ArticleRepository{
		storage: s,
	}
	return s.articleRepository
}

//public repo for User
func (s *Storage) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{storage: s}
	return s.userRepository
}
