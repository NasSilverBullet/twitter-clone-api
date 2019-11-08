package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/NasSilverBullet/twitter-clone-api/app/usecases"
)

type UserHandler struct {
	UserInteractor *usecases.UserInteractor
	Logger         usecases.Logger
}

func NewUserHandler(logger usecases.Logger, sqlHandler SQLHandler) *UserHandler {
	return &UserHandler{
		UserInteractor: &usecases.UserInteractor{
			UserRepository: &UserRepository{
				SQLHandler: sqlHandler,
			},
		},
		Logger: logger,
	}
}

func (uh *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	uh.Logger.Infof("%s: %s => %s", r.Method, r.URL.Path, "Start user handler's List")

	us, err := uh.UserInteractor.List()
	if err != nil {
		uh.Logger.Error(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}

	uh.Logger.Infof("%s: %s => Success find %d users", r.Method, r.URL.Path, len(us))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(us)
}
