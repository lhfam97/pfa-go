package main

import "time"

// func contador(tipo string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(tipo, i)
// 		time.Sleep(time.Second)
// 	}
// }

func worker(workerId int, msg chan int) {

	for res := range msg {
		println("worker", workerId, "recebeu", res)
		time.Sleep(time.Second)
	}
}

// Main eh um thread
func main() {
	//o forever cria um canal que nunca vai receber um item, logo se colocar ele no final do codigo, o codigo nunca vai encerrar
	forever := make(chan bool)
	canal := make(chan int)

	for i := 0; i < 5; i++ {
		go worker(i, canal)
	}

	for i := 0; i < 10; i++ {
		canal <- i
	}
	<-forever
}
