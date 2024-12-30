package network

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/TurgutHarunArslan/Atopix/events"
	"github.com/google/uuid"
)

type Network struct {
	Server      net.Listener
	EventBus    *events.EventBus
	Connections map[string]*Conn
}

func (s *Network) SetEventBus(bus *events.EventBus) {
	s.EventBus = bus
}

func (s *Network) Broadcast(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error serializing data:", err)
		return
	}
	for _, conn := range s.Connections {
		fmt.Fprint(conn.Conn, string(jsonData))
	}
}

func (s *Network) SetupEvents() {
	s.EventBus.Subscribe(events.PlayerInitilazedEnum, func(d events.EventInterface) {
		data, ok := d.(events.PlayerInitilazed)

		if !ok {
			return
		}

		s.Broadcast(data)

	})

	s.EventBus.Subscribe(events.ServerPlayerMovedEnum, func(d events.EventInterface) {
		data, ok := d.(events.ServerPlayerMoved)

		if !ok {
			return
		}

		s.Broadcast(data)
	})
}

func (s *Network) Init() {
	server, err := net.Listen("tcp4", ":3000")
	if err != nil {
		log.Fatal("Server couldn't be initialized: ", err)
	}
	defer server.Close()

	s.Server = server

	s.SetupEvents()

	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		id := uuid.New().String()
		c := Conn{
			Id:   id,
			Conn: client,
		}

		s.Connections[id] = &c
		go handleConnection(s.EventBus, &c)
	}
}
