package frameworks

import (
	"net/http"

	"github.com/NasSilverBullet/twitter-clone-api/app/interfaces"
	"github.com/go-chi/chi"
)

func Routes(sqlHandler interfaces.SQLHandler) error {

	r := chi.NewRouter()

	uh := interfaces.NewUserHandler(sqlHandler)

	r.Get("/users", uh.Index)

	if err := http.ListenAndServe(":8080", r); err != nil {
		return err
	}

	return nil
}
