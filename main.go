package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Conf and start VPN server
	server := NewServer("0.0.0.0", 8080)
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}

type Server struct {
	listenAddress string
	listenPort    int
}

func NewServer(address string, port int) *Server {
	return &Server{
		listenAddress: address,
		listenPort:    port,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.listenAddress, s.listenPort))
	if err != nil {
		return err
	}

	defer listener.Close()

	fmt.Printf("Server VPN started on : %s:%d\n", s.listenAddress, s.listenPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error on receiving connection : ", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

// Handle connection to manage input connection
func (s *Server) handleConnection(conn net.Conn) {
	// TODO: Add algo to manage VPN connection
}
