package fan

type (
	GroupType struct {
		Public    bool
		CardCount int
	}
)

var (
	CHI = GroupType{
		Public:    true,
		CardCount: 3,
	}
	PENG = GroupType{
		Public:    true,
		CardCount: 3,
	}
	MING_GANG = GroupType{
		Public:    true,
		CardCount: 4,
	}
	AN_GANG = GroupType{
		Public:    true,
		CardCount: 4,
	}

	KE = GroupType{
		Public:    false,
		CardCount: 3,
	}
	SHUN = GroupType{
		Public:    false,
		CardCount: 3,
	}
	DUI = GroupType{
		Public:    false,
		CardCount: 3,
	}

	ZU_HE_LONG = GroupType{
		Public:    false,
		CardCount: 9,
	}

	QUAN_BU_KAO = GroupType{
		Public:    false,
		CardCount: 14,
	}

	QI_XING_BU_KAO = GroupType{
		Public:    false,
		CardCount: 14,
	}

	SHI_SAN_YAO = GroupType{
		Public:    false,
		CardCount: 14,
	}
)
