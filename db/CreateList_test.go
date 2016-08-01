package db_test

import (
	"github.com/backwardgo/kanban/db"
	"github.com/backwardgo/kanban/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateList", func() {

	Describe("Happy Path", func() {
		var (
			txn  db.Transaction
			list models.List
		)

		BeforeEach(func() { txn = beginTransaction() })

		AfterEach(func() { rollbackTransaction(txn) })

		JustBeforeEach(func() {
			var user models.User
			createTestUser(txn, &user)

			var board models.Board
			board.CreatedBy = user.Id
			createTestBoard(txn, &board)

			list.Title = "Hello List!"
			list.BoardId = board.Id
			list.CreatedBy = user.Id

			err := db.CreateList(txn, &list)
			Expect(err).To(BeNil())
		})

		Specify("list.Id", func() {
			Expect(list.Id.Blank()).To(BeFalse())
		})

		Specify("list.CreatedAt", func() {
			Expect(list.CreatedAt.IsZero()).To(BeFalse())
			Expect(list.CreatedAt).To(BeTemporally("==", list.UpdatedAt))
		})

		Specify("list.DeletedAt", func() {
			Expect(list.DeletedAt).To(BeNil())
		})

		Specify("list.UpdatedAt", func() {
			Expect(list.UpdatedAt.IsZero()).To(BeFalse())
			Expect(list.UpdatedAt).To(BeTemporally("==", list.CreatedAt))
		})
	})

})
