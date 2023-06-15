package routes

import (
	"fmt"
	"log"

	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/porky256/mock-interview-tbr/internal/handlers"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", Index)

	router.Get("/matching/match/{matchID}", handlers.GetMatch)
	router.Get("/matching/match/findByUserID", handlers.GetMatchByUserID)
	router.Get("/matching/interview/{interviewID}", handlers.GetInterviewByID)
	router.Post("/matching/match/{matchID}", handlers.PostMatch)
	router.Post("/matching/request", handlers.PostMatchingRequest)

	router.Post("/skill", handlers.CreateSkill)
	router.Get("/skill/{skillID}", handlers.GetSkill)
	router.Put("/skill/{skillID}", handlers.UpdateSkill)
	router.Delete("/skill/{skillID}", handlers.DeleteSkill)

	router.Post("/user", handlers.CreateUser)
	router.Get("/user/login", handlers.LoginUser)
	router.Get("/user/logout", handlers.LogoutUser)
	router.Post("/user/{username}/addSkill", handlers.PostAddSkillToUser)
	router.Get("/user/{username}", handlers.GetUserByName)
	router.Put("/user/{username}", handlers.UpdateUser)
	router.Delete("/user/{username}", handlers.DeleteUser)

	return router
}

func Index(w http.ResponseWriter, _ *http.Request) {
	log.Println(fmt.Fprintf(w, "Hello World!"))
}
