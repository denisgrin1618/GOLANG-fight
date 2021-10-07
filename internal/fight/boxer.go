package fight

type Boxer struct {
	Sportsman
}

func (b *Boxer) GetDamage(d int) {
	b.Health -= d
}
