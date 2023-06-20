package user_test

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/porky256/mock-interview-tbr/internal/database"
	"github.com/porky256/mock-interview-tbr/internal/models/repomodels"
	"github.com/porky256/mock-interview-tbr/internal/user"
	"time"
)

var _ = Describe("Postgres Repo", func() {
	var mock sqlmock.Sqlmock
	var db user.DatabaseUserProvider
	var mdb *sql.DB
	BeforeEach(func() {
		var err error
		mdb, mock, err = sqlmock.New()
		db = user.NewPGUserProvider(mdb, 3*time.Second)

		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		mdb.Close()
	})

	Context("User", func() {
		var user repomodels.UserRepo
		BeforeEach(func() {
			user = repomodels.UserRepo{
				ID:          1,
				Username:    "username",
				FirstName:   "first name",
				LastName:    "last name",
				Email:       "a@h.com",
				Password:    "1234",
				Phone:       "4321",
				UserStatus:  1,
				Description: "description",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
		})

		AfterEach(func() {
			Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
		})

		Context("InsertUser", func() {
			exec := "INSERT INTO users"
			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(
					user.Username, user.FirstName, user.LastName, user.Email,
					user.Password, user.Phone, user.UserStatus, user.Description).WillReturnResult(sqlmock.NewResult(1, 1))
				id, err := db.InsertUser(user)
				Expect(err).ToNot(HaveOccurred())
				Expect(id).To(Equal(user.ID))
			})

			It("no rows inserted", func() {
				mock.ExpectExec(exec).WithArgs(
					user.Username, user.FirstName, user.LastName, user.Email,
					user.Password, user.Phone, user.UserStatus, user.Description).WillReturnResult(sqlmock.NewResult(0, 0))
				id, err := db.InsertUser(user)
				Expect(err).To(Equal(database.ErrNoRowsInserted))
				Expect(id).To(Equal(0))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(
					user.Username, user.FirstName, user.LastName, user.Email,
					user.Password, user.Phone, user.UserStatus, user.Description).WillReturnError(fmt.Errorf("error"))
				id, err := db.InsertUser(user)
				Expect(err).To(HaveOccurred())
				Expect(id).To(Equal(0))
			})
		})

		Context("GetUserByID", func() {
			query := `SELECT id, username, first_name, last_name, email, password, phone, 
        user_status, description, created_at, updated_at FROM users`
			It("all good", func() {
				rows := sqlmock.NewRows([]string{"id", "username", "first_name", "last_name", "email", "password",
					"phone", "user_status", "description", "created_at", "updated_at"}).AddRow(
					user.ID, user.Username, user.FirstName, user.LastName, user.Email, user.Password,
					user.Phone, user.UserStatus, user.Description, user.CreatedAt, user.UpdatedAt)
				mock.ExpectQuery(query).WithArgs(user.ID).WillReturnRows(rows)
				got, err := db.GetUserByID(user.ID)
				Expect(err).ToNot(HaveOccurred())
				Expect(*got).To(Equal(user))

			})

			It("handle error", func() {
				mock.ExpectQuery(query).WithArgs(user.ID).WillReturnError(fmt.Errorf("error"))
				got, err := db.GetUserByID(user.ID)
				Expect(err).To(HaveOccurred())
				Expect(got).To(BeNil())
			})
		})

		Context("GetUserByUsername", func() {
			query := `SELECT id, username, first_name, last_name, email, password, phone, 
        user_status, description, created_at, updated_at FROM users`

			It("all good", func() {
				rows := sqlmock.NewRows([]string{"id", "username", "first_name", "last_name", "email", "password",
					"phone", "user_status", "description", "created_at", "updated_at"}).AddRow(
					user.ID, user.Username, user.FirstName, user.LastName, user.Email, user.Password,
					user.Phone, user.UserStatus, user.Description, user.CreatedAt, user.UpdatedAt)
				mock.ExpectQuery(query).WithArgs(user.Username).WillReturnRows(rows)
				got, err := db.GetUserByUsername(user.Username)
				Expect(err).ToNot(HaveOccurred())
				Expect(*got).To(Equal(user))

			})

			It("handle error", func() {
				mock.ExpectQuery(query).WithArgs(user.Username).WillReturnError(fmt.Errorf("error"))
				got, err := db.GetUserByUsername(user.Username)
				Expect(err).To(HaveOccurred())
				Expect(got).To(BeNil())
			})
		})

		Context("UpdateUser", func() {
			var newUser repomodels.UserRepo
			exec := "UPDATE users SET"
			BeforeEach(func() {
				newUser = repomodels.UserRepo{
					ID:          1,
					Username:    "new username",
					FirstName:   "new first name",
					LastName:    "new last name",
					Email:       "new email",
					Password:    "new password",
					Phone:       "new phone",
					UserStatus:  4,
					Description: "new description",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}
			})

			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(
					newUser.Username, newUser.FirstName, newUser.LastName, newUser.Email,
					newUser.Password, newUser.Phone, newUser.UserStatus, newUser.Description, newUser.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				err := db.UpdateUser(newUser)
				Expect(err).ToNot(HaveOccurred())
			})

			It("no rows updated", func() {
				mock.ExpectExec(exec).WithArgs(
					newUser.Username, newUser.FirstName, newUser.LastName, newUser.Email,
					newUser.Password, newUser.Phone, newUser.UserStatus, newUser.Description, newUser.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				err := db.UpdateUser(newUser)
				Expect(err).To(Equal(database.ErrNoRowsUpdated))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(
					newUser.Username, newUser.FirstName, newUser.LastName, newUser.Email,
					newUser.Password, newUser.Phone, newUser.UserStatus, newUser.Description, newUser.ID).WillReturnError(fmt.Errorf("bad error"))
				err := db.UpdateUser(newUser)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("DeleteUserByID", func() {
			exec := `DELETE FROM users`
			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(user.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				err := db.DeleteUserByID(user.ID)
				Expect(err).ToNot(HaveOccurred())

			})

			It("no rows deleted", func() {
				mock.ExpectExec(exec).WithArgs(user.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				err := db.DeleteUserByID(user.ID)
				Expect(err).To(Equal(database.ErrNoRowsDeleted))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(user.ID).WillReturnError(fmt.Errorf("error"))
				err := db.DeleteUserByID(user.ID)
				Expect(err).To(HaveOccurred())
			})
		})

	})

	Context("UserSkill", func() {
		var userSkill repomodels.UserSkillRepo
		BeforeEach(func() {
			userSkill = repomodels.UserSkillRepo{
				ID:        1,
				SkillID:   2,
				UserID:    3,
				Score:     100,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
		})

		AfterEach(func() {
			Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
		})

		Context("InsertUserSkill", func() {
			exec := `INSERT INTO  users_skills`
			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(
					userSkill.SkillID, userSkill.UserID, userSkill.Score).WillReturnResult(sqlmock.NewResult(1, 1))
				id, err := db.InsertUserSkill(userSkill)
				Expect(err).ToNot(HaveOccurred())
				Expect(id).To(Equal(userSkill.ID))
			})

			It("no rows inserted", func() {
				mock.ExpectExec(exec).WithArgs(
					userSkill.SkillID, userSkill.UserID, userSkill.Score).WillReturnResult(sqlmock.NewResult(0, 0))
				id, err := db.InsertUserSkill(userSkill)
				Expect(err).To(Equal(database.ErrNoRowsInserted))
				Expect(id).To(Equal(0))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(
					userSkill.SkillID, userSkill.UserID, userSkill.Score).WillReturnError(fmt.Errorf("error"))
				id, err := db.InsertUserSkill(userSkill)
				Expect(err).To(HaveOccurred())
				Expect(id).To(Equal(0))
			})
		})

		Context("GetUserSkillByID", func() {
			exec := `SELECT id, skill_id, user_id, score, created_at, updated_at  FROM users_skills`

			It("all good", func() {
				rows := sqlmock.NewRows([]string{"id", "skill_id", "user_id", "score", "created_at", "updated_at"}).AddRow(
					userSkill.ID, userSkill.SkillID, userSkill.UserID, userSkill.Score, userSkill.CreatedAt, userSkill.UpdatedAt)
				mock.ExpectQuery(exec).WithArgs(userSkill.ID).WillReturnRows(rows)
				got, err := db.GetUserSkillByID(userSkill.ID)
				Expect(err).ToNot(HaveOccurred())
				Expect(*got).To(Equal(userSkill))

			})

			It("handle error", func() {
				mock.ExpectQuery(exec).WithArgs(userSkill.ID).WillReturnError(fmt.Errorf("error"))
				got, err := db.GetUserSkillByID(userSkill.ID)
				Expect(err).To(HaveOccurred())
				Expect(got).To(BeNil())
			})
		})

		Context("GetUsersSkillsByUserID", func() {
			exec := `SELECT id, skill_id, user_id, score, created_at, updated_at  FROM users_skills`

			It("all good", func() {
				rows := sqlmock.NewRows([]string{"id", "skill_id", "user_id", "score", "created_at", "updated_at"}).AddRow(
					userSkill.ID, userSkill.SkillID, userSkill.UserID, userSkill.Score, userSkill.CreatedAt, userSkill.UpdatedAt).AddRow(
					userSkill.ID, userSkill.SkillID, userSkill.UserID, userSkill.Score, userSkill.CreatedAt, userSkill.UpdatedAt)
				mock.ExpectQuery(exec).WithArgs(userSkill.UserID).WillReturnRows(rows)
				got, err := db.GetUsersSkillsByUserID(userSkill.UserID)
				Expect(err).ToNot(HaveOccurred())
				Expect(len(got)).To(Equal(2))
				Expect(got[0]).To(Equal(userSkill))
				Expect(got[1]).To(Equal(userSkill))
			})

			It("handle error", func() {
				mock.ExpectQuery(exec).WithArgs(userSkill.UserID).WillReturnError(fmt.Errorf("error"))
				got, err := db.GetUsersSkillsByUserID(userSkill.UserID)
				Expect(err).To(HaveOccurred())
				Expect(got).To(BeNil())
			})
		})

		Context("UpdateUserSkill", func() {
			var newUserSkill repomodels.UserSkillRepo
			BeforeEach(func() {
				newUserSkill = repomodels.UserSkillRepo{
					ID:        1,
					SkillID:   4,
					UserID:    5,
					Score:     99,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
			})
			exec := "UPDATE users_skills SET"

			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(
					newUserSkill.SkillID, newUserSkill.UserID, newUserSkill.Score, newUserSkill.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				err := db.UpdateUserSkill(newUserSkill)
				Expect(err).ToNot(HaveOccurred())
			})

			It("no rows updated", func() {
				mock.ExpectExec(exec).WithArgs(
					newUserSkill.SkillID, newUserSkill.UserID, newUserSkill.Score, newUserSkill.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				err := db.UpdateUserSkill(newUserSkill)
				Expect(err).To(Equal(database.ErrNoRowsUpdated))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(
					newUserSkill.SkillID, newUserSkill.UserID, newUserSkill.Score, newUserSkill.ID).WillReturnError(fmt.Errorf("bad error"))
				err := db.UpdateUserSkill(newUserSkill)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("DeleteUserSkillByID", func() {
			exec := `DELETE FROM users_skills`
			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(userSkill.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				err := db.DeleteUserSkillByID(userSkill.ID)
				Expect(err).ToNot(HaveOccurred())

			})
			It("no rows deleted", func() {
				mock.ExpectExec(exec).WithArgs(userSkill.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				err := db.DeleteUserSkillByID(userSkill.ID)
				Expect(err).To(Equal(database.ErrNoRowsDeleted))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(userSkill.ID).WillReturnError(fmt.Errorf("error"))
				err := db.DeleteUserSkillByID(userSkill.ID)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("DeleteUserSkillByUserID", func() {
			exec := `DELETE FROM users_skills`

			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(userSkill.UserID).WillReturnResult(sqlmock.NewResult(1, 1))
				err := db.DeleteUserSkillByUserID(userSkill.UserID)
				Expect(err).ToNot(HaveOccurred())

			})

			It("no rows deleted", func() {
				mock.ExpectExec(exec).WithArgs(userSkill.UserID).WillReturnResult(sqlmock.NewResult(0, 0))
				err := db.DeleteUserSkillByUserID(userSkill.UserID)
				Expect(err).To(Equal(database.ErrNoRowsDeleted))

			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(userSkill.UserID).WillReturnError(fmt.Errorf("error"))
				err := db.DeleteUserSkillByUserID(userSkill.UserID)
				Expect(err).To(HaveOccurred())
			})
		})
	})

})
