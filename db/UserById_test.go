package db_test

import (
	"github.com/backwardgo/kanban/db"
	"github.com/backwardgo/kanban/models"
	"github.com/backwardgo/kanban/test/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserById", func() {

	Describe("Happy Path", func() {
		var (
			txn  db.Transaction
			user models.User
			err  error
		)

		BeforeEach(func() {
			txn = beginTransaction()
		})

		AfterEach(func() {
			rollbackTransaction(txn)
		})

		JustBeforeEach(func() {
			helpers.CreateUser(txn, &user)
			user, err = db.UserById(txn, user.Id)
		})

		Specify("err", func() {
			Expect(err).To(BeNil())
		})

		Specify("user.Id", func() {
			Expect(user.Id.Blank()).To(BeFalse())
		})

		Specify("user.CreatedAt", func() {
			Expect(user.CreatedAt.IsZero()).To(BeFalse())
			Expect(user.CreatedAt).To(BeTemporally("==", user.UpdatedAt))
		})

		Specify("user.UpdatedAt", func() {
			Expect(user.UpdatedAt.IsZero()).To(BeFalse())
			Expect(user.UpdatedAt).To(BeTemporally("==", user.CreatedAt))
		})
	})

})
