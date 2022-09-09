package main

import (
	"fmt"
	"runtime/debug"
)

/* estrutura de lista */

type Lista struct {
	info    int
	proximo *Lista
}

/* funcoes de suporte */

func estaVazia(lista *Lista) bool {
	return (lista == nil)
}

func libera(nohAlvo *Lista) bool {
	nohAlvo = nil        // o ponteiro sera eliminado pelo GC
	debug.FreeOSMemory() //acelera a devolucao da memoria para o OS (nao precisa esperar 5 minutos)
}

/* funcoes de manipulacao de listas */

func lst_cria() *Lista {
	return nil
}

func lst_insere(lista *Lista, val int) *Lista {
	var nvLista = new(Lista)
	nvLista.info = val
	nvLista.proximo = lista
	return nvLista
}

func lst_imprime(lista *Lista) {

	if !estaVazia(lista) {
		for p := lista; p != nil; p = p.proximo {
			fmt.Println(p.info)
		}
	} else {
		fmt.Println("Lista está vazia!")
	}
}

func lst_retira(lista *Lista, val int) *Lista {
	var ant *Lista = nil //ponteiro para o elemento anterior
	var p *Lista = lista //ponteiro para percorrer a lista

	// procura elemento na lista, guardando o anterior
	for {
		if p != nil && p.info != val {
			ant = p
			p = p.proximo
		} else {
			break
		}
	}

	// verifica se achou elemento
	if estaVazia(p) {
		return lista //nao achou: retorna a lista original
	} else if ant == nil { //achou: retira elemento do inicio da lista
		lista = p.proximo
	} else { //retira elemento do meio da lista
		ant.proximo = p.proximo
	}
	libera(p)    //libera memoria para o SO
	return lista //retorna lista com item retirado
}

func lst_busca(lista *Lista, val int) *Lista {
	var p *Lista
	for p = lista; p != nil; p = p.proximo {
		if p.info == val {
			return p //retorna o elemento, caso o ache
		}
	}
	return nil //retorna NULL se nao achar o elemento
}

func main() {
	var lista = lst_cria()

	// testes das funcoes

	//teste 01: verifica se lesta esta vazia e tenta imprimir
	fmt.Println(estaVazia(lista))
	lst_imprime(lista)

	//	teste 02: insere elementos
	lista = lst_insere(lista, 7)
	lista = lst_insere(lista, 17)
	lista = lst_insere(lista, 77)
	lista = lst_insere(lista, 717)
	lista = lst_insere(lista, 737)

	//	teste 03: verifica se a lista continua vazia e imprime elementos da lista
	fmt.Println(estaVazia(lista))
	lst_imprime(lista)

	// teste 04: retira um elemento da lista e imprime
	lst_retira(lista, 77)
	lst_imprime(lista)

	var p *Lista

	// teste 04: busca por um elemento que nao exista na lista
	p = lst_busca(lista, 5)
	if p != nil {
		fmt.Println("O elemento não foi encontrado!")
	} else {
		fmt.Println("O elemento foi encontrado! Teste falhouuuuu!")
	}

	// teste 5: busca por um elemento que exista na lista no inicio, meio e fim
	p = lst_busca(lista, 7) //inicio
	if p != nil {
		fmt.Println("O elemento do inicio foi encontrado! O valor dele é: %d", p.info)
	} else {
		fmt.Println("O elemento não foi encontrado! Teste falhouuuuu!")
	}

	p = lst_busca(lista, 77) //meio
	if p != nil {
		fmt.Println("O elemento do meio foi encontrado! O valor dele é: %d", p.info)
	} else {
		fmt.Println("O elemento não foi encontrado! Teste falhouuuuu!")
	}

	p = lst_busca(lista, 737) //fim
	if p != nil {
		fmt.Println("O elemento do fim foi encontrado! O valor dele é: %d", p.info)
	} else {
		fmt.Println("O elemento não foi encontrado! Teste falhouuuuu!")
	}

}
