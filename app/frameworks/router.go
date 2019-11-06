package frameworks

import (
	"fmt"
	"net/http"
	"os"

	"github.com/NasSilverBullet/twitter-clone-api/app/interfaces"
	"github.com/go-chi/chi"
)

func Routes() {

	r := chi.NewRouter()

	uh := interfaces.NewUserHandler()

	r.Get("/users", uh.Index)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
