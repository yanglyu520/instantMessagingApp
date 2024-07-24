# Simple TCP Server with Broadcasting

This Go program demonstrates a basic TCP server that listens for incoming connections on a specified IP address and port, and broadcasts messages to all connected users.

## Key Points

- **Server Initialization**: A `Server` struct is created with IP, port, a map to track online users, and a channel for broadcasting messages.
- **Starting the Server**: The server listens for incoming connections and handles each in a separate goroutine.
- **Connection Handling**: Each accepted connection is processed concurrently, adding the user to an online map and broadcasting messages.
- **Broadcasting**: The server sends messages to all connected users when a user comes online, goes offline, or sends a message.
