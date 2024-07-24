# Simple TCP Server

This Go program demonstrates a basic TCP server that listens for incoming connections on a specified IP address and port.

## Key Points

- **Initialization**: A `Server` struct is created with IP and port.
- **Starting the Server**: The server listens for incoming connections and handles each in a separate goroutine.
- **Connection Handling**: Each accepted connection is processed concurrently with goroutine

