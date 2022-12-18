package worker_pool

import "fmt"

func DoWorker() {

	tasks := []int{2, 3, 4, 5, 7, 10, 12, 35, 37, 40, 41, 42}

	//limite de proceso concurrentes al mismo tiempo
	nWorkers := 5
	//declara un canal con la capacidad maxima de el array de task
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	// va a generar 5 workers, y estos quedan a la espera de la emision del canal de solo lectura
	//para dar paso a la logica y emiter el canal de escritura
	for w := 1; w <= nWorkers; w++ {
		go Worker(w, jobs, results)
	}

	// va a mete en el canal de jobs los task cada inicia una task los workers declarados anteriormente se ponen a trabajar
	// l;uego de iterar cierra el canl y desbloque los proceso subsecuentes
	for _, t := range tasks {
		jobs <- t
	}

	close(jobs)

	// va a  tomar y lerr los valosres de task poner finalizar el canal
	for a := 1; a <= len(tasks); a++ {
		<-results
	}
}

// recibe el index, un chanel solo de lectura llamado jobs y un chanel de escritura llamado results
// no empieza a trabajar hasta que el canal emite
func Worker(id int, jobs <-chan int, results chan<- int) {
	//itera los jobs conforme se van ocupando estos van tomando las tareas
	//1 toma y se bloquea queda ejecutando el fibonacci, mientras ehjecuta la logica  el siguiente toma otra tarea
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		fib := Fibonnaci(j)
		//mete en el canal results lo que el fibonacci va generando
		results <- fib
		fmt.Println("worker", id, "finished job", j, "result", fib)
	}
}

// esta funcion carlcula el fibonacci de el numero dado con recursividad
func Fibonnaci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonnaci(n-1) + Fibonnaci(n-2)
}

//este codigo basicamente hace que a cada uno de los elementos de task se le calcule su valor de fibonacci
// al mismo tiempo
