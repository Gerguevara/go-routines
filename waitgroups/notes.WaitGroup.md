Un WaitGroup es un contador de esperas basicamente cada que se agrega una goRoutin se suma uno al contador don el metodo wg.Add)(1) o
las que se requieran, para disminuir elk contador se utiliza wg.Done() el cual no recibe parametros y solo reduce en un el contador
el metodo wg.Wait pone en espera la aplicacion hasta que se complete la espera del wait group