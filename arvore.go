package main

import "fmt"

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

//função que remove um elemento de um vetor desordenadamente
func remover(s []Record, i int) []Record {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

//função que percorre a árvore e retorna o nódulo que tem o ID pai
func percorrer(raiz *Node, pai int) *Node {
	if raiz.ID == pai {
		return raiz
	}

	for i := 0; i < len(raiz.Children); i++ {
		aux := percorrer(raiz.Children[i], pai)
		if aux != nil {
			return aux
		}
	}

	return nil
}


func Build(records []Record) (*Node, error) {

	raiz := &Node{}//ponteiro pra Node

	//se não tiver nenhum record, retorna nulo
	if len(records) == 0 {
		return nil, nil
	}

	//Ao enncontrar o elemento com ID = 0, ele atribui isso à raiz e remove da array
	for i := 0; i < len(records); i++ {
		if records[i].ID == 0 {
			raiz.ID = 0
			remover(records,i)
		}
	}

	//falta utilizar a função que acha e retorna o pai (implementada acima) para construir a árvore

}

//perguntar sobre como percorrer árvore em go para o LG