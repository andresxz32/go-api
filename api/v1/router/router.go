package router

import (
	"fmt"
	"go-api/api/v1/handlers"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func Initialize() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), //forces Content-type
		middleware.RedirectSlashes,
		middleware.Recoverer,            //middleware to recover from panics
		middleware.Heartbeat("/health"), //for heartbeat process such as Kubernetes liveprobeness
		cors.Handler(cors.Options{
			// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"https://*", "http://*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}),
	)

	//Sets context for all requests
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(30 * time.Second))

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/", handlers.Routes()) //Implementation of routes from handlers.go
		//r.Mount("/metrics", nil)        //for monitoring agents such as prometheus
	})

	return router
}

func ServeRouter() {
	router := Initialize()

	fmt.Println("Server Running: http://localhost:8000/api/v1")
	err := http.ListenAndServe(":8000", router)
	log.Println(err)
	if err != nil {
		log.Fatal("Error serving router")
	}
}
