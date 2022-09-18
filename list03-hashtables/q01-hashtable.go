/*
******************************************************************************************

	@header
	Universidade Federal do Ceará – UFC
	Centro de Ciências – CC
	Mestrado e Doutorado em Ciências da Computação - MDCC
	Estruturas de Dados
	Prof. Dr. José Maria Filho

	@description Programa que implementa a Questão 03 da lista de exercícios EDA_List3.pdf
	@author Wellington Wagner F. Sarmento
	@version 1.0
	@date 18/09/2022

******************************************************************************************
*/
package main

import (
	"fmt"
	"runtime/debug"
)

/* estruturas de lista e tabela de dispercao (hash table) */
type Lista struct {
	valor   string
	proximo *Lista
}
type HashTable struct {
	chave   int
	valor   *Lista
	proximo *HashTable
}

/* funcao de suporte */

func libera(nohAlvo *Lista) { //libera noh da lista
	nohAlvo = nil        // o ponteiro sera eliminado pelo GC
	debug.FreeOSMemory() //acelera a devolucao da memoria para o OS (nao precisa esperar 5 minutos)
}

/* funcoes de manipulacao da hash table*/

// returns a hash between 0 and m-1
func hash(key int) int {
	return (key & 0x7fffffff) % m
}

func has_cria() *HashTable {
	return nil
}

func has_insere(hash *HashTable, val string) *HashTable {
	var lista *Lista
	lista = new(Lista)
	lista.valor = val
	lista.proximo = nil

	var item *HashTable
	item = new(HashTable)
	item.chave = 0 // *** falta
	item.valor = list
	item.proximo = hash

	return item
}

func has_busca_lista(list *Lista, val string) *Lista {
	var p *Lista
	for p = list; p != nil; p = p.next {
		if p.valor == val {
			return p
		}
	}
	return nil
}

func has_busca(hash *HashTable, val string) *HashTable {
	var p *HashTable
	for p = hash; p != nil; p = p.next {
		for i := p.value; i != nil; i = i.next {
			if i.value == val {
				return p
			}
		}
	}
	return nil
}

func has_retira(lista *HashTable, val string) *HashTable {
	var anterior *HashTable = nil /* ponteiro para elemento anterior */
	var p *HashTable = lista      /* ponteiro para percorrer a tabela */
	var item *Lista

	/* procura elem na lista, guardando anterior  */

	for {
		item = has_busca_lista(p.valor, val)
		if p == nil || item.valor == val {
			break
		}
		anterior = p
		p = p.proximo
	}
	/* verifica se achou elemento */
	if p == nil {
		return lista /* não achou: ret lista original */
	}
	/* achou: retira */
	if anterior == nil { /* retira elemento do inicio */
		lista = p.proximo
	} else { /* retira elemento do meio da lista */
		anterior.proximo = p.proximo
	}
	p = nil /* libera espaço ocupado pelo elemento */
	return lista
}

func estaVazia(lista *HashTable) bool {
	if lista == nil {
		return true
	} else {
		return false
	}

}

func has_libera(lista **HashTable) **HashTable {
	lista = nil          // a lista apontara para NULL
	debug.FreeOSMemory() //  a memoria voltara para o SO
	return lista
}

func main() {

	var n float64

	//fmt.Print("Informe um num inteiro:")
	//fmt.Scanln("%d\n", &n)

	n = 12
	// m := int(math.Round(n / 2))

	var hash = has_cria()
	fmt.Println(estaVazia(hash))

	hash = has_insere(hash, "23")
	hash = has_insere(hash, "oi")
	hash = has_insere(hash, "-")
	hash = has_insere(hash, "9")
	hash = has_insere(hash, "abc")

	print(hash)
	fmt.Println(estaVazia(hash))

	node := has_busca(hash, "9")
	if node != nil {
		fmt.Println(node.valor, node.proximo)
	}

	hash = has_libera(hash)
	fmt.Println(estaVazia(hash))
}
