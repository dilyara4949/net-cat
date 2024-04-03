package main

import (
	"log"
	"net"
)

func main() {
	port := ReadArg()
	chat := NewChat()
	greating := Greating()

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("Listening on the port :" + port + "\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}
		client := NewClient(conn)
		if len(chat.Clients) >= chat.MaxConn {
			client.Msg("Sorry, but chat is too full, try later")
			conn.Close()
			continue
		}

		go func() {
			client.ShowHistory(chat.History)
			client.Msg(greating)
			client.NameClient(chat)
			

			chat.Broadcast(client, client.Name+" has joined our chat...\n")
			log.Printf("new client has joined: %s", client.Name)

			chat.Clients = append(chat.Clients, client)

			go client.Run(chat)
			go client.Read(chat)
		}()
	}
}

func errHandleLogPrint(err error, msg string) {
	if err != nil {
		log.Print(msg, "error:", err.Error())
	}
}
