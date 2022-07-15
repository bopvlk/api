package api

import (
	"github.com/BO/6.ServerAndDB/storage"
	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
)

//хочемо зконфігурувати наш API инстанс (поле logger)
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

//Хочемо відконфыгурувати маршрутизатор (поле Router API)
func (api *API) configureRouterField() {
	api.router.HandleFunc(prefix+"/articles", api.GetAllArticles).Methods("GET")
	api.router.HandleFunc(prefix+"/articles/{id}", api.GetArticlesById).Methods("GET")
	api.router.HandleFunc(prefix+"/articles/{id}", api.DeleteArticleById).Methods("DELETE")
	api.router.HandleFunc(prefix+"/articles", api.PostArticle).Methods("POST")
	api.router.HandleFunc(prefix+"/user/register", api.PostUserRegister).Methods("POST")
}

//хочемо зконфігурувати наш Stogage
func (a *API) configureStorageField() error {
	storage := storage.New(a.config.Storage)
	//хочемо створити зв'язок, якщо не мож вертажмо помилку.
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
