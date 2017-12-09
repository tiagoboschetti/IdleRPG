package classes

func (a *Arqueiro) Desviar() bool {
	// toda vez que receber ataque terá uma chance de desviar do ataque
	return true
}

func (g *Guerreiro) Defender() bool {
	// chance de ADO
	return true
}

func (m *Mago) Castar() bool {
	// Tem chance de causar double damage
	return true
}
