package main

import (
	"errors"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

const raizID = 0

var arvore_base = &Node{}

type Node struct {
	ID       int
	Children []*Node
}

func main(){
	app := fiber.New()

	app.Static("/", "../static")

	app.Get("/submit", fazer_arvore)

	app.Listen(":3000")
}

//função que percorre a árvore e retorna o nódulo que tem o ID pai
func adicionar(raiz *Node, pai int, filho *Node) error {
	if raiz == nil {
		*raiz = *filho
		return 	nil	
	} 


    if raiz.ID == pai {
        raiz.Children = append(raiz.Children, filho) // Adicionando o novo nó como filho
		return nil
    }

    // Percorra recursivamente os filhos
    for i := range raiz.Children {
		adicionar(raiz.Children[i], pai, filho)	
		return nil	
    }

	return errors.New("O pai não está presente na árvore")
	
}

func fazer_arvore(c *fiber.Ctx) error{
	id := c.FormValue("ID")

	int_id, err := strconv.Atoi(id)

	if err != nil {
		return c.Redirect("/erro.html")
	}

	pai := c.FormValue("pai")

	int_pai, err := strconv.Atoi(pai)

	if err != nil {
		return c.Redirect("/erro.html")
	}

	if int_pai>int_id {
		return c.Redirect("/erro.html")
	}

	filho := &Node{ID: int_id, Children: []*Node{}}

	adicionar(arvore_base, int_pai, filho)

	return c.Redirect("/sucesso.html")
}