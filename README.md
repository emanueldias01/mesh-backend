# Mesh Backend

The Mesh Backend is a lightweight signaling server built with Go for the Mesh real-time meeting platform.

It is responsible for:

* Creating and managing meeting rooms
* Handling participant connections
* WebSocket-based signaling
* Exchanging WebRTC offers, answers, and ICE candidates
* Broadcasting room events such as user join and leave notifications

The server does not process audio or video streams. Media communication is established directly between participants using WebRTC peer-to-peer connections.

## Features

* Room creation
* Room discovery
* WebSocket signaling
* WebRTC negotiation support
* Participant presence events

## Architecture

```text
Client A
    |
    | WebRTC Signaling
    |
Mesh Backend
    |
    | WebRTC Signaling
    |
Client B

Audio/Video Stream
      ↕
 Peer-to-Peer
```

## API Overview

### Rooms

* Create a room
* Retrieve room information

### WebSocket

* Connect to a room
* Exchange signaling messages
* Receive participant events

## Getting Started

### Requirements

* Go 1.24+

### Install Dependencies

```bash
go mod download
```

### Run

```bash
go run .
```

The server will start on:

```text
http://localhost:8080
```

## Run
```bash
    go run ./cmd
```


## Build
```bash
    go build ./cmd
```

## How It Works

1. A user creates a meeting room.
2. Participants join the room through a WebSocket connection.
3. The backend exchanges signaling messages between participants.
4. WebRTC establishes direct peer-to-peer media connections.
5. Audio and video streams bypass the server and flow directly between clients.

## License

This project is available for educational and personal use.
