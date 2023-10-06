package tree
import "fmt"

const raizID = 0

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
    return append(s[:i], s[i+1:]...)
}

//função que percorre a árvore e retorna o nódulo que tem o ID pai
func adicionar(raiz *Node, pai int, filho *Node) {
    if raiz.ID == pai {
        raiz.Children = append(raiz.Children, filho) // Adicione o novo nó como filho
    }

    // Percorra recursivamente os filhos
    for i := range raiz.Children {
		adicionar(raiz.Children[i], pai, filho)
    }


}

func partition(arr []Record, low, high int, modo int) ([]Record, int) {
	pivot := arr[high]
	i := low
	if modo == 0 {
		for j := low; j < high; j++ {
			if arr[j].Parent < pivot.Parent {
				arr[i], arr[j] = arr[j], arr[i]
				i++
			}
		}
	}
	if modo == 1 {
		for j := low; j < high; j++ {
			if arr[j].ID < pivot.ID {
				arr[i], arr[j] = arr[j], arr[i]
				i++
			}
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []Record, low, high int, modo int) []Record {
	if low < high {
		var p int
		arr, p = partition(arr, low, high, modo)
		arr = quickSort(arr, low, p-1, modo)
		arr = quickSort(arr, p+1, high, modo)
	}
	return arr
}

func quickSortStart(arr []Record, modo int) []Record {
	return quickSort(arr, 0, len(arr)-1, modo)
}


func Build(records []Record) (*Node, error) {

	raiz := &Node{}//ponteiro pra Node

	//se não tiver nenhum record, retorna nulo
	if len(records) == 0 {
		return nil, nil
	} 
	fila := make([]int, len(records))
    
    for j := range records{
        if records[j].ID < raizID || records[j].ID >= len(records){
            return nil, fmt.Errorf("Nó com ID menor que a raiz e maior que o máximo de nós")
        }
    	if (records[j].ID < records[j].Parent) || (records[j].ID == raizID && records[j].Parent != raizID){
            return nil, fmt.Errorf("Nós que não tem possibilidade de existirem")
        }
    	if (records[j].ID != raizID && records[j].ID == records[j].Parent){
            return nil, fmt.Errorf("Nó que é filho de si mesmo")
        }
    	fila[records[j].ID] = j
    }
	for j := range fila{
        if records[fila[j]].ID != j{
            return nil, fmt.Errorf("Nós não continuos")
        }
    } 

	//Ao enncontrar o elemento com ID = 0, ele atribui isso à raiz e remove da array
	for i := 0; i < len(records); i++ {
		if records[i].ID == 0 {
			raiz.ID = 0
			records=remover(records,i)
		}
	}

	//ordenando os recordes em ordem crescente pelo pai
	records=quickSortStart(records,0)

	records=quickSortStart(records,1)
    

	for i := 0; i < len(records); i++ {	
		aux2 := &Node{ID: records[i].ID, Children: []*Node{}}

		adicionar(raiz,records[i].Parent,aux2)

		}


	return raiz, nil
}
