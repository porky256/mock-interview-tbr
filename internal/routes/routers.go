package routes

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/porky256/mock-interview-tbr/internal/match"
	"github.com/porky256/mock-interview-tbr/internal/skill"
	"github.com/porky256/mock-interview-tbr/internal/user"
	"net/http"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", Index)

	router.Get("/matching/match/{matchID}", match.GetMatch)
	router.Get("/matching/match/findByUserID", match.GetMatchByUserID)
	router.Get("/matching/interview/{interviewID}", match.GetInterviewByID)
	router.Post("/matching/match/{matchID}", match.PostMatch)
	router.Post("/matching/request", match.PostMatchingRequest)

	router.Post("/skill", skill.CreateSkill)
	router.Get("/skill/{skillID}", skill.GetSkill)
	router.Put("/skill/{skillID}", skill.UpdateSkill)
	router.Delete("/skill/{skillID}", skill.DeleteSkill)

	router.Post("/user", user.CreateUser)
	router.Get("/user/login", user.LoginUser)
	router.Get("/user/logout", user.LogoutUser)
	router.Post("/user/{username}/addSkill", user.PostAddSkillToUser)
	router.Get("/user/{username}", user.GetUserByName)
	router.Put("/user/{username}", user.UpdateUser)
	router.Delete("/user/{username}", user.DeleteUser)

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
