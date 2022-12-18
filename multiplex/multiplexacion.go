package multiplex

//dado que el codigo se ejecuta por orden de linea y no por orden de llgada se ocupa esta tecnica , para forzar el ver
//el primero que llego no el orden de ejcucion sequencial, en este ejemplo c1, d1 toman el domble de tiempo pero dao que fueron
//declarados antes este siempre seria el primero en mostrar la salida

//basicamente este programa en doSomething hace una espera de la ejecucion del programa, y con select
//fuerza a mostrase con el orden de ejecucion
import (
	"fmt"
	"time"
)

func LunchMultiplexation() {
	c1 := make(chan int)
	c2 := make(chan int)

	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go doSomething(d1, c1, 1)
	go doSomething(d2, c2, 2)

	for i := 0; i < 2; i++ {
		// un select es igual a un switch, especifico para canales, nos permite saber que canal ha sido activado
		select {
		case channelMsg1 := <-c1:
			fmt.Println(channelMsg1)
		case channelMsg2 := <-c2:
			fmt.Println(channelMsg2)
		}
	}
}

func doSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
