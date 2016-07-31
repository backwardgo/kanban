package db_test

import (
	"github.com/backwardgo/kanban/db"
	"github.com/backwardgo/kanban/ids"
	"github.com/backwardgo/kanban/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserFilter", func() {
	var (
		txn   db.Transaction
		user1 models.User
		user2 models.User
		user3 models.User
	)

	BeforeEach(func() {
		txn = beginTransaction()

		createTestUser(txn, &user1)
		createTestUser(txn, &user2)
		createTestUser(txn, &user3)
	})

	AfterEach(func() {
		rollbackTransaction(txn)
	})

	Describe("filter: nothing; pager: nil", func() {
		var (
			count uint
			users []models.User
		)

		JustBeforeEach(func() {
			var err error
			var userFilter = db.UserFilter{}

			count, err = db.UserCount(txn, userFilter)
			Expect(err).To(BeNil())

			users, err = db.UserSlice(txn, userFilter, nil)
			Expect(err).To(BeNil())
		})

		Specify("count", func() {
			Expect(count).To(BeEquivalentTo(3))
		})

		Specify("users", func() {
			Expect(users).To(HaveLen(3))
		})
	})

	Describe("filter: nothing", func() {
		var (
			count uint
			users []models.User
		)

		JustBeforeEach(func() {
			var err error
			var userFilter = db.UserFilter{}

			count, err = db.UserCount(txn, userFilter)
			Expect(err).To(BeNil())

			users, err = db.UserSlice(txn, userFilter, nil)
			Expect(err).To(BeNil())
		})

		Specify("count", func() {
			Expect(count).To(BeEquivalentTo(3))
		})

		Specify("users", func() {
			Expect(users).To(HaveLen(3))
		})
	})

	Describe("filter: one UserIdIn; pager nil", func() {
		var (
			count uint
			users []models.User
		)

		JustBeforeEach(func() {
			var err error
			var userFilter = db.UserFilter{
				UserIdIn: ids.UserIdIn(user1.Id),
			}

			count, err = db.UserCount(txn, userFilter)
			Expect(err).To(BeNil())

			users, err = db.UserSlice(txn, userFilter, nil)
			Expect(err).To(BeNil())
		})

		Specify("count", func() {
			Expect(count).To(BeEquivalentTo(1))
		})

		Specify("users", func() {
			Expect(users).To(HaveLen(1))
			Expect(users).To(ContainElement(user1))
			Expect(users).ToNot(ContainElement(user2))
			Expect(users).ToNot(ContainElement(user3))
		})
	})

	Describe("filter: two UserIdIn; pager: nil", func() {
		var (
			count uint
			users []models.User
		)

		JustBeforeEach(func() {
			var err error
			var userFilter = db.UserFilter{
				UserIdIn: ids.UserIdIn(
					user2.Id,
					user3.Id,
				),
			}

			count, err = db.UserCount(txn, userFilter)
			Expect(err).To(BeNil())

			users, err = db.UserSlice(txn, userFilter, nil)
			Expect(err).To(BeNil())
		})

		Specify("count", func() {
			Expect(count).To(BeEquivalentTo(2))
		})

		Specify("users", func() {
			Expect(users).To(HaveLen(2))
			Expect(users).ToNot(ContainElement(user1))
			Expect(users).To(ContainElement(user2))
			Expect(users).To(ContainElement(user3))
		})
	})

})
