package interfaces

import (
	"encoding/json"
	"net/http"
)

type UsersHandler struct{}

func NewUsersHandler() *UsersHandler {
	return &UsersHandler{}
}

func (uh *UsersHandler) Index(w http.ResponseWriter, r *http.Request) {

	users := []struct {
		Name string `json:"name"`
		Sex  string `json:"sex"`
	}{
		{"Luke Skywalker", "male"},
		{"Leia Organa", "female"},
		{"Han Solo", "male"},
		{"Chewbacca", "male"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
