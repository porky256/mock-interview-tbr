package matchrepo

import "github.com/porky256/mock-interview-tbr/internal/models/repomodels"

type userMatchProvider interface {
	InsertUserMatch(match repomodels.UserMatchRepo) (int, error)
	GetUserMatchByID(id int) (*repomodels.UserMatchRepo, error)
	GetUserMatchByUserAskerID(userID int) (*repomodels.UserMatchRepo, error)
	UpdateUserMatch(match repomodels.UserMatchRepo) error
	DeleteUserMatchByID(id int) error
}

type matchRequestProvider interface {
	InsertMatchRequest(req repomodels.MatchRequestRepo) (int, error)
	GetMatchRequestByID(id int) (*repomodels.MatchRequestRepo, error)
	GetMatchRequestByUserID(userID int) (*repomodels.MatchRequestRepo, error)
	UpdateMatchRequest(req repomodels.MatchRequestRepo) error
	DeleteMatchRequestByID(id int) error
}

type skillInRequestProvider interface {
	InsertSkillInRequest(skill repomodels.SkillInRequestRepo) (int, error)
	GetSkillInRequestByID(id int) (*repomodels.SkillInRequestRepo, error)
	GetSkillInRequestByRequestID(requestID int) (*repomodels.SkillInRequestRepo, error)
	DeleteSkillInRequestByID(id int) error
	DeleteSkillInRequestByRequestID(requestID int) error
}

type interviewsProvider interface {
	InsertInterview(interview repomodels.InterviewRepo) (int, error)
	GetInterviewByID(id int) (*repomodels.InterviewRepo, error)
	UpdateInterview(interview repomodels.InterviewRepo) error
	DeleteInterviewByID(id int) error
}

// DatabaseMatchProvider provider for db related functions for users
type DatabaseMatchProvider interface {
	userMatchProvider
	matchRequestProvider
	skillInRequestProvider
	interviewsProvider
}
