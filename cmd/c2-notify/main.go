package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"encoding/base64"
	"encoding/json"

	"golang.org/x/net/websocket"
)

// Subscribes to a notification and exits on first message
//
// this can be expanded to repeat indefinitely as an option
// or supply an alternate value for 'origin' should the default
// not be valid for some reason
func main() {
	if len(os.Args) != 3 {
		usage()
	}
	wsUrl := os.Args[1]
	urlParts, err := url.Parse(wsUrl)
	if err != nil {
		log.Fatal(err)
	}
	origin := urlParts.Host
	ws, err := websocket.Dial(wsUrl, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	path := os.Args[2]
	if _, err := ws.Write(subscribe("+", path)); err != nil {
		log.Fatal(err)
	}
	var notification map[string]interface{}
	if err := json.NewDecoder(ws).Decode(&notification); err != nil {
		log.Fatal(err)
	}
	var payload string
	if payloadData, exists := notification["payload"]; !exists {
		log.Fatal("No payload found")
	} else {
		if payloadDecoded, err := base64.StdEncoding.DecodeString(payloadData.(string)); err != nil {
			log.Fatal(err)
		} else {
			payload = string(payloadDecoded)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
	if notification["type"] == "error" {
		log.Fatal(payload)
	}
	fmt.Println(payload)
	if _, err := ws.Write(subscribe("-", path)); err != nil {
		log.Fatal(err)
	}
	ws.Close()
}

func subscribe(op string, path string) []byte {
	return []byte(fmt.Sprintf(`{"op":"%s","path":"%s","group":"n2-notify"}`, op, path))
}

func usage() {
	log.Fatalf(`usage : %s ws://url xpath`, os.Args[0])
}
