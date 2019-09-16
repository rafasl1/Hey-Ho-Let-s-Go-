package main

import "fmt"

const MAX = 100

type registro struct {
    chave int
    prox int
}
type lista struct {
    A [MAX]registro
    dispo int
    ini int
}

func (l* lista) inicializa (){
    l.ini = -1
    l.dispo = -1
    for i := 0; i < MAX; i++ {
        l.A[i].prox = i+1 
    }
}

func main() {
    var l lista
    l.inicializa()
    fmt.Println("A lista tem inicio em",l.ini,"e o numero de elementos disponiveis eh",l.dispo+1)
}
