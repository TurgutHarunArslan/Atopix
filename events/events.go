package events

import "github.com/TurgutHarunArslan/Atopix/game/utils"

const (
	PlayerJoinEnum     string = "PlayerJoin"
	PositionChangeEnum string = "PositionChange"
)

type EventInterface interface {
	Type() string
}

type PlayerJoin struct {
	PlayerId string
}

func (p PlayerJoin) Type() string {
	return PlayerJoinEnum
}

type PositionChange struct {
	PlayerId string
	Vector   utils.Vector
}

func (p PositionChange) Type() string {
	return PositionChangeEnum
}
