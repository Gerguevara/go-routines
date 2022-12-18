package waitgroups

import (
	"fmt"
	"sync"
	"time"
)

func UsingWaitGroups() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		//se pasa el wg como referencia
		go doSomething(i, &wg)
	}
	wg.Wait()
}

// parsea la referencia de wg a su valor
func doSomething(i int, wg *sync.WaitGroup) {
	// la linea que se ejecutar al final, cada que esta funcion finaliza reduce el wg en 1
	defer wg.Done()
	fmt.Printf("Started %d\n", i)
	//simular un proceso largo
	time.Sleep(time.Duration(2) * time.Second)

	fmt.Printf("Done: %d\n", i)
}
