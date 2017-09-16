package main

import (
	"errors"
	"fmt"
)

func main() {
	pilha := Pilha{}
	fmt.Println("Pilha criada com tamanho ", pilha.Tamanho())
	fmt.Println("Vazia?", pilha.Vazia())

	pilha.Empilhar("Go")
	pilha.Empilhar(2009)
	pilha.Empilhar(3.14)
	pilha.Empilhar("Fim")
	fmt.Println("Tamanho após empilhar 4 valores: ", pilha.Tamanho())
	fmt.Println("Vazia?", pilha.Vazia())

	for !pilha.Vazia() {
		v, _ := pilha.Desempilhar()
		fmt.Println("Desempilhando ", v)
		fmt.Println("Tamanho: ", pilha.Tamanho())
		fmt.Println("Vazia?", pilha.Vazia())
	}

	_, err := pilha.Desempilhar()
	if err != nil {
		fmt.Println("Erro: ", err)
	}
}

// Pilha Representa uma pilha
type Pilha struct {
	valores []interface{}
}

// Vazia Verifica se a pilha está vazia
func (pilha Pilha) Vazia() bool {
	return pilha.Tamanho() == 0
}

// Tamanho Retorna o tamanho da pilha
func (pilha Pilha) Tamanho() int {
	return len(pilha.valores)
}

// Empilhar adiciona um elemento á pilha
func (pilha *Pilha) Empilhar(valor interface{}) {
	pilha.valores = append(pilha.valores, valor)
}

// Desempilhar Remove o último elemento da pilha
func (pilha *Pilha) Desempilhar() (interface{}, error) {
	if pilha.Vazia() {
		return nil, errors.New("pilha vazia")
	}
	valor := pilha.valores[pilha.Tamanho()-1]
	pilha.valores = pilha.valores[:pilha.Tamanho()-1]
	return valor, nil
}
