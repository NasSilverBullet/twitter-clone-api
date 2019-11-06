package frameworks

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func Routes() {
	r := chi.NewRouter()
	r.Get("/", handler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
