package main

import "fmt"

type Classe struct {
	vida   int
	defesa int
	ataque int
}

func (c *Classe) Vida() int {
	return c.vida
}

func (c *Classe) Defesa() int {
	return c.defesa
}

func (c *Classe) Ataque() int {
	return c.ataque
}

type Arqueiro struct {
	Classe
}

type Guerreiro struct {
	Classe
}

type Mago struct {
	Classe
}

func main() {
	a := Arqueiro{Classe: Classe{vida: 10, defesa: 3, ataque: 5}}
	fmt.Println("Arqueiro")
	fmt.Println("  > Vida: ", a.vida)
	fmt.Println("  > Defesa: ", a.defesa)
	fmt.Println("  > Ataque: ", a.ataque)

	g := Guerreiro{Classe: Classe{vida: 15, defesa: 6, ataque: 2}}
	fmt.Println("Guerreiro")
	fmt.Println("  > Vida: ", g.vida)
	fmt.Println("  > Defesa: ", g.defesa)
	fmt.Println("  > Ataque: ", g.ataque)

	m := Mago{Classe: Classe{vida: 8, defesa: 2, ataque: 8}}
	fmt.Println("Mago")
	fmt.Println("  > Vida: ", m.vida)
	fmt.Println("  > Defesa: ", m.defesa)
	fmt.Println("  > Ataque: ", m.ataque)
}
