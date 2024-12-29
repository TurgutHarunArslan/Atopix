package main

import (
	"github.com/TurgutHarunArslan/Atopix/events"
	"github.com/TurgutHarunArslan/Atopix/game"
	"github.com/TurgutHarunArslan/Atopix/game/interfaces"
	Player "github.com/TurgutHarunArslan/Atopix/game/player"
	"github.com/TurgutHarunArslan/Atopix/network"
)

func main() {
	eventBus := events.New()
	game := game.Game{
		EventBus: eventBus,
		Players:  map[string]*Player.Player{},
		Entities: []interfaces.EntitiyInterface{},
	}
	go network.StartServer(game.EventBus)
	game.Init()
}
