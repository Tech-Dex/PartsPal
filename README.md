# PartsPal

Whether you're a car enthusiast, a mechanic, or a regular vehicle owner, PartsPal is designed to help you quickly find
the exact parts you need for your vehicle maintenance or repair projects.

## Features

- Search and compare car parts from multiple providers
- Real-time price and availability updates
- User-friendly Fyne desktop application
- Websocket API for easy integration into other applications

## Getting Started

### Prerequisites

- Go (version 1.16 or higher)
- Fyne (installation guide: https://github.com/fyne-io/fyne)
- Gorilla Websocket (installation guide: https://github.com/gorilla/websocket)

### Installation

1. Clone the repository:

```shell
git clone <repository-url>
```

2. Install dependencies:

```shell
go mod download
```

3.1. Build the Fyne application:

```shell
go build cmd/parts-pal/main.go
```

3.2. Build the Websocket server:

```shell
go build cmd/api/main.go
```

4. Run the built executables:

```shell
./main
```

5. Packaging the Fyne application (optional):

```shell
fyne package -os <operating-system> -icon <icon-file> # https://developer.fyne.io/started/packaging
```

### Usage

#### Fyne Application

1. Run the Fyne application:

```shell
./main
```

2. Enter the product code information and click the "Search" button

#### Websocket API

1. Run the Websocket server:

```shell
./main
```

2. Connect to the server using a websocket client on localhost:3000/ws/scrape
3. Send a simple byte message containing the product code information: "27025"
4. The server will respond with a JSON message containing the product information. There are 2 types of messages: "
   bestDeal" and "deal"
5. The "bestDeal" message contains the best deal for the product, while the "deal" message contains a deal from one of
   the providers

#### Example WebSocket Client Code (using Gorilla Websocket):

```go
package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	// Establish a WebSocket connection
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:3000/ws/scrape", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Send a simple message with the product code
	err = conn.WriteMessage(websocket.TextMessage, []byte("PRODUCT_CODE"))
	if err != nil {
		log.Fatal(err)
	}

	// Receive and handle WebSocket messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the received message
		log.Println("Received message:", string(message))
	}
}
```

Replace "PRODUCT_CODE" in the code above with the actual product code you want to search for.