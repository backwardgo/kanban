package db_test

import (
	"github.com/backwardgo/kanban/db"
	"github.com/backwardgo/kanban/models"
	"github.com/backwardgo/kanban/test/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateBoard", func() {

	Describe("Happy Path", func() {
		var (
			txn   db.Transaction
			board models.Board
		)

		BeforeEach(func() {
			txn = beginTransaction()
		})

		AfterEach(func() {
			rollbackTransaction(txn)
		})

		JustBeforeEach(func() {
			var user models.User
			helpers.CreateUser(txn, &user)

			board.Name = "Welcome Board"
			board.CreatedBy = user.Id

			err := db.CreateBoard(txn, &board)
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
