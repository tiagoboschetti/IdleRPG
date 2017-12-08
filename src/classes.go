package classes

type Classe interface {
	vida() int
	defesa() int
	ataque() int
}

type Arqueiro struct {
}

func (a *Arqueiro) Vida() int {
	return 10
}

func (a *Arqueiro) Defesa() int {
	return 4
}

func (a *Arqueiro) Ataque() int {
	return 5
}

func (a Arqueiro) Desviar() bool {
	// toda vez que receber ataque terá uma chance de desviar do ataque
	return true
}

type Guerreiro struct {
}

func (a *Guerreiro) Vida() int {
	return 13
}

func (a *Guerreiro) Defesa() int {
	return 6
}

func (a *Guerreiro) Ataque() int {
	return 3
}

func (g Guerreiro) Defender() bool {
	// chance de ADO
	return true
}

type Mago struct {
}

func (a *Mago) Vida() int {
	return 8
}

func (a *Mago) Defesa() int {
	return 2
}

func (a *Mago) Ataque() int {
	return 8
}

func (m Mago) Castar() bool {
	// Tem chance de causar double damage
	return true
}
