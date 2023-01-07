package main

import "fmt"

func main() {
	// cint := make(chan int) // unbuffered channel

	cint := make(chan int, 2) // buffered channel
	cint <- 1
	cint <- 2
	fmt.Println(<-cint)
	fmt.Println(<-cint)

	// Unbuffered channel: Espera una función o una rutina para recibir el mensaje, es bloqueada por ser llamada en la misma función

	// Buffered channel: Se puede llamar de manera inmediata, en el siguiente ejemplo 2 es el numero de canales que pueden ser usados
}
