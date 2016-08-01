package db_test

import (
	"github.com/backwardgo/kanban/db"
	"github.com/backwardgo/kanban/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserByEmail", func() {

	Describe("Happy Path", func() {
		var txn db.Transaction
		BeforeEach(func() { txn = beginTransaction() })
		AfterEach(func() { rollbackTransaction(txn) })

		var (
			user models.User
			err  error
		)

		JustBeforeEach(func() {
			createTestUser(txn, &user)

			user, err = db.UserByEmail(txn, user.Email)
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
