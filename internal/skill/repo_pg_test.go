package skill_test

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/porky256/mock-interview-tbr/internal/database"
	"github.com/porky256/mock-interview-tbr/internal/models/repomodels"
	"github.com/porky256/mock-interview-tbr/internal/skill"
	"time"
)

var _ = Describe("Postgres Repo", func() {
	var mock sqlmock.Sqlmock
	var db skill.DatabaseSkillProvider
	var mdb *sql.DB
	BeforeEach(func() {
		var err error
		mdb, mock, err = sqlmock.New()
		db = skill.NewPGSkillProvider(mdb, 3*time.Second)

		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		mdb.Close()
	})

	Context("Skill", func() {
		var skill repomodels.SkillRepo
		BeforeEach(func() {
			skill = repomodels.SkillRepo{
				ID:          1,
				Name:        "name",
				Description: "description",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
		})

		AfterEach(func() {
			Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
		})

		Context("InsertSkill", func() {
			exec := "INSERT INTO  skills"
			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(
					skill.Name, skill.Description).WillReturnResult(sqlmock.NewResult(1, 1))
				id, err := db.InsertSkill(skill)
				Expect(err).ToNot(HaveOccurred())
				Expect(id).To(Equal(skill.ID))
			})

			It("no rows inserted", func() {
				mock.ExpectExec(exec).WithArgs(
					skill.Name, skill.Description).WillReturnResult(sqlmock.NewResult(0, 0))
				id, err := db.InsertSkill(skill)
				Expect(err).To(Equal(database.ErrNoRowsInserted))
				Expect(id).To(Equal(0))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(
					skill.Name, skill.Description).WillReturnError(fmt.Errorf("error"))
				id, err := db.InsertSkill(skill)
				Expect(err).To(HaveOccurred())
				Expect(id).To(Equal(0))
			})
		})

		Context("GetSkillByID", func() {
			query := `SELECT id, name, description, created_at, updated_at FROM skills`
			It("all good", func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"}).AddRow(
					skill.ID, skill.Name, skill.Description, skill.CreatedAt, skill.UpdatedAt)
				mock.ExpectQuery(query).WithArgs(skill.ID).WillReturnRows(rows)
				got, err := db.GetSkillByID(skill.ID)
				Expect(err).ToNot(HaveOccurred())
				Expect(*got).To(Equal(skill))

			})

			It("handle error", func() {
				mock.ExpectQuery(query).WithArgs(skill.ID).WillReturnError(fmt.Errorf("error"))
				got, err := db.GetSkillByID(skill.ID)
				Expect(err).To(HaveOccurred())
				Expect(got).To(BeNil())
			})
		})

		Context("GetSkillByName", func() {
			query := `SELECT id, name, description, created_at, updated_at FROM skills`

			It("all good", func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"}).AddRow(
					skill.ID, skill.Name, skill.Description, skill.CreatedAt, skill.UpdatedAt)
				mock.ExpectQuery(query).WithArgs(skill.Name).WillReturnRows(rows)
				got, err := db.GetSkillByName(skill.Name)
				Expect(err).ToNot(HaveOccurred())
				Expect(*got).To(Equal(skill))

			})

			It("handle error", func() {
				mock.ExpectQuery(query).WithArgs(skill.Name).WillReturnError(fmt.Errorf("error"))
				got, err := db.GetSkillByName(skill.Name)
				Expect(err).To(HaveOccurred())
				Expect(got).To(BeNil())
			})
		})

		Context("UpdateSkill", func() {
			var newSkill repomodels.SkillRepo
			exec := "UPDATE skills SET"
			BeforeEach(func() {
				newSkill = repomodels.SkillRepo{
					ID:          1,
					Name:        "new name",
					Description: "new description",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}
			})

			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(
					newSkill.Name, newSkill.Description, newSkill.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				err := db.UpdateSkill(newSkill)
				Expect(err).ToNot(HaveOccurred())
			})

			It("no rows updated", func() {
				mock.ExpectExec(exec).WithArgs(
					newSkill.Name, newSkill.Description, newSkill.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				err := db.UpdateSkill(newSkill)
				Expect(err).To(Equal(database.ErrNoRowsUpdated))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(
					newSkill.Name, newSkill.Description, newSkill.ID).WillReturnError(fmt.Errorf("bad error"))
				err := db.UpdateSkill(newSkill)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("DeleteSkillByID", func() {
			exec := `DELETE FROM skills`
			It("all good", func() {
				mock.ExpectExec(exec).WithArgs(skill.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				err := db.DeleteSkillByID(skill.ID)
				Expect(err).ToNot(HaveOccurred())

			})

			It("no rows deleted", func() {
				mock.ExpectExec(exec).WithArgs(skill.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				err := db.DeleteSkillByID(skill.ID)
				Expect(err).To(Equal(database.ErrNoRowsDeleted))
			})

			It("handle error", func() {
				mock.ExpectExec(exec).WithArgs(skill.ID).WillReturnError(fmt.Errorf("error"))
				err := db.DeleteSkillByID(skill.ID)
				Expect(err).To(HaveOccurred())
			})
		})

	})
})
