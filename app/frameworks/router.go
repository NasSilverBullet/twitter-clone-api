package frameworks

import (
	"fmt"
	"net/http"
	"os"

	"github.com/NasSilverBullet/twitter-clone-api/app/interfaces"
	"github.com/NasSilverBullet/twitter-clone-api/app/usecases"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(logger usecases.Logger, sqlHandler interfaces.SQLHandler) error {
	logger.Info("Start running router..")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)

	uh := interfaces.NewUserHandler(logger, sqlHandler)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", uh.List)
		r.Get("/{id}", uh.Get)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), r); err != nil {
		return err
	}

	return nil
}
