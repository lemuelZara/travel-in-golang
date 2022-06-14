package basics

import (
	"encoding/json"
	"os"
)

type Permissions struct {
	Read  bool `json:"read"`
	Write bool `json:"write"`
}

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Permissions `json:"permissions"`
}

func GetAll() (users []User) {
	file, err := os.ReadFile("../tmp/users.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &users)
	if err != nil {
		panic(err)
	}

	return users
}

func Add(user User) int {
	users := GetAll()

	users = append(users, user)

	data, _ := json.Marshal(users)

	_ = os.WriteFile("../tmp/users.json", data, 0644)

	return user.ID
}
