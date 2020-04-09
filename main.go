package main

import (
	"fmt"
	// "html/template"
	// "net/http"
	"os"

	dem "github.com/markus-wa/demoinfocs-golang"
	events "github.com/markus-wa/demoinfocs-golang/events"
)

type PlayerStructure struct {
	plName string
	kills  int
	deaths int
	round  [25]RoundStructure
}

type RoundStructure struct {
	kills int
}

func main() {
	// http.HandleFunc("/", index)
	// http.ListenAndServe(":8080", nil)

	// pega os argumentos passados como parametro
	args := os.Args[1:]

	// abre a demo
	f, err := os.Open("./demos/demo.dem")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	p := dem.NewParser(f)
	header, err := p.ParseHeader()
	fmt.Println("Nome da demo: ", args, "\nMapa: ", header.MapName)

	player := make(map[string]*PlayerStructure)
	initializePlayers(p, player)
	registerKillsAndDeaths(p, player)

	round := 0
	roundStart := 0
	roundEnd := 0

	//Registro de inicio de round
	p.RegisterEventHandler(func(e events.RoundStart) {
		round++
		roundStart = round
	})

	//Registro de término de round
	p.RegisterEventHandler(func(e events.RoundEndOfficial) {
		roundEnd = round
		fmt.Println("ROUND ATUAL:", roundEnd)
	})

	if roundStart == roundEnd {
		p.RegisterEventHandler(func(e events.Kill) {
			fmt.Println(e.Killer, "killed", e.Victim, "with", e.Weapon)
		})

	}

	// parse to end
	err = p.ParseToEnd()
	if err != nil {
		panic(err)
	}

	printFinalFrags(player)

}

func diedBlind(e events.Kill, tickNumber int) {
	if e.Victim.IsBlinded() {
		fmt.Println("[tick: ", tickNumber, "]", e.Victim.Name, "died blinded")
	}
}

func printFinalFrags(player map[string]*PlayerStructure) {
	for _, entry := range player {
		if entry.kills != 0 || entry.deaths != 0 {
			fmt.Printf("%s %d/%d\n", entry.plName, entry.kills, entry.deaths)
		}
	}
}

func registerKillsAndDeaths(p *dem.Parser, player map[string]*PlayerStructure) {
	p.RegisterEventHandler(func(e events.Kill) {
		player[e.Killer.Name].kills++
		player[e.Victim.Name].deaths++
	})
}

func initializePlayers(p *dem.Parser, player map[string]*PlayerStructure) {
	p.RegisterEventHandler(func(e events.PlayerConnect) {
		newPlayerStructure := new(PlayerStructure)
		newPlayerStructure.plName = e.Player.Name
		newPlayerStructure.kills = 0
		newPlayerStructure.deaths = 0
		player[e.Player.Name] = newPlayerStructure
	})
}

func getChatHistory(p *dem.Parser) {
	p.RegisterEventHandler(func(e events.ChatMessage) {
		if e.IsChatAll {
			fmt.Print("*(ALL) ")
		} else {
			fmt.Println("*(TEAM) ")
		}
		fmt.Println(e.Sender, ":", e.Text)
	})
}
