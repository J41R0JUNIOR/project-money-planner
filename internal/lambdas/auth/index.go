package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Method(http.MethodPost, "/signup", SignUpHandler)
}
