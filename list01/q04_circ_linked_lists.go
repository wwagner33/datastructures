/*
******************************************************************************************

	@header
	Universidade Federal do Ceará – UFC
	Centro de Ciências – CC
	Mestrado e Doutorado em Ciências da Computação - MDCC
	Estruturas de Dados
	Prof. Dr. José Maria Filho

	@description Programa que implementa a Questão 04 da lista de exercícios EDA_List1.pdf
	@author Wellington Wagner F. Sarmento
	@version 1.0
	@date 08/09/2022

******************************************************************************************
*/
package main

import (
	"fmt"
	"runtime/debug"
)

/* estruturas  */

type Lista struct {
	info    int
	proximo *Lista
}

var head *Lista = nil
var tail *Lista = nil

/* funcao de suporte */

func libera(nohAlvo *Lista) { //libera noh da lista
	nohAlvo = nil        // o ponteiro sera eliminado pelo GC
	debug.FreeOSMemory() //acelera a devolucao da memoria para o OS (nao precisa esperar 5 minutos)
}

/* funcoes de manipulacao de listas */

// item 1. Criar uma lista vazia;
func lst_cria() *Lista {
	return nil
}

// item 2. Inserir elemento no início;
func lst_insere(lista *Lista, val int) *Lista {
	var node *Lista
	node = new(Lista)
	node.info = val
	node.proximo = lista
	if tail == nil {
		tail = node
	}
	head = node
	tail.proximo = node
	return node
}

// item 3. Imprimir os valores armazenados na lista;
func lst_imprime(lista *Lista) {

	if estaVazia(lista) != 1 {
		p := lista
		for {
			fmt.Println(p.info, p.proximo.info)
			if p.proximo == head {
				break
			}
			p = p.proximo
		}
	} else {
		fmt.Println("Lista está vazia!")
	}
}

// item 4. Imprimir os valores armazenados na lista usando recursão;
func lst_imprime_rec(lista *Lista) {

	if estaVazia(lista) != 1 {
		fmt.Println(lista.info, lista.proximo.info)
		if lista.proximo != head {
			lst_imprime_rec(lista.proximo)
		}
	}
}

// item 5. Imprimir os valores armazenados na lista em ordem reversa (da cauda para a cabeça da lista);
func lst_imprime_inv(lista *Lista) {

	if estaVazia(lista) != 1 {
		if lista.proximo != head {
			lst_imprime_inv(lista.proximo)
		}
		fmt.Println(lista.info, lista.proximo.info)
	} else {
		fmt.Println("Lista está vazia!")
	}
}

// item 6. Verificar se a lista está vazia (retorna 1 se vazia ou 0 se não vazia);
func estaVazia(lista *Lista) int {
	if lista == nil {
		return 1
	} else {
		return 0
	}

}

// item 7. Recuperar/Buscar um determinado elemento da lista;
func lst_busca(lista *Lista, val int) *Lista {
	var p *Lista = lista
	for {
		if p.info == val {
			return p
		}
		if p.proximo.info == head.info {
			break
		}
		p = p.proximo
	}
	return nil
}

// item 8. Remover um determinado elemento da lista;
func lst_retira(lista *Lista, val int) *Lista {
	var ant *Lista = nil /* ponteiro para elemento anterior */
	var p *Lista = lista /* ponteiro para percorrer a lista */
	var head int = lista.info

	for {
		if p.proximo.info == head || p.info == val {
			break
		}
		ant = p
		p = p.proximo
	}
	/* verifica se achou elemento */
	if p == nil {
		return lista /* não achou: ret lista original */
	}
	/* achou: retira */

	ant.proximo = p.proximo

	libera(p) /* libera espaço ocupado pelo elemento */
	return lista
}

// item 9. Remover um determinado elemento da lista usando recursão;
func lst_retira_rec(lista *Lista, val int) *Lista {
	var p *Lista = lista //ponteiro para percorrer a lista
	var temp *Lista

	if estaVazia(p) != 1 {
		if lista.info == val {
			temp = lista
			lista = lista.proximo
			libera(temp)

		} else {
			lista.proximo = lst_retira_rec(lista.proximo, val)
		}

	} else {
		fmt.Println("\nElemento não encontr  ado! Segue a lista dos elementos da lista:")
	}
	return lista
}

// item 10. Liberar a lista;

func lst_libera(lista *Lista) *Lista {
	lista = nil          // a lista apontara para NULL
	debug.FreeOSMemory() //  a memoria voltara para o SO
	return lista
}

func main() {
	var lista = lst_cria()

	// testes das funcoes

	//teste 01: verifica se lesta esta vazia e tenta imprimir
	fmt.Println("\n // teste 01: verifica se lesta esta vazia e tenta imprimir")
	fmt.Println(estaVazia(lista) == 1)
	lst_imprime(lista)

	//	teste 02: insere elementos e imprime a lista
	fmt.Println("\n//	teste 02: insere elementos e imprime a lista")
	lista = lst_insere(lista, 5)
	lista = lst_insere(lista, 51)
	lista = lst_insere(lista, 3)
	lista = lst_insere(lista, 17)
	lista = lst_insere(lista, 13)
	lst_imprime(lista)

	// //	teste 03: imprime elementos da lista sem recursao
	// fmt.Println("\n//	teste 03: imprime elementos da lista sem recursao")
	// lst_imprime(lista)

	// //	teste 04: imprime elementos da lista usando recursao
	// fmt.Println("\n//	teste 04: imprime elementos da lista usando recursao")
	// lst_imprime_rec(lista)

	// //	teste 05: imprime elementos da lista no sentido inverso
	// fmt.Println("\n//	teste 05: imprime elementos da lista no sentido inverso")
	// lst_imprime_inv(lista)

	// //	teste 06: verifica se a lista esta vazia
	// fmt.Println("\n //	teste 06: verifica se a lista esta vazia")
	// fmt.Println(estaVazia(lista) == 1)

	// // teste 07: busca por um elemento que nao exista na lista
	// fmt.Println("\n// teste 07: busca por um elemento que nao exista na lista")
	// var p *Lista
	// p = lst_busca(lista, 7)
	// if p == nil {
	// 	fmt.Println("O elemento não foi encontrado!")
	// } else {
	// 	fmt.Println("O elemento foi encontrado! Teste falhouuuuu!")
	// }

	// // teste 08: busca por um elemento que exista na lista no inicio, meio e fim
	// fmt.Println("\n// teste 08: busca por um elemento que exista na lista no inicio, meio e fim")

	// p = lst_busca(lista, 3) //inicio
	// if p != nil {
	// 	fmt.Println("O elemento do inicio foi encontrado! O valor dele é: ", p.info)
	// } else {
	// 	fmt.Println("O elemento não foi encontrado! Teste falhouuuuu!")
	// }

	// p = lst_busca(lista, 17) //meio
	// if p != nil {
	// 	fmt.Println("O elemento do meio foi encontrado! O valor dele é: ", p.info)
	// } else {
	// 	fmt.Println("O elemento não foi encontrado! Teste falhouuuuu!")
	// }

	// p = lst_busca(lista, 51) //fim
	// if p != nil {
	// 	fmt.Println("O elemento do fim foi encontrado! O valor dele é: ", p.info)
	// } else {
	// 	fmt.Println("O elemento não foi encontrado! Teste falhouuuuu!")
	// }

	// // teste 09: retira um elemento da lista e imprime
	// fmt.Println("\n// teste 09: retira um elemento da lista e imprime")
	// lst_retira(lista, 5)
	// lst_imprime(lista)

	// //	teste 10: retira elemento com recursao  e imprime a lista
	// fmt.Println("\n//	teste 10: remove elemento com recursao")
	// lst_retira_rec(lista, 3)
	// lst_imprime(lista)

	// teste 11: libera a lista e verifica se esta vazia
	// fmt.Println("\n// teste 10: libera a lista e verifica se esta vazia")
	// lista = lst_libera(lista)
	// fmt.Println(estaVazia(lista) == 1)
	// lst_imprime(lista)

}
