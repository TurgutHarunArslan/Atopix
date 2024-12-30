package network

import (
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



func (s *Network) Broadcast(data interface{}){
	for _, conn := range s.Connections {
		fmt.Fprint(conn.Conn,data)
	}
}

func (s *Network) SetupEvents()  {
	s.EventBus.Subscribe(events.PlayerInitilazedEnum,func(d events.EventInterface) {
		data,ok := d.(events.PlayerInitilazed)

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
