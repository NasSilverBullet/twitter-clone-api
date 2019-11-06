package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/NasSilverBullet/twitter-clone-api/app/usecases"
)

type UsersHandler struct {
	UserInteractor *usecases.UsersInteractor
}

func NewUsersHandler() *UsersHandler {
	return &UsersHandler{
		UserInteractor: &usecases.UsersInteractor{},
	}
}

func (uh *UsersHandler) Index(w http.ResponseWriter, r *http.Request) {
	u, err := uh.UserInteractor.Index()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}
