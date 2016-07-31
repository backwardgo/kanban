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

	Describe("filter: nothing", func() {
		var userFilter db.UserFilter

		Context("UserCount", func() {
			var (
				count uint
				err   error
			)

			JustBeforeEach(func() {
				count, err = db.UserCount(txn, userFilter)
			})

			Specify("count", func() {
				Expect(count).To(BeEquivalentTo(3))
			})

			Specify("err", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("UserSlice: pager: nil; orderBy: nil", func() {
			var (
				users []models.User
				err   error
			)

			JustBeforeEach(func() {
				users, err = db.UserSlice(txn, userFilter, nil)
			})

			Specify("users", func() {
				Expect(users).To(HaveLen(3))
				Expect(users).To(ConsistOf(
					user1,
					user2,
					user3,
				))
			})

			Specify("err", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("filter: one UserIdIn", func() {
		var userFilter db.UserFilter

		BeforeEach(func() {
			userFilter = db.UserFilter{
				UserIdIn: ids.UserIdIn(
					user1.Id,
				),
			}
		})

		Context("UserCount", func() {
			var (
				count uint
				err   error
			)

			JustBeforeEach(func() {
				count, err = db.UserCount(txn, userFilter)
			})

			Specify("count", func() {
				Expect(count).To(BeEquivalentTo(1))
			})

			Specify("err", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("UserSlice: pager: nil; orderBy: nil", func() {
			var (
				users []models.User
				err   error
			)

			JustBeforeEach(func() {
				users, err = db.UserSlice(txn, userFilter, nil)
			})

			Specify("users", func() {
				Expect(users).To(HaveLen(1))
				Expect(users).To(ConsistOf(
					user1,
				))
			})

			Specify("err", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("filter: two UserIdIn", func() {
		var userFilter db.UserFilter

		BeforeEach(func() {
			userFilter = db.UserFilter{
				UserIdIn: ids.UserIdIn(
					user2.Id,
					user3.Id,
				),
			}
		})

		Context("UserCount", func() {
			var (
				count uint
				err   error
			)

			JustBeforeEach(func() {
				count, err = db.UserCount(txn, userFilter)
			})

			Specify("count", func() {
				Expect(count).To(BeEquivalentTo(2))
			})

			Specify("err", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("UserSlice: pager: nil; orderBy: nil", func() {
			var (
				users []models.User
				err   error
			)

			JustBeforeEach(func() {
				users, err = db.UserSlice(txn, userFilter, nil)
			})

			Specify("users", func() {
				Expect(users).To(HaveLen(2))
				Expect(users).To(ConsistOf(
					user2,
					user3,
				))
			})

			Specify("err", func() {
				Expect(err).To(BeNil())
			})
		})
	})

})
