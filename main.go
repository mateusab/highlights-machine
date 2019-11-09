package main

import (
	"fmt"
	"os"

	dem "github.com/markus-wa/demoinfocs-golang"
	// events "github.com/markus-wa/demoinfocs-golang/events"
)

func main() {
	// pega os argumentos passados como parametro
	args := os.Args[1:]

	// abre a demo
	f, err := os.Open("./demos/mibr-vs-eunited.dem")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	p := dem.NewParser(f)
	header, err := p.ParseHeader()
	fmt.Println("Nome da demo: ", args, "\nMapa: ", header.MapName)

	// parse to end
	err = p.ParseToEnd()
	if err != nil {
		panic(err)
	}
}
