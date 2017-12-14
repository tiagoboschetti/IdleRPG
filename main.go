package main

import (
	"fmt"
	"reflect"
)

type ClasseProvider interface {
	GetClasse() *Classe
	Desviar() bool
	Defender() bool
	Castar() bool
}

func (c *Classe) GetClasse() *Classe {
	return c
}

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

func Atacar(ppAtacante, ppDefensor ClasseProvider) *Classe {

	if reflect.TypeOf(ppAtacante).Name() == "Mago" && ppAtacante.Castar() {
		// procou o double damage
		aAtaque := ppAtacante.GetClasse().ataque
		aDanoCausadoAoDefensor := aAtaque * 2
		fmt.Println(">>>>DOUBLE DAMAGE<<<<")
	}

	if reflect.TypeOf(ppDefensor).Name() == "Arqueiro" && ppDefensor.Desviar() {
		// procou o evasion
		aDanoCausadoAoDefensor := 0
		fmt.Println(">>>>EVASION<<<<")
	}

	if reflect.TypeOf(ppDefensor).Name() == "Guerreiro" && ppDefensor.Defender() {
		// procou o ADO
		aDanoCausadoAoDefensor := ppAtacante.GetClasse().ataque
		aDanoCausadoAoAtacante := ppDefensor.GetClasse().ataque
		fmt.Println(">>>>ATAQUE DE OPORTUNIDADE<<<<")
	}
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

func (a *Arqueiro) Desviar() bool {
	// toda vez que receber ataque terá uma chance de desviar do ataque
	// RETORNA TRUE CASO O EVADE TENHA SIDO BEM SUCEDIDO
	return true

}

func (g *Guerreiro) Defender() bool {
	// chance de ADO
	// RETORNA TRUE CASO O ATAQUE DE OPORTUNIDADE TENHA PROCADO
	return true
}

func (m *Mago) Castar() bool {
	// Tem chance de causar double damage
	// RETORNAR TRUE CASO PROQUE O DOUBLE DAMAGE
	return true
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
