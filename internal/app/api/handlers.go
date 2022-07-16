package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BO/6.ServerAndDB/internal/app/models"
	"github.com/gorilla/mux"
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
	initHeaders(w)
	api.logger.Info("get article by id /api/v1/articles")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		api.logger.Info("Trobles by parsing [id] param", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. Don`t use ID as uncasting to int value",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	article, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Trobles while accessing database table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some trobles to accesing database. Try again later.",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Can not find article with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with that ID does not in database",
			IsError:    true,
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(article)
}

func (api *API) DeleteArticleById(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Delete article by ID DELETE /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		api.logger.Info("Trobles by parsing [id] param", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. Don`t use ID as uncasting to int value",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}

	_, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Trobles while accessing database table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some trobles to accesing database. Try again later.",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}

	if !ok {
		api.logger.Info("Can not find article with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with that ID does not in database",
			IsError:    true,
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(msg)
		return
	}

	_, err = api.storage.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("Trobles while deleting database elemett from table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some trobles to accesing database. Try again later.",
			IsError:    true,
		}
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("Article with ID %d successfuly deleted", id),
		IsError:    false,
	}
	json.NewEncoder(w).Encode(msg)
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
	initHeaders(w)
	api.logger.Info("Post User Register POST /api/v1/register")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
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
	//хочемо найти корислувача з таким логыном в бд
	_, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Trobles while accessing database table (users) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some trobles to accesing database. Try again later.",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}

	//дивимося, якщо такий користувач уже є, то регістрацію не робимо
	if ok {
		api.logger.Info("User with that ID already exists")
		msg := Message{
			StatusCode: 400,
			Message:    "User with that ID already existsin database",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	//тепер хочемо додати дані в БД
	newUser, err := api.storage.User().Create(&user)
	if err != nil {
		api.logger.Info("Trobles while accessing database table (users) with id, err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some trobles to accesing database. Try again later.",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("User with {login:%s} successfuly registred!", newUser.Login),
		IsError:    false,
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(msg)
}

func (api *API) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Delete user by ID, DELETE /api/v1/user/{id}")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		api.logger.Info("Trobles by parsing [id] param", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. Don`t use ID as uncasting to int value",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	_, ok, err := api.storage.User().FindUserById(id)
	if err != nil {
		api.logger.Info("Trobles while accessing database table (users) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some trobles to accesing database. Try again later.",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Can not find user with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "User with that ID does not in database",
			IsError:    true,
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(msg)
		return
	}
	user, err := api.storage.User().DeleteUser(id)
	if err != nil {
		api.logger.Info("Trobles while deleting database elemett from table (user) with id. err:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some trobles to accesing database. Try again later.",
			IsError:    true,
		}
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("User with ID: {%d} and Login: {%s} successfuly deleted", id, user.Login),
		IsError:    false,
	}
	json.NewEncoder(w).Encode(msg)
}
