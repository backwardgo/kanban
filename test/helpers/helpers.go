package helpers

import (
	"fmt"
	"sync/atomic"

	"github.com/backwardgo/kanban/db"
	"github.com/backwardgo/kanban/models"
	"github.com/backwardgo/kanban/models/roles"
	"github.com/backwardgo/kanban/test/fixtures"
	"github.com/icrowley/fake"

	. "github.com/onsi/gomega"
)

// -------

var (
	testEmailCounter uint64
	testIdSequence   uint64
)

func nextTestEmail() models.Email {
	atomic.AddUint64(&testEmailCounter, 1)
	return models.Email(fmt.Sprintf("test-email-%v@fake.com", testEmailCounter))
}

func nextTestId() uint64 {
	atomic.AddUint64(&testIdSequence, 1)
	return testIdSequence
}

// -------

func CreateUser(txn db.Transaction, m *models.User) {
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

func CreateList(txn db.Transaction, m *models.List) {
	if m.Title == "" {
		m.Title = fmt.Sprintf("List %v", nextTestId())
	}

	err := db.CreateList(txn, m)
	Expect(err).To(BeNil())
}

func CreateBoard(txn db.Transaction, m *models.Board) {
	if m.Name == "" {
		m.Name = fmt.Sprintf("Board %v", nextTestId())
	}

	err := db.CreateBoard(txn, m)
	Expect(err).To(BeNil())
}

func CreateMember(txn db.Transaction, m *models.Member) {
	if m.Role.Blank() {
		m.Role = roles.Default
	}

	err := db.CreateMember(txn, m)
	Expect(err).To(BeNil())
}

func CreateCard(txn db.Transaction, m *models.Card) {
	if m.Title == "" {
		m.Title = fmt.Sprintf("Card %v", nextTestId())
	}

	err := db.CreateCard(txn, m)
	Expect(err).To(BeNil())
}

func CreateBasicTeam(txn db.Transaction) (m fixtures.BasicTeam) {
	{ // create the users
		CreateUser(txn, &m.User1)
		CreateUser(txn, &m.User2)
		CreateUser(txn, &m.User3)
	}

	{ // create the boards
		m.BoardA.CreatedBy = m.User1.Id
		m.BoardB.CreatedBy = m.User2.Id
		m.BoardC.CreatedBy = m.User3.Id

		CreateBoard(txn, &m.BoardA)
		CreateBoard(txn, &m.BoardB)
		CreateBoard(txn, &m.BoardC)
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

			CreateMember(txn, &m.MemberA1)
			CreateMember(txn, &m.MemberA2)
			CreateMember(txn, &m.MemberA3)
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

			CreateMember(txn, &m.MemberB1)
			CreateMember(txn, &m.MemberB2)
			CreateMember(txn, &m.MemberB3)
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

			CreateMember(txn, &m.MemberC1)
			CreateMember(txn, &m.MemberC2)
			CreateMember(txn, &m.MemberC3)
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

		CreateList(txn, &m.ListA1)
		CreateList(txn, &m.ListA2)
		CreateList(txn, &m.ListA3)
		CreateList(txn, &m.ListB1)
		CreateList(txn, &m.ListB2)
		CreateList(txn, &m.ListB3)
		CreateList(txn, &m.ListC1)
		CreateList(txn, &m.ListC2)
		CreateList(txn, &m.ListC3)
	}

	{ // create the cards
		{ // for ListA1
			m.CardA11.ListId = m.ListA1.Id
			m.CardA12.ListId = m.ListA1.Id
			m.CardA13.ListId = m.ListA1.Id
			m.CardA11.CreatedBy = m.User1.Id
			m.CardA12.CreatedBy = m.User2.Id
			m.CardA13.CreatedBy = m.User3.Id

			CreateCard(txn, &m.CardA11)
			CreateCard(txn, &m.CardA12)
			CreateCard(txn, &m.CardA13)
		}
		{ // for ListA2
			m.CardA21.ListId = m.ListA2.Id
			m.CardA22.ListId = m.ListA2.Id
			m.CardA23.ListId = m.ListA2.Id
			m.CardA21.CreatedBy = m.User1.Id
			m.CardA22.CreatedBy = m.User2.Id
			m.CardA23.CreatedBy = m.User3.Id

			CreateCard(txn, &m.CardA21)
			CreateCard(txn, &m.CardA22)
			CreateCard(txn, &m.CardA23)
		}
		{ // for ListA3
			m.CardA31.ListId = m.ListA3.Id
			m.CardA32.ListId = m.ListA3.Id
			m.CardA33.ListId = m.ListA3.Id
			m.CardA31.CreatedBy = m.User1.Id
			m.CardA32.CreatedBy = m.User2.Id
			m.CardA33.CreatedBy = m.User3.Id

			CreateCard(txn, &m.CardA31)
			CreateCard(txn, &m.CardA32)
			CreateCard(txn, &m.CardA33)
		}
		{ // for ListB1
			m.CardB11.ListId = m.ListB1.Id
			m.CardB12.ListId = m.ListB1.Id
			m.CardB13.ListId = m.ListB1.Id
			m.CardB11.CreatedBy = m.User1.Id
			m.CardB12.CreatedBy = m.User2.Id
			m.CardB13.CreatedBy = m.User3.Id

			CreateCard(txn, &m.CardB11)
			CreateCard(txn, &m.CardB12)
			CreateCard(txn, &m.CardB13)
		}
		{ // for ListB2
			m.CardB21.ListId = m.ListB2.Id
			m.CardB22.ListId = m.ListB2.Id
			m.CardB23.ListId = m.ListB2.Id
			m.CardB21.CreatedBy = m.User1.Id
			m.CardB22.CreatedBy = m.User2.Id
			m.CardB23.CreatedBy = m.User3.Id

			CreateCard(txn, &m.CardB21)
			CreateCard(txn, &m.CardB22)
			CreateCard(txn, &m.CardB23)
		}
		{ // for ListB3
			m.CardB31.ListId = m.ListB3.Id
			m.CardB32.ListId = m.ListB3.Id
			m.CardB33.ListId = m.ListB3.Id
			m.CardB31.CreatedBy = m.User1.Id
			m.CardB32.CreatedBy = m.User2.Id
			m.CardB33.CreatedBy = m.User3.Id

			CreateCard(txn, &m.CardB31)
			CreateCard(txn, &m.CardB32)
			CreateCard(txn, &m.CardB33)
		}
		{ // for ListC1
			m.CardC11.ListId = m.ListC1.Id
			m.CardC12.ListId = m.ListC1.Id
			m.CardC13.ListId = m.ListC1.Id
			m.CardC11.CreatedBy = m.User1.Id
			m.CardC12.CreatedBy = m.User2.Id
			m.CardC13.CreatedBy = m.User3.Id

			CreateCard(txn, &m.CardC11)
			CreateCard(txn, &m.CardC12)
			CreateCard(txn, &m.CardC13)
		}
		{ // for ListC2
			m.CardC21.ListId = m.ListC2.Id
			m.CardC22.ListId = m.ListC2.Id
			m.CardC23.ListId = m.ListC2.Id
			m.CardC21.CreatedBy = m.User1.Id
			m.CardC22.CreatedBy = m.User2.Id
			m.CardC23.CreatedBy = m.User3.Id

			CreateCard(txn, &m.CardC21)
			CreateCard(txn, &m.CardC22)
			CreateCard(txn, &m.CardC23)
		}
		{ // for ListC3
			m.CardC31.ListId = m.ListC3.Id
			m.CardC32.ListId = m.ListC3.Id
			m.CardC33.ListId = m.ListC3.Id
			m.CardC31.CreatedBy = m.User1.Id
			m.CardC32.CreatedBy = m.User2.Id
			m.CardC33.CreatedBy = m.User3.Id

			CreateCard(txn, &m.CardC31)
			CreateCard(txn, &m.CardC32)
			CreateCard(txn, &m.CardC33)
		}
	}

	return m
}
