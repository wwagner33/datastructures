/*
******************************************************************************************

	@header
	Universidade Federal do Ceará – UFC
	Centro de Ciências – CC
	Mestrado e Doutorado em Ciências da Computação - MDCC
	Estruturas de Dados
	Prof. Dr. José Maria Filho

	@description Programa que implementa a Questão 06 da lista de exercícios EDA_List1.pdf
	@author Wellington Wagner F. Sarmento
	@version 1.0
	@date 08/09/2022

******************************************************************************************
*/
package main

import (
	"fmt"
)

// incompleta

type Conta struct {
	tipoConta int // 1 corrente, 2 poupanca, 3 fidelidade
	numero    float32
	proximo   *Conta
}
type checaConta struct { //corrente
	numero float32
	saldo  float64
}
type SavingsConta struct { //poupanca
	numero float32
	saldo  float64
}
type LoyaltyConta struct { //fidelidade
	numero float32
	saldo  float64
	bonus  float64
}

func create() *Conta {
	return nil
}

func insereContas *Conta, tipoConta int, numero float32, saldo float64, bonus float64) *Conta {
	var novaConta *Conta
	novaConta = new(Conta)
	novaConta.tipoConta = tipoConta
	novaConta.numero = numero
	novaConta.proximo = Contas
	switch tipoConta {
	case 1:
		addchecaConta(numero, saldo)
	case 2:
		addSavingsConta(numero, saldo)
	case 3:
		addLoyaltyConta(numero, saldo, bonus)
	}
	return novaConta
}
func addchecaConta(numero float32, saldo float64) {
	var checa *checaConta
	checa = new(checaConta)
	checa.saldo = saldo
	checa.numero = numero
}
func addSavingsConta(numero float32, saldo float64) {
	var savings *SavingsConta
	savings = new(SavingsConta)
	savings.saldo = saldo
	savings.numero = numero
}
func addLoyaltyConta(numero float32, saldo float64, bonus float64) {
	var loyalty *LoyaltyConta
	loyalty = new(LoyaltyConta)
	loyalty.saldo = saldo
	loyalty.numero = numero
	loyalty.bonus = bonus
}

func remove(Conta *Conta, numero int) *Conta {
	var previous *Conta = nil /* ponteiro para elemento anterior */
	var p *Conta = Conta      /* ponteiro para percorrer a lista */

	/* procura elem na lista, guardando anterior  - while com for */
	for {
		if p == nil || int(p.numero) == numero {
			break
		}
		previous = p
		p = p.proximo
	}
	/* verifica se achou elemento */
	if p == nil {
		return Conta /* não achou: ret lista original */
	}
	/* achou: retira */
	if previous == nil { /* retira elemento do inicio */
		Conta = p.proximo
	} else { /* retira elemento do meio da lista */
		previous.proximo = p.proximo
	}
	p = nil /* libera espaço ocupado pelo elemento */
	return Conta
}

func recursive_remove(Conta *Conta, numero float32) *Conta {
	if !intToBool(empty(Conta)) {
		if Conta.numero == numero {
			Conta = Conta.proximo
		} else {
			Conta.proximo = recursive_remove(Conta.proximo, numero)
		}
	}
	return Conta
}

func empty(Conta *Conta) int {
	//return (Conta == nil)
	if Conta == nil {
		return 1
	}
	return 0
}

func search(Conta *Conta, numero float32) *Conta {
	var p *Conta
	for p = Conta; p != nil; p = p.proximo {
		if p.numero == numero {
			return p
		}
	}
	return nil
}

func print(Conta *Conta) {
	if !intToBool(empty(Conta)) {
		for p := Conta; p != nil; p = p.proximo {
			fmt.Println(p.numero, p.tipoConta)
		}
	}
}

func reverse_print(Conta *Conta) {
	if !intToBool(empty(Conta)) {
		reverse_print(Conta.proximo)
		fmt.Println(Conta.numero, Conta.tipoConta)
	}
}

func recursive_print(Conta *Conta) {
	if !intToBool(empty(Conta)) {
		fmt.Println(Conta.numero, Conta.tipoConta)
		if Conta.proximo != nil {
			recursive_print(Conta.proximo)
		}
	}
}

func intToBool(b int) bool {
	if b == 1 {
		return true
	}
	return false
}

// O GC do Go eh quem faz a liberação
func liberaConta *Conta) *Conta {
	return nil
}

func main() {
	var Conta = create()
	fmt.Println(intToBool(empty(Conta)))

	Conta = insereConta, 1, 235, 230, 0)
	Conta = insereConta, 3, 563, 230, 50)
	Conta = insereConta, 2, 123, 230, 0)
	Conta = insereConta, 2, 654, 230, 0)
	Conta = insereConta, 1, 205, 150, 0)
	Conta = insereConta, 3, 241, 80, 10)
	Conta = insereConta, 1, 794, 100, 0)

	print(Conta)
	recursive_print(Conta)
	reverse_print(Conta)
	fmt.Println(intToBool(empty(Conta)))

	var ac = search(Conta, 654)
	if ac != nil {
		fmt.Println(Conta.numero, Conta.tipoConta)
	}

	ac = remove(Conta, 205)
	print(Conta)
	ac = recursive_remove(Conta, 794)
	print(Conta)

	Conta = liberaConta) // O GC faz a liberacao
	fmt.Println(intToBool(empty(Conta)))
}
