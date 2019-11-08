package interfaces

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"

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

func (uh *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	uh.Logger.Infof("%s: %s => %s", r.Method, r.URL.Path, "Start user handler's Get")

	_, idstr := path.Split(r.URL.Path)

	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		uh.Logger.Error(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)

		return
	}

	uh.Logger.Infof("%s: %s => Retrieved url param id %d", r.Method, r.URL.Path, id)

	u, err := uh.UserInteractor.Get(id)
	if err != nil {
		uh.Logger.Error(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)

		return
	}

	uh.Logger.Infof("%s: %s => Success get user", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}
