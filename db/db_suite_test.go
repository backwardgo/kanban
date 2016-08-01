package db_test

import (
	"fmt"
	"sync/atomic"

	"github.com/backwardgo/kanban/db"
	"github.com/backwardgo/kanban/models"
	"github.com/backwardgo/kanban/models/roles"
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
	testIdSequence   uint64
)

func nextTestEmail() models.Email {
	atomic.AddUint64(&testEmailCounter, 1)
	return models.Email(fmt.Sprintf("email-%v@fake.com", testEmailCounter))
}

func nextTestId() uint64 {
	atomic.AddUint64(&testIdSequence, 1)
	return testIdSequence
}

// -------

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

	err := db.CreateUser(txn, m)
	Expect(err).To(BeNil())
}

func createTestList(txn db.Transaction, m *models.List) {
	if m.Title == "" {
		m.Title = fmt.Sprintf("List %v", nextTestId())
	}

	err := db.CreateList(txn, m)
	Expect(err).To(BeNil())
}

func createTestBoard(txn db.Transaction, m *models.Board) {
	if m.Name == "" {
		m.Name = fmt.Sprintf("Board %v", nextTestId())
	}

	err := db.CreateBoard(txn, m)
	Expect(err).To(BeNil())
}

func createTestMember(txn db.Transaction, m *models.Member) {
	if m.Role.Blank() {
		m.Role = roles.Default
	}

	err := db.CreateMember(txn, m)
	Expect(err).To(BeNil())
}

func createTestCard(txn db.Transaction, m *models.Card) {
	if m.Title == "" {
		m.Title = fmt.Sprintf("Card %v", nextTestId())
	}

	err := db.CreateCard(txn, m)
	Expect(err).To(BeNil())
}

// -------

type basicTeamFixture struct {
	User1 models.User
	User2 models.User
	User3 models.User

	BoardA models.Board // CreatedBy User1
	BoardB models.Board // CreatedBy User2
	BoardC models.Board // CreatedBy User3

	MemberA1 models.Member // BoardA; User1; admin
	MemberA2 models.Member // BoardA; User2; default
	MemberA3 models.Member // BoardA; User3; observer

	MemberB1 models.Member // BoardB; User1; observer
	MemberB2 models.Member // BoardB; User2; admin
	MemberB3 models.Member // BoardB; User3; default

	MemberC1 models.Member // BoardC; User1; default
	MemberC2 models.Member // BoardC; User2; observer
	MemberC3 models.Member // BoardC; User3; admin

	// ListA* belongs to BoardA
	ListA1 models.List // CreatedBy User1
	ListA2 models.List // CreatedBy User2
	ListA3 models.List // CreatedBy User3

	// ListB* belongs to BoardB
	ListB1 models.List // CreatedBy User1
	ListB2 models.List // CreatedBy User2
	ListB3 models.List // CreatedBy User3

	// ListC* belongs to BoardC
	ListC1 models.List // CreatedBy User1
	ListC2 models.List // CreatedBy User2
	ListC3 models.List // CreatedBy User3

	// CardA1* belongs to ListA1
	CardA11 models.Card // CreatedBy User1
	CardA12 models.Card // CreatedBy User2
	CardA13 models.Card // CreatedBy User3

	// CardA2* belongs to ListA2
	CardA21 models.Card // CreatedBy User1
	CardA22 models.Card // CreatedBy User2
	CardA23 models.Card // CreatedBy User3

	// CardA3* belongs to ListA3
	CardA31 models.Card // CreatedBy User1
	CardA32 models.Card // CreatedBy User2
	CardA33 models.Card // CreatedBy User3

	// CardB1* belongs to ListB1
	CardB11 models.Card // CreatedBy User1
	CardB12 models.Card // CreatedBy User2
	CardB13 models.Card // CreatedBy User3

	// CardB2* belongs to ListB2
	CardB21 models.Card // CreatedBy User1
	CardB22 models.Card // CreatedBy User2
	CardB23 models.Card // CreatedBy User3

	// CardB3* belongs to ListB3
	CardB31 models.Card // CreatedBy User1
	CardB32 models.Card // CreatedBy User2
	CardB33 models.Card // CreatedBy User3

	// CardC1* belongs to ListC1
	CardC11 models.Card // CreatedBy User1
	CardC12 models.Card // CreatedBy User2
	CardC13 models.Card // CreatedBy User3

	// CardC2* belongs to ListC2
	CardC21 models.Card // CreatedBy User1
	CardC22 models.Card // CreatedBy User2
	CardC23 models.Card // CreatedBy User3

	// CardC3* belongs to ListC3
	CardC31 models.Card // CreatedBy User1
	CardC32 models.Card // CreatedBy User2
	CardC33 models.Card // CreatedBy User3
}

func createBasicTeamFixture(txn db.Transaction) (m basicTeamFixture) {
	{ // create the users
		createTestUser(txn, &m.User1)
		createTestUser(txn, &m.User2)
		createTestUser(txn, &m.User3)
	}

	{ // create the boards
		m.BoardA.CreatedBy = m.User1.Id
		m.BoardB.CreatedBy = m.User2.Id
		m.BoardC.CreatedBy = m.User3.Id

		createTestBoard(txn, &m.BoardA)
		createTestBoard(txn, &m.BoardB)
		createTestBoard(txn, &m.BoardC)
	}

	{ // create the members
		{ // for BoardA
			m.MemberA1.BoardId = m.BoardA.Id
			m.MemberA2.BoardId = m.BoardA.Id
			m.MemberA3.BoardId = m.BoardA.Id

			m.MemberA1.Role = roles.Admin
			m.MemberA2.Role = roles.Default
			m.MemberA3.Role = roles.Observer

			m.MemberA1.UserId = m.User1.Id
			m.MemberA2.UserId = m.User2.Id
			m.MemberA3.UserId = m.User3.Id

			createTestMember(txn, &m.MemberA1)
			createTestMember(txn, &m.MemberA2)
			createTestMember(txn, &m.MemberA3)
		}
		{ // for BoardB
			m.MemberB1.BoardId = m.BoardB.Id
			m.MemberB2.BoardId = m.BoardB.Id
			m.MemberB3.BoardId = m.BoardB.Id

			m.MemberB1.Role = roles.Observer
			m.MemberB2.Role = roles.Admin
			m.MemberB3.Role = roles.Default

			m.MemberB1.UserId = m.User1.Id
			m.MemberB2.UserId = m.User2.Id
			m.MemberB3.UserId = m.User3.Id

			createTestMember(txn, &m.MemberB1)
			createTestMember(txn, &m.MemberB2)
			createTestMember(txn, &m.MemberB3)
		}
		{ // for BoardC
			m.MemberC1.BoardId = m.BoardC.Id
			m.MemberC2.BoardId = m.BoardC.Id
			m.MemberC3.BoardId = m.BoardC.Id

			m.MemberC1.Role = roles.Default
			m.MemberC2.Role = roles.Observer
			m.MemberC3.Role = roles.Admin

			m.MemberC1.UserId = m.User1.Id
			m.MemberC2.UserId = m.User2.Id
			m.MemberC3.UserId = m.User3.Id

			createTestMember(txn, &m.MemberC1)
			createTestMember(txn, &m.MemberC2)
			createTestMember(txn, &m.MemberC3)
		}
	}

	{ // create the lists
		m.ListA1.BoardId = m.BoardA.Id
		m.ListA2.BoardId = m.BoardA.Id
		m.ListA3.BoardId = m.BoardA.Id
		m.ListB1.BoardId = m.BoardB.Id
		m.ListB2.BoardId = m.BoardB.Id
		m.ListB3.BoardId = m.BoardB.Id
		m.ListC1.BoardId = m.BoardC.Id
		m.ListC2.BoardId = m.BoardC.Id
		m.ListC3.BoardId = m.BoardC.Id

		m.ListA1.CreatedBy = m.User1.Id
		m.ListA2.CreatedBy = m.User1.Id
		m.ListA3.CreatedBy = m.User1.Id
		m.ListB1.CreatedBy = m.User2.Id
		m.ListB2.CreatedBy = m.User2.Id
		m.ListB3.CreatedBy = m.User2.Id
		m.ListC1.CreatedBy = m.User3.Id
		m.ListC2.CreatedBy = m.User3.Id
		m.ListC3.CreatedBy = m.User3.Id

		createTestList(txn, &m.ListA1)
		createTestList(txn, &m.ListA2)
		createTestList(txn, &m.ListA3)
		createTestList(txn, &m.ListB1)
		createTestList(txn, &m.ListB2)
		createTestList(txn, &m.ListB3)
		createTestList(txn, &m.ListC1)
		createTestList(txn, &m.ListC2)
		createTestList(txn, &m.ListC3)
	}

	{ // create the cards
		{ // for ListA1
			m.CardA11.ListId = m.ListA1.Id
			m.CardA12.ListId = m.ListA1.Id
			m.CardA13.ListId = m.ListA1.Id
			m.CardA11.CreatedBy = m.User1.Id
			m.CardA12.CreatedBy = m.User2.Id
			m.CardA13.CreatedBy = m.User3.Id

			createTestCard(txn, &m.CardA11)
			createTestCard(txn, &m.CardA12)
			createTestCard(txn, &m.CardA13)
		}
		{ // for ListA2
			m.CardA21.ListId = m.ListA2.Id
			m.CardA22.ListId = m.ListA2.Id
			m.CardA23.ListId = m.ListA2.Id
			m.CardA21.CreatedBy = m.User1.Id
			m.CardA22.CreatedBy = m.User2.Id
			m.CardA23.CreatedBy = m.User3.Id

			createTestCard(txn, &m.CardA21)
			createTestCard(txn, &m.CardA22)
			createTestCard(txn, &m.CardA23)
		}
		{ // for ListA3
			m.CardA31.ListId = m.ListA3.Id
			m.CardA32.ListId = m.ListA3.Id
			m.CardA33.ListId = m.ListA3.Id
			m.CardA31.CreatedBy = m.User1.Id
			m.CardA32.CreatedBy = m.User2.Id
			m.CardA33.CreatedBy = m.User3.Id

			createTestCard(txn, &m.CardA31)
			createTestCard(txn, &m.CardA32)
			createTestCard(txn, &m.CardA33)
		}
		{ // for ListB1
			m.CardB11.ListId = m.ListB1.Id
			m.CardB12.ListId = m.ListB1.Id
			m.CardB13.ListId = m.ListB1.Id
			m.CardB11.CreatedBy = m.User1.Id
			m.CardB12.CreatedBy = m.User2.Id
			m.CardB13.CreatedBy = m.User3.Id

			createTestCard(txn, &m.CardB11)
			createTestCard(txn, &m.CardB12)
			createTestCard(txn, &m.CardB13)
		}
		{ // for ListB2
			m.CardB21.ListId = m.ListB2.Id
			m.CardB22.ListId = m.ListB2.Id
			m.CardB23.ListId = m.ListB2.Id
			m.CardB21.CreatedBy = m.User1.Id
			m.CardB22.CreatedBy = m.User2.Id
			m.CardB23.CreatedBy = m.User3.Id

			createTestCard(txn, &m.CardB21)
			createTestCard(txn, &m.CardB22)
			createTestCard(txn, &m.CardB23)
		}
		{ // for ListB3
			m.CardB31.ListId = m.ListB3.Id
			m.CardB32.ListId = m.ListB3.Id
			m.CardB33.ListId = m.ListB3.Id
			m.CardB31.CreatedBy = m.User1.Id
			m.CardB32.CreatedBy = m.User2.Id
			m.CardB33.CreatedBy = m.User3.Id

			createTestCard(txn, &m.CardB31)
			createTestCard(txn, &m.CardB32)
			createTestCard(txn, &m.CardB33)
		}
		{ // for ListC1
			m.CardC11.ListId = m.ListC1.Id
			m.CardC12.ListId = m.ListC1.Id
			m.CardC13.ListId = m.ListC1.Id
			m.CardC11.CreatedBy = m.User1.Id
			m.CardC12.CreatedBy = m.User2.Id
			m.CardC13.CreatedBy = m.User3.Id

			createTestCard(txn, &m.CardC11)
			createTestCard(txn, &m.CardC12)
			createTestCard(txn, &m.CardC13)
		}
		{ // for ListC2
			m.CardC21.ListId = m.ListC2.Id
			m.CardC22.ListId = m.ListC2.Id
			m.CardC23.ListId = m.ListC2.Id
			m.CardC21.CreatedBy = m.User1.Id
			m.CardC22.CreatedBy = m.User2.Id
			m.CardC23.CreatedBy = m.User3.Id

			createTestCard(txn, &m.CardC21)
			createTestCard(txn, &m.CardC22)
			createTestCard(txn, &m.CardC23)
		}
		{ // for ListC3
			m.CardC31.ListId = m.ListC3.Id
			m.CardC32.ListId = m.ListC3.Id
			m.CardC33.ListId = m.ListC3.Id
			m.CardC31.CreatedBy = m.User1.Id
			m.CardC32.CreatedBy = m.User2.Id
			m.CardC33.CreatedBy = m.User3.Id

			createTestCard(txn, &m.CardC31)
			createTestCard(txn, &m.CardC32)
			createTestCard(txn, &m.CardC33)
		}
	}

	return m
}
