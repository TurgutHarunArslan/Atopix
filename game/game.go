package game

import (
	"time"

	"github.com/TurgutHarunArslan/Atopix/events"
	"github.com/TurgutHarunArslan/Atopix/game/interfaces"
	pmanager "github.com/TurgutHarunArslan/Atopix/game/player"
	"github.com/TurgutHarunArslan/Atopix/game/utils"
)

type Game struct {
	Players  map[string]*pmanager.Player
	Entities []interfaces.EntitiyInterface
	EventBus *events.EventBus
	Ticker   time.Ticker
}

func (g *Game) Init() {
	g.EventBus.Subscribe(events.PlayerJoinEnum, func(d events.EventInterface) {
		data, ok := d.(events.PlayerJoin)
		if !ok {
			return
		}

		g.Players[data.PlayerId] = &pmanager.Player{
			Vector: utils.Vector{
				X: 0,
				Y: 0,
			},
		}

		pl := g.Players[data.PlayerId]

		g.EventBus.Publish(events.PlayerInitilazed{
			PlayerId: data.PlayerId,
			Vector:   pl.Vector,
		})

	})

	defer g.Ticker.Stop()
	for range g.Ticker.C {
		g.Update()
	}
}

func (g *Game) Update() {
	for _, v := range g.Entities {
		v.Update(g)
	}
}

func (g *Game) GetNearestPlayer(vector utils.Vector) *pmanager.Player {
	var distance float32 = 999.00
	var Player *pmanager.Player
	for _, player := range g.Players {
		distancefrom := player.Vector.DistanceFrom(vector)
		if distancefrom < distance {
			Player = player
			distance = distancefrom
		}
	}

	return Player
}

func (g *Game) GetEventBus() *events.EventBus {
	return g.EventBus
}
