package channels

import "fmt"

//los buffer en los canales son basicamente cuando manejasn mas de 1 a la vez, si ne la declaracion de un canal no
//se le pone el segundo parametro  inmediatamente el canal emitar pero eos causara un  error, ya que lo interpreta como buffer 0

// 1 espacio de canal no se considera liberado hasta que alguien lee su valor o arbitratiamente se libera el espacio con 	<-c

func ChannelWithoutBuffer() {
	c := make(chan int) // Buffered
	// se bloquea en este punto esperon que otra funcion lea el mensaje pero eso jamas pasa
	c <- 1
	//no llega a  este punto dado que quedo esperando,
	fmt.Println(c)
	//por lo tanto siewmpre debe haber alguien esperando y escuchando el canal
}

// cuando un canal no tiene buffer y nadie espera la funcion esto se bloquea a diferencia de con buffer que bien puede ejecutarse
// sin problemas porque el canal esta abierto
func ChannelWithBuffer() {
	// c := make(chan int) // Unbuffered
	c := make(chan int, 2) // Buffered

	c <- 1
	c <- 2

	fmt.Println(<-c)
	fmt.Println(<-c)
}

// con canal de buffer y la ejecucuin inmediata el prblema basicamente es que se desborda la capacidad del canal
//porque es 0, entonces para vaciar el canal es necesario que alguien escuche esa emicion para poder continuar
//y quue no se desborde, esto mismo pasaria en un canal con buffer pero que le pasemos mas valores y desbordemos su capacidad
