/* Hey-Ho-Let-s-Go-
Meus primeiros passos em Go
*/
package main

import "fmt"

type Batata struct {
     tipo string
     tempoColheita int
     estragada bool
 }

func (batata *Batata) getTipo () string{
     return batata.tipo
}

func (batata *Batata) setTempoColheita (tempo int) {
     batata.tempoColheita = tempo
}

func main() {
     ruffles := Batata{"chips",10,false}
     fmt.Println("Sou uma batata do tipo",ruffles.getTipo())

     ruffles.setTempoColheita(20)
     fmt.Println("Levei",ruffles.tempoColheita,"dias ate ser colhido!")
}

