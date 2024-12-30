package events

import "github.com/TurgutHarunArslan/Atopix/game/utils"

const (
	PlayerJoinEnum        string = "PlayerJoin"
	PlayerInitilazedEnum  string = "PlayerInitilazed"
	ClientPlayerMovedEnum string = "ClientPlayerMoved"
	ServerPlayerMovedEnum string = "ServerPlayerMoved"
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

type PlayerInitilazed struct {
	PlayerId string
	Vector utils.Vector
}

func (p PlayerInitilazed) Type() string {
	return PlayerInitilazedEnum
}

type ClientPlayerMoved struct {
	PlayerId string
	Vector   utils.Vector
}

func (p ClientPlayerMoved) Type() string {
	return ClientPlayerMovedEnum
}

type ServerPlayerMoved struct {
	PlayerId string
	Vector   utils.Vector
}

func (p ServerPlayerMoved) Type() string {
	return ServerPlayerMovedEnum
}
