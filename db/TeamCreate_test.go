package db_test

import (
	"github.com/backwardgo/kanban/db"
	"github.com/backwardgo/kanban/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TeamCreate", func() {

	Describe("Happy Path", func() {
		var (
			txn  db.Transaction
			team models.Team
		)

		BeforeEach(func() { txn = beginTransaction() })

		AfterEach(func() { rollbackTransaction(txn) })

		JustBeforeEach(func() {
			var user models.User
			createTestUser(txn, &user)

			team.Name = "The Simpsons"
			team.CreatedBy = user.Id

			err := db.TeamCreate(txn, &team)
			Expect(err).To(BeNil())
		})

		Specify("team.Id", func() {
			Expect(team.Id.Blank()).To(BeFalse())
		})

		Specify("team.CreatedAt", func() {
			Expect(team.CreatedAt.IsZero()).To(BeFalse())
			Expect(team.CreatedAt).To(BeTemporally("==", team.UpdatedAt))
		})

		Specify("team.DeletedAt", func() {
			Expect(team.DeletedAt).To(BeNil())
		})

		Specify("team.UpdatedAt", func() {
			Expect(team.UpdatedAt.IsZero()).To(BeFalse())
			Expect(team.UpdatedAt).To(BeTemporally("==", team.CreatedAt))
		})
	})

})
