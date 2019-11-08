package frameworks

import (
	"fmt"
	"net/http"
	"os"

	"github.com/NasSilverBullet/twitter-clone-api/app/interfaces"
	"github.com/NasSilverBullet/twitter-clone-api/app/usecases"
	"github.com/go-chi/chi"
)

func Routes(logger usecases.Logger, sqlHandler interfaces.SQLHandler) error {
	logger.Info("Start running router")

	r := chi.NewRouter()

	uh := interfaces.NewUserHandler(logger, sqlHandler)

	r.Get("/users", uh.Index)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), r); err != nil {
		return err
	}

	return nil
}
