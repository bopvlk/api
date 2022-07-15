package api

import (
	"encoding/json"
	"net/http"

	"github.com/BO/6.ServerAndDB/internal/app/models"
)

//full API handler initialization file
//помогающа структура для формування повідомлень
type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func (api *API) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	//ініціалізуємо хедери
	initHeaders(w)
	//логіруємо початок опрацювання запита
	api.logger.Info("Get All Articles GET /api/v1/articles")
	//хочемо щось получити від бд
	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		// що ми робимо коли помилка на етапы пыдключення
		api.logger.Info("Error while Articles.SelectAll:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some trobles in database. Try agane later",
			IsError:    true,
		}
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(articles)
}

func (api *API) GetArticlesById(w http.ResponseWriter, r *http.Request) {

}

func (api *API) DeleteArticleById(w http.ResponseWriter, r *http.Request) {

}

func (api *API) PostArticle(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Post Article POST /api/v1/articles")
	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		api.logger.Info("Invalid  json resived from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	a, err := api.storage.Article().Create(&article)
	if err != nil {
		api.logger.Info("Trobles while creating new article", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troble is database. Try again later",
			IsError:    true,
		}
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(msg)
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(a)
}

func (api *API) PostUserRegister(w http.ResponseWriter, r *http.Request) {

}
