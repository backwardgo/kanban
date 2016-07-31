package db_test

import (
	"github.com/backwardgo/kanban/db"
	"github.com/backwardgo/kanban/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BoardCreate", func() {

	Describe("Happy Path", func() {
		var (
			txn   db.Transaction
			board models.Board
		)

		BeforeEach(func() { txn = beginTransaction() })

		AfterEach(func() { rollbackTransaction(txn) })

		JustBeforeEach(func() {
			var user models.User
			createTestUser(txn, &user)

			var team models.Team
			team.CreatedBy = user.Id
			createTestTeam(txn, &team)

			board.Name = "Welcome Board"
			board.TeamId = team.Id
			board.CreatedBy = user.Id

			err := db.BoardCreate(txn, &board)
			Expect(err).To(BeNil())
		})

		Specify("board.Id", func() {
			Expect(board.Id.Blank()).To(BeFalse())
		})

		Specify("board.CreatedAt", func() {
			Expect(board.CreatedAt.IsZero()).To(BeFalse())
			Expect(board.CreatedAt).To(BeTemporally("==", board.UpdatedAt))
		})

		Specify("board.DeletedAt", func() {
			Expect(board.DeletedAt).To(BeNil())
		})

		Specify("board.UpdatedAt", func() {
			Expect(board.UpdatedAt.IsZero()).To(BeFalse())
			Expect(board.UpdatedAt).To(BeTemporally("==", board.CreatedAt))
		})
	})

})
