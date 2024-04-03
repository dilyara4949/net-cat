# TCPChat

This project aims to recreate the functionality of NetCat in a Server-Client Architecture using Go programming language. The server mode listens on a specified port for incoming connections while the client mode connects to a specified port to transmit information to the server.

## Features

- **TCP Connection**: Establishes a TCP connection between the server and multiple clients in a 1-to-many relationship.
- **Client Naming**: Requires clients to provide a name upon connection.
- **Controlled Connections**: Limits the number of connections to a maximum of 10.
- **Messaging**: Clients can send messages to the chat.
- **Message Format**: Messages are timestamped and prefixed with the sender's name.
- **Message History**: Newly joined clients receive previous messages sent to the chat.
- **Notification**: Informs other clients when a new client joins or leaves the group.
- **Continuous Chat**: Clients remain connected even if one leaves the chat.
- **Default Port**: If no port is specified, the default port 8989 is used.

## How to Run

```bash
./TCPChat [port]
```

Run clients to connect to the server:

```bash
./TCPChatClient [server_address] [server_port] [client_name]
```
