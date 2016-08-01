package fixtures

import "github.com/backwardgo/kanban/models"

type BasicTeam struct {
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
