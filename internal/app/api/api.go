package api

import (
	"net/http"

	"github.com/BO/6.ServerAndDB/storage"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//базовий апі опис екземпляра
//Base API Server instance description
type API struct {
	//UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
	router *mux.Router
	//добавка поля для роботи зі Storage
	storage *storage.Storage
}

//API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Start http server/configure logger, router, database conection and etc...
func (api *API) Start() error {
	//truing to configure logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	//підтвердження того, що логгер зконфігурований
	api.logger.Info("starting api server at port:", api.config.BindAddr)

	//конфігуруємо наш маршрутизатор
	api.configureRouterField()
	//конфігуруємо сховище (Storage)
	if err := api.configureStorageField(); err != nil {
		return err
	}
	//на етапі валідного завершения стартуємо http-сервер
	return http.ListenAndServe(api.config.BindAddr, api.router)

}
