package db_test

import (
	"github.com/backwardgo/kanban/db"
	"github.com/backwardgo/kanban/models"
	"github.com/backwardgo/kanban/test/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateCard", func() {

	Describe("Happy Path", func() {
		var (
			txn  db.Transaction
			card models.Card
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

			var board models.Board
			board.CreatedBy = user.Id
			helpers.CreateBoard(txn, &board)

			var list models.List
			list.BoardId = board.Id
			list.CreatedBy = user.Id
			helpers.CreateList(txn, &list)

			card.Title = "Hello Card!"
			card.ListId = list.Id
			card.CreatedBy = user.Id

			err := db.CreateCard(txn, &card)
			Expect(err).To(BeNil())
		})

		Specify("card.Id", func() {
			Expect(card.Id.Blank()).To(BeFalse())
		})

		Specify("card.CreatedAt", func() {
			Expect(card.CreatedAt.IsZero()).To(BeFalse())
			Expect(card.CreatedAt).To(BeTemporally("==", card.UpdatedAt))
		})

		Specify("card.DeletedAt", func() {
			Expect(card.DeletedAt).To(BeNil())
		})

		Specify("card.UpdatedAt", func() {
			Expect(card.UpdatedAt.IsZero()).To(BeFalse())
			Expect(card.UpdatedAt).To(BeTemporally("==", card.CreatedAt))
		})
	})

})
