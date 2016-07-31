package db_test

import (
	"fmt"
	"sync/atomic"

	"github.com/backwardgo/kanban/db"
	"github.com/backwardgo/kanban/models"
	"github.com/icrowley/fake"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Db Suite")
}

var testDB *sqlx.DB

var _ = BeforeSuite(func() {
	var err error

	var databaseURL = `postgres://localhost/kanban_test?sslmode=disable`
	db.MigrateRedo(databaseURL)

	testDB, err = sqlx.Connect(`postgres`, databaseURL)
	Expect(err).To(BeNil())

	truncateEverything()
})

var _ = AfterSuite(func() {
	err := testDB.Close()
	Expect(err).To(BeNil())
})

func beginTransaction() db.Transaction {
	Expect(testDB).ToNot(BeNil())

	con := db.NewConnectionWithDB(testDB)

	txn, err := con.Begin()
	Expect(err).To(BeNil())

	return txn
}

func commitTransaction(txn db.Transaction) {
	err := txn.Commit()
	Expect(err).To(BeNil())
}

func rollbackTransaction(txn db.Transaction) {
	err := txn.Rollback()
	Expect(err).To(BeNil())
}

func truncateEverything() {
	testDB.MustExec(`
		TRUNCATE users CASCADE;
	`)
}

// -------
var (
	testEmailCounter uint64
)

func nextTestEmail() models.Email {
	atomic.AddUint64(&testEmailCounter, 1)
	return models.Email(fmt.Sprintf("email-%v@fake.com", testEmailCounter))
}

func createTestUser(txn db.Transaction, m *models.User) {
	if m.FirstName == "" {
		m.FirstName = fake.FirstName()
	}

	if m.LastName == "" {
		m.LastName = fake.LastName()
	}

	if m.Email.Blank() {
		m.Email = nextTestEmail()
	}

	err := db.UserCreate(txn, m)
	Expect(err).To(BeNil())
}

func createTestTeam(txn db.Transaction, m *models.Team) {
	if m.Name == "" {
		m.Name = fake.Company()
	}

	err := db.TeamCreate(txn, m)
	Expect(err).To(BeNil())
}
