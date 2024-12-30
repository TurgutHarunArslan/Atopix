package packets

type PlayerInitilizedPacket struct {
	PlayerId string
	X,Y float32
}

type PlayerClientPositionMovedPacket struct{
	X,Y float32
}

type PlayerServerPositionMovedPacket struct {
	PlayerId string
	X,Y float32
}