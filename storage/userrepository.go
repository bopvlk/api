package storage

import (
	"fmt"
	"log"

	"github.com/BO/6.ServerAndDB/internal/app/models"
)

//instance of article repository
type UserRepository struct {
	storage *Storage
}

var (
	tableUser string = "users"
)

//Create user
func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password) VALUES ($1, $2) RETURNING id", tableUser)
	if err := ur.storage.db.QueryRow(query, u.Login, u.Password).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

//Find user by login
func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	users, err := ur.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var userFinded *models.User
	for _, u := range users {
		if u.Login == login {
			userFinded = u
			founded = true
			break
		}
	}
	return userFinded, founded, nil
}

//select users in db
func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//підгтуємо куди будемо записувати наших юзерів
	users := make([]*models.User, 0)
	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.ID, &u.Login, &u.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return users, nil
}

//шукати користувача по id
func (ur *UserRepository) FindUserById(id int) (*models.User, bool, error) {
	users, err := ur.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var usersFounded *models.User
	for _, a := range users {
		if a.ID == id {
			usersFounded = a
			founded = true
			break
		}
	}
	return usersFounded, founded, nil
}

// Видалити користувача по ID
func (ur *UserRepository) DeleteUser(id int) (*models.User, error) {
	user, ok, err := ur.FindUserById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tableUser)
		_, err := ur.storage.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}
