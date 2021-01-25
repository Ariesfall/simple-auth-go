package serv

import (
	"log"
	"net/http"
	"time"

	"github.com/Ariesfall/simple-auth-go/pkg/conn"
	"github.com/Ariesfall/simple-auth-go/pkg/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/jmoiron/sqlx"
)

const (
	jwtKey = "1e3dqwddf21dwd1"
)

var (
	db        *sqlx.DB
	tokenAuth *jwtauth.JWTAuth
)

// Run start init and api service
func Run() {
	inits()
	service()
}

func inits() {
	// init db connection
	log.Println("connecting to db")
	db = conn.Connectdb(conn.SampleConn)

	// init jwt token
	tokenAuth = jwtauth.New("HS256", []byte(jwtKey), nil)
}

func service() {
	r := chi.NewRouter()

	// request real ip and log
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	// request timeout
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", newJwtCookie)

	r.Group(func(r chi.Router) {
		// verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		// api path
		r.Get("/task", handler.TaskGet(db))
		r.Post("/task", handler.TaskCreate(db))
		r.Put("/task", handler.TaskUpdate(db))
		r.Post("/task", handler.TaskDelete(db))
	})

	// port listen
	log.Println("start service on port :8080")
	http.ListenAndServe(":8080", r)
}
