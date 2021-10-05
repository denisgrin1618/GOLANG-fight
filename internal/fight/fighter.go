package fight

type Fighter interface {
	Hit() int
	Dodge() bool
	GetDamage(int)
	CanFight() bool
}

func Round(b Fighter, k Fighter) {
	if !k.Dodge() {
		k.GetDamage( b.Hit() )
	}

}




