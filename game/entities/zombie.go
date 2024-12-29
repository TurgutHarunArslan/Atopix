package entities

import (
	"github.com/TurgutHarunArslan/Atopix/game/interfaces"
	pmanager "github.com/TurgutHarunArslan/Atopix/game/player"
	"github.com/TurgutHarunArslan/Atopix/game/utils"
)

type Zombie struct {
	utils.Vector
}

func (e *Zombie) GetNearestPlayer(game interfaces.GameInterface) *pmanager.Player {
	return game.GetNearestPlayer(e.Vector)
}

func (e *Zombie) Update(game interfaces.GameInterface) {
	// Update Logic
	pl := e.GetNearestPlayer(game)
	println(pl.X)
}
