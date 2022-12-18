package channels

import (
	"fmt"
	"time"
)

func RunChannel() {
	fmt.Println("Star running")

	//1 crea y pasa el channel
	c := make(chan int)
	go doSomethingHere(c)

	//3 toma del channel el valor emitido, lo que da paso a la siguiente linea de codigo
	<-c
	//3.1 opcional, se puede tomar el valor emitido de un channel y utlizarlo
	// d := <-c
	// fmt.Println("Completed with value", d)
	fmt.Println("Completed")
}

func doSomethingHere(c chan int) {
	//simular una espera larga
	time.Sleep(time.Second * 3)
	fmt.Println("Done")
	//2 pone en el chanel el valor de 1
	c <- 10
}
