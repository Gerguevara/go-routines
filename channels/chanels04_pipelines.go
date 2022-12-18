package channels

import "fmt"

func LoadChannels() {
	generator := make(chan int)
	doubles := make(chan int)
	//4 ejecuta funciones para estos channels
	go Generator(generator)
	go Double(generator, doubles)
	//5 no cre una go routing, y como el cana doubles ya fue cerrado puede ejecutarlo normalmente
	Print(doubles)
}

// 1 esta funcion recibe un un canal y lo llena contando de 1 a 10 agregando el contador contador
func Generator(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	// cierra el canal para que ya no recibira mas datos, y desbloquera proceso congelados que esperarn respuesta de estecanal
	close(c)
}

// 2 esta function recibe  2 channels , el primero solo de lectura (se designa asi con <-chan) para recorrer el for
// el segundo solo de escritura (se designa asi con chan<-)  para cargarlo con la data (que el valor del canal solo lectura x 2)
func Double(chanel_In <-chan int, chanel_out chan<- int) {
	for v := range chanel_In {
		chanel_out <- 2 * v
		// si este cana no estuviera limitado a solo lectura he intentaramos agregar mas datos, daria error
		//chanel_In <-1
	}
	// esta funcion cierra el canal para que ya no reciba mas datos y desbloquera proceso congelados que esperarn respuesta de estecanal
	close(chanel_out)
}

// 3 esta funcion solo imprime el canal
func Print(c chan int) {
	for value := range c {
		fmt.Printf("#%d\n", value)
	}
}
