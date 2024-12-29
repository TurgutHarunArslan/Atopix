package interfaces

import (
	"github.com/TurgutHarunArslan/Atopix/events"
	pmanager "github.com/TurgutHarunArslan/Atopix/game/player"
	"github.com/TurgutHarunArslan/Atopix/game/utils"
)

type GameInterface interface {
	GetNearestPlayer(vector utils.Vector) *pmanager.Player
	GetEventBus() *events.EventBus
}
