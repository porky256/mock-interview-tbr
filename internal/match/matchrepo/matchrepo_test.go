package matchrepo_test

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/porky256/mock-interview-tbr/internal/database"
	"github.com/porky256/mock-interview-tbr/internal/match/matchrepo"
	"github.com/porky256/mock-interview-tbr/internal/models/repomodels"
	"time"
)

var _ = Describe("Postgres Repo", func() {
	var mock sqlmock.Sqlmock
	var db matchrepo.DatabaseMatchProvider
	var mdb *sql.DB
	BeforeEach(func() {
		var err error
		mdb, mock, err = sqlmock.New()
		db = matchrepo.NewPGMatchProvider(mdb, 3*time.Second)

		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		mdb.Close()
	})

	Context("UserMatch", func() {
		var match repomodels.UserMatchRepo
		BeforeEach(func() {
			match = repomodels.UserMatchRepo{
				ID:         1,
				UserAsker:  2,
				UserMatch:  3,
				MatchScore: 100,
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			}
		})

		AfterEach(func() {
			Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
		})

		Context("InsertUserMatch", func() {
			exec := "INSERT INTO users_matches"
			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(
					match.UserAsker, match.UserMatch, match.MatchScore).WillReturnResult(sqlmock.NewResult(1, 1))
				id, err := db.InsertUserMatch(match)
				Expect(err).ToNot(HaveOccurred())
				Expect(id).To(Equal(match.ID))
			})

			It("no rows inserted", func() {
				mock.ExpectExec(exec).WithArgs(
					match.UserAsker, match.UserMatch, match.MatchScore).WillReturnResult(sqlmock.NewResult(0, 0))
				id, err := db.InsertUserMatch(match)
				Expect(err).To(Equal(database.ErrNoRowsInserted))
				Expect(id).To(Equal(0))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(
					match.UserAsker, match.UserMatch, match.MatchScore).WillReturnError(fmt.Errorf("error"))
				id, err := db.InsertUserMatch(match)
				Expect(err).To(HaveOccurred())
				Expect(id).To(Equal(0))
			})
		})

		Context("GetUserMatchByID", func() {
			query := `SELECT id, user_asker, user_match, match_score, created_at, updated_at 
				FROM users_matches`
			It("all good", func() {
				rows := sqlmock.NewRows([]string{"id", "user_asker", "user_match",
					"match_score", "created_at", "updated_at"}).AddRow(
					match.ID, match.UserAsker, match.UserMatch, match.MatchScore, match.CreatedAt, match.UpdatedAt)
				mock.ExpectQuery(query).WithArgs(match.ID).WillReturnRows(rows)
				got, err := db.GetUserMatchByID(match.ID)
				Expect(err).ToNot(HaveOccurred())
				Expect(*got).To(Equal(match))

			})

			It("handle error", func() {
				mock.ExpectQuery(query).WithArgs(match.ID).WillReturnError(fmt.Errorf("error"))
				got, err := db.GetUserMatchByID(match.ID)
				Expect(err).To(HaveOccurred())
				Expect(got).To(BeNil())
			})
		})

		Context("GetUserMatchByUserAskerID", func() {
			query := `SELECT id, user_asker, user_match, match_score, created_at, updated_at 
				FROM users_matches`

			It("all good", func() {
				rows := sqlmock.NewRows([]string{"id", "user_asker", "user_match",
					"match_score", "created_at", "updated_at"}).AddRow(
					match.ID, match.UserAsker, match.UserMatch, match.MatchScore, match.CreatedAt, match.UpdatedAt)
				mock.ExpectQuery(query).WithArgs(match.UserAsker).WillReturnRows(rows)
				got, err := db.GetUserMatchByUserAskerID(match.UserAsker)
				Expect(err).ToNot(HaveOccurred())
				Expect(*got).To(Equal(match))

			})

			It("handle error", func() {
				mock.ExpectQuery(query).WithArgs(match.UserAsker).WillReturnError(fmt.Errorf("error"))
				got, err := db.GetUserMatchByUserAskerID(match.UserAsker)
				Expect(err).To(HaveOccurred())
				Expect(got).To(BeNil())
			})
		})

		Context("UpdateUserMatch", func() {
			var newUserMatch repomodels.UserMatchRepo
			exec := "UPDATE users_matches SET"
			BeforeEach(func() {
				newUserMatch = repomodels.UserMatchRepo{
					ID:         1,
					UserAsker:  6,
					UserMatch:  7,
					MatchScore: 8,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}
			})

			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(
					newUserMatch.UserAsker, newUserMatch.UserMatch,
					newUserMatch.MatchScore, newUserMatch.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				err := db.UpdateUserMatch(newUserMatch)
				Expect(err).ToNot(HaveOccurred())
			})

			It("no rows updated", func() {
				mock.ExpectExec(exec).WithArgs(
					newUserMatch.UserAsker, newUserMatch.UserMatch,
					newUserMatch.MatchScore, newUserMatch.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				err := db.UpdateUserMatch(newUserMatch)
				Expect(err).To(Equal(database.ErrNoRowsUpdated))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(
					newUserMatch.UserAsker, newUserMatch.UserMatch, newUserMatch.MatchScore, newUserMatch.ID).WillReturnError(fmt.Errorf("bad error"))
				err := db.UpdateUserMatch(newUserMatch)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("DeleteUserMatchByID", func() {
			exec := `DELETE FROM users_matches`
			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(match.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				err := db.DeleteUserMatchByID(match.ID)
				Expect(err).ToNot(HaveOccurred())

			})

			It("no rows deleted", func() {
				mock.ExpectExec(exec).WithArgs(match.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				err := db.DeleteUserMatchByID(match.ID)
				Expect(err).To(Equal(database.ErrNoRowsDeleted))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(match.ID).WillReturnError(fmt.Errorf("error"))
				err := db.DeleteUserMatchByID(match.ID)
				Expect(err).To(HaveOccurred())
			})
		})

	})

	Context("MatchRequest", func() {
		var req repomodels.MatchRequestRepo
		BeforeEach(func() {
			req = repomodels.MatchRequestRepo{
				ID:        1,
				UserID:    2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
		})

		AfterEach(func() {
			Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
		})

		Context("InsertMatchRequest", func() {
			exec := "INSERT INTO match_requests"
			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(
					req.UserID).WillReturnResult(sqlmock.NewResult(1, 1))
				id, err := db.InsertMatchRequest(req)
				Expect(err).ToNot(HaveOccurred())
				Expect(id).To(Equal(req.ID))
			})

			It("no rows inserted", func() {
				mock.ExpectExec(exec).WithArgs(
					req.UserID).WillReturnResult(sqlmock.NewResult(0, 0))
				id, err := db.InsertMatchRequest(req)
				Expect(err).To(Equal(database.ErrNoRowsInserted))
				Expect(id).To(Equal(0))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(
					req.UserID).WillReturnError(fmt.Errorf("error"))
				id, err := db.InsertMatchRequest(req)
				Expect(err).To(HaveOccurred())
				Expect(id).To(Equal(0))
			})
		})

		Context("GetMatchRequestByID", func() {
			query := `SELECT id, user_id, created_at, updated_at FROM match_requests`
			It("all good", func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "created_at", "updated_at"}).AddRow(
					req.ID, req.UserID, req.CreatedAt, req.UpdatedAt)
				mock.ExpectQuery(query).WithArgs(req.ID).WillReturnRows(rows)
				got, err := db.GetMatchRequestByID(req.ID)
				Expect(err).ToNot(HaveOccurred())
				Expect(*got).To(Equal(req))

			})

			It("handle error", func() {
				mock.ExpectQuery(query).WithArgs(req.ID).WillReturnError(fmt.Errorf("error"))
				got, err := db.GetMatchRequestByID(req.ID)
				Expect(err).To(HaveOccurred())
				Expect(got).To(BeNil())
			})
		})

		Context("GetMatchRequestByUserID", func() {
			query := `SELECT id, user_id, created_at, updated_at FROM match_requests`

			It("all good", func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "created_at", "updated_at"}).AddRow(
					req.ID, req.UserID, req.CreatedAt, req.UpdatedAt)
				mock.ExpectQuery(query).WithArgs(req.UserID).WillReturnRows(rows)
				got, err := db.GetMatchRequestByUserID(req.UserID)
				Expect(err).ToNot(HaveOccurred())
				Expect(*got).To(Equal(req))

			})

			It("handle error", func() {
				mock.ExpectQuery(query).WithArgs(req.UserID).WillReturnError(fmt.Errorf("error"))
				got, err := db.GetMatchRequestByUserID(req.UserID)
				Expect(err).To(HaveOccurred())
				Expect(got).To(BeNil())
			})
		})

		Context("UpdateMatchRequest", func() {
			var newMatchRequest repomodels.MatchRequestRepo
			exec := "UPDATE match_requests SET"
			BeforeEach(func() {
				newMatchRequest = repomodels.MatchRequestRepo{
					ID:        1,
					UserID:    6,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
			})

			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(
					newMatchRequest.UserID, newMatchRequest.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				err := db.UpdateMatchRequest(newMatchRequest)
				Expect(err).ToNot(HaveOccurred())
			})

			It("no rows updated", func() {
				mock.ExpectExec(exec).WithArgs(
					newMatchRequest.UserID, newMatchRequest.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				err := db.UpdateMatchRequest(newMatchRequest)
				Expect(err).To(Equal(database.ErrNoRowsUpdated))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(
					newMatchRequest.UserID, newMatchRequest.ID).WillReturnError(fmt.Errorf("bad error"))
				err := db.UpdateMatchRequest(newMatchRequest)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("DeleteMatchRequestByID", func() {
			exec := `DELETE FROM match_requests`
			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(req.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				err := db.DeleteMatchRequestByID(req.ID)
				Expect(err).ToNot(HaveOccurred())

			})

			It("no rows deleted", func() {
				mock.ExpectExec(exec).WithArgs(req.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				err := db.DeleteMatchRequestByID(req.ID)
				Expect(err).To(Equal(database.ErrNoRowsDeleted))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(req.ID).WillReturnError(fmt.Errorf("error"))
				err := db.DeleteMatchRequestByID(req.ID)
				Expect(err).To(HaveOccurred())
			})
		})

	})

	Context("SkillInRequest", func() {
		var skill repomodels.SkillInRequestRepo
		BeforeEach(func() {
			skill = repomodels.SkillInRequestRepo{
				ID:        1,
				RequestID: 2,
				SkillID:   3,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
		})

		AfterEach(func() {
			Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
		})

		Context("InsertSkillInRequest", func() {
			exec := "INSERT INTO skills_in_requests"
			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(
					skill.RequestID, skill.SkillID).WillReturnResult(sqlmock.NewResult(1, 1))
				id, err := db.InsertSkillInRequest(skill)
				Expect(err).ToNot(HaveOccurred())
				Expect(id).To(Equal(skill.ID))
			})

			It("no rows inserted", func() {
				mock.ExpectExec(exec).WithArgs(
					skill.RequestID, skill.SkillID).WillReturnResult(sqlmock.NewResult(0, 0))
				id, err := db.InsertSkillInRequest(skill)
				Expect(err).To(Equal(database.ErrNoRowsInserted))
				Expect(id).To(Equal(0))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(
					skill.RequestID, skill.SkillID).WillReturnError(fmt.Errorf("error"))
				id, err := db.InsertSkillInRequest(skill)
				Expect(err).To(HaveOccurred())
				Expect(id).To(Equal(0))
			})
		})

		Context("GetSkillInRequestByID", func() {
			query := `SELECT id, request_id, skill_id, created_at, updated_at 
				FROM skills_in_requests`
			It("all good", func() {
				rows := sqlmock.NewRows([]string{"id", "request_id", "skill_id", "created_at", "updated_at"}).AddRow(
					skill.ID, skill.RequestID, skill.SkillID, skill.CreatedAt, skill.UpdatedAt)
				mock.ExpectQuery(query).WithArgs(skill.ID).WillReturnRows(rows)
				got, err := db.GetSkillInRequestByID(skill.ID)
				Expect(err).ToNot(HaveOccurred())
				Expect(*got).To(Equal(skill))

			})

			It("handle error", func() {
				mock.ExpectQuery(query).WithArgs(skill.ID).WillReturnError(fmt.Errorf("error"))
				got, err := db.GetSkillInRequestByID(skill.ID)
				Expect(err).To(HaveOccurred())
				Expect(got).To(BeNil())
			})
		})

		Context("GetSkillInRequestByUserAskerID", func() {
			query := `SELECT id, request_id, skill_id, created_at, updated_at 
				FROM skills_in_requests`

			It("all good", func() {
				rows := sqlmock.NewRows([]string{"id", "request_id", "skill_id", "created_at", "updated_at"}).AddRow(
					skill.ID, skill.RequestID, skill.SkillID, skill.CreatedAt, skill.UpdatedAt)
				mock.ExpectQuery(query).WithArgs(skill.RequestID).WillReturnRows(rows)
				got, err := db.GetSkillInRequestByRequestID(skill.RequestID)
				Expect(err).ToNot(HaveOccurred())
				Expect(*got).To(Equal(skill))

			})

			It("handle error", func() {
				mock.ExpectQuery(query).WithArgs(skill.RequestID).WillReturnError(fmt.Errorf("error"))
				got, err := db.GetSkillInRequestByRequestID(skill.RequestID)
				Expect(err).To(HaveOccurred())
				Expect(got).To(BeNil())
			})
		})

		Context("DeleteSkillInRequestByID", func() {
			exec := `DELETE FROM skills_in_requests`
			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(skill.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				err := db.DeleteSkillInRequestByID(skill.ID)
				Expect(err).ToNot(HaveOccurred())

			})

			It("no rows deleted", func() {
				mock.ExpectExec(exec).WithArgs(skill.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				err := db.DeleteSkillInRequestByID(skill.ID)
				Expect(err).To(Equal(database.ErrNoRowsDeleted))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(skill.ID).WillReturnError(fmt.Errorf("error"))
				err := db.DeleteSkillInRequestByID(skill.ID)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("DeleteSkillInRequestByRequestID", func() {
			exec := `DELETE FROM skills_in_requests`
			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(skill.RequestID).WillReturnResult(sqlmock.NewResult(1, 1))
				err := db.DeleteSkillInRequestByRequestID(skill.RequestID)
				Expect(err).ToNot(HaveOccurred())

			})

			It("no rows deleted", func() {
				mock.ExpectExec(exec).WithArgs(skill.RequestID).WillReturnResult(sqlmock.NewResult(0, 0))
				err := db.DeleteSkillInRequestByRequestID(skill.RequestID)
				Expect(err).To(Equal(database.ErrNoRowsDeleted))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(skill.RequestID).WillReturnError(fmt.Errorf("error"))
				err := db.DeleteSkillInRequestByRequestID(skill.RequestID)
				Expect(err).To(HaveOccurred())
			})
		})

	})

	Context("Interview", func() {
		var interview repomodels.InterviewRepo
		BeforeEach(func() {
			interview = repomodels.InterviewRepo{
				ID:            1,
				MatchID:       2,
				Status:        3,
				InterviewDate: time.Now().Add(time.Hour * 3 * 24),
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			}
		})

		AfterEach(func() {
			Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
		})

		Context("InsertInterview", func() {
			exec := "INSERT INTO interviews"
			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(
					interview.MatchID, interview.Status, interview.InterviewDate).WillReturnResult(sqlmock.NewResult(1, 1))
				id, err := db.InsertInterview(interview)
				Expect(err).ToNot(HaveOccurred())
				Expect(id).To(Equal(interview.ID))
			})

			It("no rows inserted", func() {
				mock.ExpectExec(exec).WithArgs(
					interview.MatchID, interview.Status, interview.InterviewDate).WillReturnResult(sqlmock.NewResult(0, 0))
				id, err := db.InsertInterview(interview)
				Expect(err).To(Equal(database.ErrNoRowsInserted))
				Expect(id).To(Equal(0))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(
					interview.MatchID, interview.Status, interview.InterviewDate).WillReturnError(fmt.Errorf("error"))
				id, err := db.InsertInterview(interview)
				Expect(err).To(HaveOccurred())
				Expect(id).To(Equal(0))
			})
		})

		Context("GetInterviewByID", func() {
			query := `SELECT id, match_id, status, interview_date, created_at, updated_at 
				FROM interviews`
			It("all good", func() {
				rows := sqlmock.NewRows([]string{"id", "match_id", "status",
					"interview_date", "created_at", "updated_at"}).AddRow(
					interview.ID, interview.MatchID, interview.Status, interview.InterviewDate, interview.CreatedAt, interview.UpdatedAt)
				mock.ExpectQuery(query).WithArgs(interview.ID).WillReturnRows(rows)
				got, err := db.GetInterviewByID(interview.ID)
				Expect(err).ToNot(HaveOccurred())
				Expect(*got).To(Equal(interview))

			})

			It("handle error", func() {
				mock.ExpectQuery(query).WithArgs(interview.ID).WillReturnError(fmt.Errorf("error"))
				got, err := db.GetInterviewByID(interview.ID)
				Expect(err).To(HaveOccurred())
				Expect(got).To(BeNil())
			})
		})

		Context("UpdateInterview", func() {
			var newInterview repomodels.InterviewRepo
			exec := "UPDATE interviews SET"
			BeforeEach(func() {
				newInterview = repomodels.InterviewRepo{
					ID:            1,
					MatchID:       4,
					Status:        5,
					InterviewDate: time.Now().Add(time.Hour * 4 * 24),
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
				}
			})

			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(
					newInterview.MatchID, newInterview.Status,
					newInterview.InterviewDate, newInterview.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				err := db.UpdateInterview(newInterview)
				Expect(err).ToNot(HaveOccurred())
			})

			It("no rows updated", func() {
				mock.ExpectExec(exec).WithArgs(
					newInterview.MatchID, newInterview.Status,
					newInterview.InterviewDate, newInterview.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				err := db.UpdateInterview(newInterview)
				Expect(err).To(Equal(database.ErrNoRowsUpdated))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(
					newInterview.MatchID, newInterview.Status,
					newInterview.InterviewDate, newInterview.ID).WillReturnError(fmt.Errorf("bad error"))
				err := db.UpdateInterview(newInterview)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("DeleteInterviewByID", func() {
			exec := `DELETE FROM interviews`
			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(interview.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				err := db.DeleteInterviewByID(interview.ID)
				Expect(err).ToNot(HaveOccurred())

			})

			It("no rows deleted", func() {
				mock.ExpectExec(exec).WithArgs(interview.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				err := db.DeleteInterviewByID(interview.ID)
				Expect(err).To(Equal(database.ErrNoRowsDeleted))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(interview.ID).WillReturnError(fmt.Errorf("error"))
				err := db.DeleteInterviewByID(interview.ID)
				Expect(err).To(HaveOccurred())
			})
		})

	})

})
