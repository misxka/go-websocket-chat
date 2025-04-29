# Go WebSocket Chat Application

This is a WebSocket-based chat application built with Go, using Redis Pub/Sub for message broadcasting. The application supports multiple clients and can scale horizontally by leveraging Redis for distributed messaging.

## Features

- Real-time chat using WebSocket.
- Redis Pub/Sub for message broadcasting.
- Dockerized setup for easy deployment.
- Scalable architecture for distributed environments.

---

## Prerequisites

- [Go](https://golang.org/) (1.18 or later)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/misxka/go-websocket-chat.git
cd go-websocket-chat
```

### 2. Set Up Environment Variables

Create a `.env` file in the project root and define the following variables:

```bash
REDIS_PASSWORD=your_redis_password
REDIS_USER=your_redis_user
REDIS_USER_PASSWORD=your_redis_user_password
```

### 3. Start the Application

Use Docker Compose to start the Redis server:

```bash
docker-compose up -d
```

This will start a Redis container with the configuration specified in `docker-compose.yaml`.

### 4. Run the Application

Build & run the Go application:

```bash
go build -o bin/
./bin/go-websocket-chat
```


## How It Works

1. WebSocket Communication:
  - Clients connect to the server via WebSocket.
  - Messages sent by clients are received in the `read` method of the `Client` struct.

2. Redis Pub/Sub:
  - Messages from WebSocket clients are published to a Redis channel.
  - The `write` method subscribes to the Redis channel and forwards messages to WebSocket clients.

3. Client Management:
  - The `ClientManager` struct handles client registration, unregistration, and broadcasting.


## Configuration

### Redis Configuration

The Redis service is configured in `docker-compose.yaml`. You can customize the following:

  - Ports: Change the `6379:6379` mapping to use a different port.
  - Volumes: Use a different directory for persistent data storage.
  - Environment Variables: Update the `.env` file to set Redis credentials.