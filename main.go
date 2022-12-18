package main

// paso 1: primero importar, nota los paquetes en go siempre deben ser en minuscula y llamarse como su carpeta contendora
import mp "Buffer-Class-Go/multiplex"

// import wp "Buffer-Class-Go/worker_pool"
// ch "Buffer-Class-Go/channels"
// wg "Buffer-Class-Go/waitgroups"

func main() {
	// Paso 2: luego buscar por la funcion, recordar que debe empezar con mayuscula que sea exportable

	//ch.ChannelWithoutBuffer()
	//ch.ChannelWithBuffer()
	// wg.UsingWaitGroups()
	// ch.RunChannel()
	// ch.RunSemaphore()
	// ch.LoadChannels()
	mp.LunchMultiplexation()
}

//paso 3:  ejecutar go run main.go
