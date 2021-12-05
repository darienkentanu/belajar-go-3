package repositories

import (
	"belajar-go-3/mvc/models"
	"time"
)

var users []models.User

func init() {
	users = append(users, models.User{ID: 1, Name: "joss fong"})
	users = append(users, models.User{ID: 2, Name: "clairine", Email: "clairine@gmail.com"})
}

func GetUsers() []models.User {
	return users
}

func CreateUser(user models.User) models.User {
	user.ID = int64(len(users) + 1)
	user.CreatedAt = time.Now()

	users = append(users, user)
	return user
}

func GetUser(id int64) models.User {
	for _, v := range users {
		if v.ID == id {
			return v
		}
	}

	return models.User{}
}
