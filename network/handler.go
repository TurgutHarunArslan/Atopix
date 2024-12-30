package network

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/TurgutHarunArslan/Atopix/events"
	"github.com/TurgutHarunArslan/Atopix/game/utils"
	"github.com/TurgutHarunArslan/Atopix/network/packets"
)

func handleConnection(EventBus *events.EventBus, c *Conn) {
	tmp := make([]byte, 4096)
	defer c.Conn.Close()

	for {
		n, err := c.Conn.Read(tmp)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}

		var tempMap map[string]interface{}
		if err := json.Unmarshal(tmp[:n], &tempMap); err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return
		}

		msgType, ok := tempMap["type"].(string)

		if !ok {
			fmt.Println("Invalid Or Missing type data")
			continue
		}

		switch msgType {
		case "move":
			var packet packets.PlayerClientPositionMovedPacket
			if err := json.Unmarshal(tmp[:n], &tempMap); err != nil {
				continue
			}
			event := events.ClientPlayerMoved{
				PlayerId: c.Id,
				Vector:   utils.Vector{X: packet.X, Y: packet.Y},
			}
			EventBus.Publish(event)
		}

	}
}
