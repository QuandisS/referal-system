package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/markbates/goth/gothic"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))

	r.Get("/", s.HandleIndex)
	r.Get("/dbhealth", s.HandleDbHealth)
	r.Get("/auth/{provider}", s.HandleProviderLogin)
	r.Get("/auth/{provider}/callback", s.HandleProviderCallbackFunction)
	return r
}

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Printf("error handling JSON marshal. Err: %v", err)
		return
	}

	_, err = w.Write(jsonResp)
	if err != nil {
		log.Printf("error writing response. Err: %v", err)
		return
	}
}

func (s *Server) HandleDbHealth(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, err := w.Write(jsonResp)
	if err != nil {
		log.Printf("error writing response. Err: %v", err)
		return
	}
}

func (s *Server) HandleProviderLogin(w http.ResponseWriter, r *http.Request) {
	if u, err := gothic.CompleteUserAuth(w, r); err == nil {
		log.Printf("User %s already authenticated!", u)
		return
	}

	gothic.BeginAuthHandler(w, r)
}

func (s *Server) HandleProviderCallbackFunction(w http.ResponseWriter, r *http.Request) {
	u, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Printf("error completing user authentication. Err: %v", err)
		return
	}
	
	log.Printf("User %s authenticated!", u)

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}