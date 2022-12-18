package channels

// Los semaforos son una forma de limitar en go los procesos concurrentes, no dejando espacio para ejecucion de otro, hasta
//que la memoria se libere  para vitar colapsos
import (
	"fmt"
	"sync"
	"time"
)

func RunSemaphore() {
	c := make(chan int, 2) // creates a buffered channel with a capacity of two
	var wg sync.WaitGroup  // creates wait group

	for i := 0; i < 10; i++ {
		//  emite el valor de 1 a el chanel, pero recordar que el chanel esta limitado a 2 por lo tanto no recibe el valor
		// hasta que no se libera un espacio, lo cual pasa en la linea 36 con <-c, lo que obliga una espera hasta entonces,
		//pero ya que esto ocurre en 2 thrads diferentes puede ejecutar hasta 2 al mismo tiempo
		c <- 1
		// agrega 1 al waitgroup
		wg.Add(1)

		go doSomething(i, &wg, c)
	}

	//pone en pausa el programa hasta que se el contador del wg llega a 0
	// lo que quiere decir que el hilo principal de go no se detendra hasta que el contador llegue a 0
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	//resta 1 a wailgroup pero cuando termina la ejecucion de esta funcion
	defer wg.Done()
	fmt.Printf("Id: %d -> started...\n", i)
	//simula una espera de un proceso largo
	time.Sleep(time.Second * 4)

	fmt.Printf("Id: %d -> finished...\n", i)

	<-c // frees the space for new routines
}
