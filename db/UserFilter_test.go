package db_test

import (
	"github.com/backwardgo/kanban/db"
	"github.com/backwardgo/kanban/ids"
	"github.com/backwardgo/kanban/models"
	"github.com/backwardgo/kanban/test/helpers"

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

		helpers.CreateUser(txn, &user1)
		helpers.CreateUser(txn, &user2)
		helpers.CreateUser(txn, &user3)
	})

	AfterEach(func() {
		rollbackTransaction(txn)
	})

	Describe("filter: nothing", func() {
		var userFilter db.UserFilter

		Context("CountUsers", func() {
			var (
				count uint
				err   error
			)

			JustBeforeEach(func() {
				count, err = db.CountUsers(txn, userFilter)
			})

			Specify("count", func() {
				Expect(count).To(BeEquivalentTo(3))
			})

			Specify("err", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("SelectUsers: pager: nil; orderBy: nil", func() {
			var (
				users []models.User
				err   error
			)

			JustBeforeEach(func() {
				users, err = db.SelectUsers(txn, userFilter, nil)
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

		Context("SelectUsers: pager: {page: 1, perPage: 1}; orderBy: nil", func() {
			var (
				pager db.Pager
				users []models.User
				err   error
			)

			JustBeforeEach(func() {
				pager = db.NewPager(1, 1)
				users, err = db.SelectUsers(txn, userFilter, pager)
			})

			Specify("pager.Page", func() {
				Expect(pager.Page()).To(BeEquivalentTo(1))
			})

			Specify("pager.PerPage", func() {
				Expect(pager.PerPage()).To(BeEquivalentTo(1))
			})

			Specify("pager.TotalRecords", func() {
				Expect(pager.TotalRecords()).To(BeEquivalentTo(3))
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

		Context("SelectUsers: pager: {page: 2, perPage: 1}; orderBy: nil", func() {
			var (
				pager db.Pager
				users []models.User
				err   error
			)

			JustBeforeEach(func() {
				pager = db.NewPager(2, 1)
				users, err = db.SelectUsers(txn, userFilter, pager)
			})

			Specify("pager.Page", func() {
				Expect(pager.Page()).To(BeEquivalentTo(2))
			})

			Specify("pager.PerPage", func() {
				Expect(pager.PerPage()).To(BeEquivalentTo(1))
			})

			Specify("pager.TotalRecords", func() {
				Expect(pager.TotalRecords()).To(BeEquivalentTo(3))
			})

			Specify("users", func() {
				Expect(users).To(HaveLen(1))
				Expect(users).To(ConsistOf(
					user2,
				))
			})

			Specify("err", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("SelectUsers: pager: {page: 3, perPage: 1}; orderBy: nil", func() {
			var (
				pager db.Pager
				users []models.User
				err   error
			)

			JustBeforeEach(func() {
				pager = db.NewPager(3, 1)
				users, err = db.SelectUsers(txn, userFilter, pager)
			})

			Specify("pager.Page", func() {
				Expect(pager.Page()).To(BeEquivalentTo(3))
			})

			Specify("pager.PerPage", func() {
				Expect(pager.PerPage()).To(BeEquivalentTo(1))
			})

			Specify("pager.TotalRecords", func() {
				Expect(pager.TotalRecords()).To(BeEquivalentTo(3))
			})

			Specify("users", func() {
				Expect(users).To(HaveLen(1))
				Expect(users).To(ConsistOf(
					user3,
				))
			})

			Specify("err", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("SelectUsers: pager: {page: 4, perPage: 1}; orderBy: nil", func() {
			var (
				pager db.Pager
				users []models.User
				err   error
			)

			JustBeforeEach(func() {
				pager = db.NewPager(4, 1)
				users, err = db.SelectUsers(txn, userFilter, pager)
			})

			Specify("pager.Page", func() {
				Expect(pager.Page()).To(BeEquivalentTo(4))
			})

			Specify("pager.PerPage", func() {
				Expect(pager.PerPage()).To(BeEquivalentTo(1))
			})

			Specify("pager.TotalRecords", func() {
				Expect(pager.TotalRecords()).To(BeEquivalentTo(3))
			})

			Specify("users", func() {
				Expect(users).To(HaveLen(0))
			})

			Specify("err", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("SelectUsers: pager: {page: 0, perPage: 0}; orderBy: nil", func() {
			var (
				pager db.Pager
				users []models.User
				err   error
			)

			JustBeforeEach(func() {
				pager = db.NewPager(0, 0)
				users, err = db.SelectUsers(txn, userFilter, pager)
			})

			Specify("pager.Page", func() {
				Expect(pager.Page()).To(BeEquivalentTo(1))
			})

			Specify("pager.PerPage", func() {
				Expect(pager.PerPage()).To(BeEquivalentTo(200))
			})

			Specify("pager.TotalRecords", func() {
				Expect(pager.TotalRecords()).To(BeEquivalentTo(3))
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

		Context("CountUsers", func() {
			var (
				count uint
				err   error
			)

			JustBeforeEach(func() {
				count, err = db.CountUsers(txn, userFilter)
			})

			Specify("count", func() {
				Expect(count).To(BeEquivalentTo(1))
			})

			Specify("err", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("SelectUsers: pager: nil; orderBy: nil", func() {
			var (
				users []models.User
				err   error
			)

			JustBeforeEach(func() {
				users, err = db.SelectUsers(txn, userFilter, nil)
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

		Context("CountUsers", func() {
			var (
				count uint
				err   error
			)

			JustBeforeEach(func() {
				count, err = db.CountUsers(txn, userFilter)
			})

			Specify("count", func() {
				Expect(count).To(BeEquivalentTo(2))
			})

			Specify("err", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("SelectUsers: pager: nil; orderBy: nil", func() {
			var (
				users []models.User
				err   error
			)

			JustBeforeEach(func() {
				users, err = db.SelectUsers(txn, userFilter, nil)
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
