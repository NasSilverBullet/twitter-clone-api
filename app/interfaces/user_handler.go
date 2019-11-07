package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/NasSilverBullet/twitter-clone-api/app/usecases"
)

type UserHandler struct {
	UserInteractor *usecases.UserInteractor
}

func NewUserHandler(sqlHandler SQLHandler) *UserHandler {
	return &UserHandler{
		UserInteractor: &usecases.UserInteractor{
			UserRepository: &UserRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (uh *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	us, err := uh.UserInteractor.Index()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(us)
}
