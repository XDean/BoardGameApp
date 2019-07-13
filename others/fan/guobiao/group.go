package guobiao

type (
	Group struct {
		Type  GroupType
		Cards Cards
	}
)

func (g Group) isJiang() bool {
	return g.Type == GT_JIANG
}

func (g Group) isShunZi() bool {
	return g.Type == GT_SHUN || g.Type == GT_CHI
}

func (g Group) isKeZi() bool {
	return g.Type == GT_KE || g.Type == GT_PENG || g.Type == GT_MING_GANG || g.Type == GT_AN_GANG
}

func (g Group) isGang() bool {
	return g.Type == GT_MING_GANG || g.Type == GT_AN_GANG
}

func (g Group) isAnKe() bool {
	return g.Type == GT_KE || g.Type == GT_AN_GANG
}

func (g Group) isChiPengMing() bool {
	return g.Type == GT_CHI || g.Type == GT_PENG || g.Type == GT_MING_GANG
}
