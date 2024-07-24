# Group Chatting App

## Overview

This project is a group chatting application built in Go, demonstrating TCP server-client communication and broadcasting functionality.

## Features and Summary for Each Step

### Step 1: Build TCP Server
- **Feature**: Establishes a TCP server to handle client connections.
- **Summary**: Sets up a server that listens on a specified IP and port, accepting incoming client connections.

### Step 2: User Management and Broadcasting
- **Feature**: Tracks online users and maintains a user map.
- **Summary**: When a user connects, they are added to an online map, and the server broadcasts their status to all connected users.

### Step 3: Real-Time Messaging
- **Feature**: Notifies all users when someone comes online or sends a message.
- **Summary**: Implements functionality for real-time broadcasting of messages from one user to all other connected users.

### Step 4: Reorgnizing user components
- **Feature**: Detects when a user disconnects and notifies others.
- **Summary**: Adds functionality to remove users from the online map and broadcast their disconnection to remaining users.

### Step 5: Command Handling: Query all users online
- **Feature**: Allows users to send private messages.
- **Summary**: Implements direct messaging between users, ensuring messages are only received by the intended recipient.

### Step 6: Command Handling: rename user
- **Feature**: Supports user commands for various actions.
- **Summary**: Introduces a command system allowing users to execute specific actions, such as listing online users or sending private messages, through predefined commands.

## Usage

- Start the server to listen on the specified IP and port.
- Connect multiple clients to the server using `nc 127.0.0.1 8888`
- Each client will receive notifications of other users' activities in real-time.

## Contributing

Feel free to open issues or submit pull requests for improvements and bug fixes.

## License

This project is open-source and available under the MIT License.

For more details, visit the [repository](https://github.com/yanglyu520/instantMessagingApp).