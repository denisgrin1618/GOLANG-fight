package fight

type Fighter interface {
	Hit() int
	Dodge() bool
	GetDamage(int)
	CanFight() bool
}

func Round(f1 Fighter, f2 Fighter) {
	if !f2.Dodge() {
		f2.GetDamage(f1.Hit())
	}
}
