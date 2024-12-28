package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"github.com/google/uuid"
)

func handleConnection(client net.Conn) {
	id := uuid.New().String()
	c := Conn{
		Id:   id,
		Conn: client,
	}

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

	}
}
