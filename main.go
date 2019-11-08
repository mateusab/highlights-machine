package main

import "fmt"
import "os"

func main() {
	fmt.Printf("hello, world\n")

	//pega os argumentos passados como parametro
    args := os.Args[1:]

	//printa os argumentos
    fmt.Println("demo: ", args)
}