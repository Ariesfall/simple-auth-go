package serv

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
)

// generate and set jwt cookie
func newJwtCookie(w http.ResponseWriter, r *http.Request) {
	id := rand.Intn(100)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{
		"user_id": id,
		"iat":     jwtauth.EpochNow(),
		"exp":     jwtauth.ExpireIn(3600 * time.Second),
	})

	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: tokenString})
	w.Write([]byte("You got a new jwt Cookie, it will expire in 3600s"))
}
