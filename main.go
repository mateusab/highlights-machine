package main

import (
	"fmt"
	"os"

	dem "github.com/markus-wa/demoinfocs-golang"
	events "github.com/markus-wa/demoinfocs-golang/events"
)

type player struct {
	name  string
	kills int
}

func main() {
	// pega os argumentos passados como parametro
	args := os.Args[1:]

	round := 0
	roundStart := 0
	roundEnd := 0
	// abre a demo
	f, err := os.Open("./demos/mibr.dem")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	p := dem.NewParser(f)
	header, err := p.ParseHeader()
	fmt.Println("Nome da demo: ", args, "\nMapa: ", header.MapName)

	p.RegisterEventHandler(func(e events.RoundStart) {
		round++
		roundStart = round
	})

	p.RegisterEventHandler(func(e events.RoundEndOfficial) {
		roundEnd = round
		fmt.Println("Round", roundEnd)
	})

	if roundStart == roundEnd {
		p.RegisterEventHandler(func(e events.Kill) {
			killCounter(e)
		})
	}

	// parse to end
	err = p.ParseToEnd()
	if err != nil {
		panic(err)
	}
}

func killCounter(e events.Kill) {
	fmt.Println(e.Killer.Name, "killed", e.Victim.Name)
}

func diedBlind(e events.Kill, tickNumber int) {
	if e.Victim.IsBlinded() {
		fmt.Println("[tick: ", tickNumber, "]", e.Victim.Name, "died blinded")
	}
}
