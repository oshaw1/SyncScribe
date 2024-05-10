package handlers

import (
	"SyncScribe/backend/crdt"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: true,
}

var clients = make(map[string]map[*websocket.Conn]uint)
var crdts = make(map[string]*crdt.OrderedListCRDT)
var mutex = &sync.Mutex{}
var nextSiteID uint = 1

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading WebSocket connection:", err)
		return
	}
	defer conn.Close()

	noteID := strings.TrimPrefix(r.URL.Path, "/ws/")
	log.Printf("WebSocket connection established for note: %s", noteID)

	mutex.Lock()
	if _, ok := crdts[noteID]; !ok {
		crdts[noteID] = &crdt.OrderedListCRDT{}
	}
	if _, ok := clients[noteID]; !ok {
		clients[noteID] = make(map[*websocket.Conn]uint)
	}
	siteID := nextSiteID
	nextSiteID++
	clients[noteID][conn] = siteID
	mutex.Unlock()

	log.Printf("New WebSocket connection for note: %s, SiteID: %d", noteID, siteID)

	initialState, _ := json.Marshal(crdts[noteID])
	conn.WriteMessage(websocket.TextMessage, initialState)

	defer func() {
		mutex.Lock()
		delete(clients[noteID], conn)
		mutex.Unlock()
		log.Printf("WebSocket connection closed for note: %s, SiteID: %d", noteID, siteID)
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading WebSocket message:", err)
			break
		}

		log.Printf("Received WebSocket message for note %s: %s", noteID, string(message))

		var opMsg OperationMessage
		err = json.Unmarshal(message, &opMsg)
		if err != nil {
			log.Println("Error parsing WebSocket message:", err)
			continue
		}

		var op crdt.Operation
		switch opMsg.Type {
		case insertOperation:
			op = &crdt.InsertOperation{
				Element: crdt.Element{
					ID:      opMsg.Element.ID,
					Content: opMsg.Element.Content,
				},
				SiteID: siteID,
			}
		case deleteOperation:
			op = &crdt.DeleteOperation{
				ElementID: opMsg.Element.ID,
				SiteID:    siteID,
			}
		default:
			log.Println("Unknown operation type:", opMsg.Type)
			continue
		}

		mutex.Lock()
		crdts[noteID].Apply(op)
		mutex.Unlock()
		log.Printf("Applied operation to CRDT for note %s: %+v", noteID, op)

		broadcastUpdate(noteID, op, conn)
	}
}

func broadcastUpdate(noteID string, op crdt.Operation, sender *websocket.Conn) {
	for client := range clients[noteID] {
		if client != sender {
			var opMsg OperationMessage
			switch op := op.(type) {
			case *crdt.InsertOperation:
				opMsg = OperationMessage{
					Type: insertOperation,
					Element: ElementMessage{
						ID:       op.Element.ID,
						Content:  op.Element.Content,
						Position: op.Element.Position,
					},
				}
			case *crdt.DeleteOperation:
				opMsg = OperationMessage{
					Type: deleteOperation,
					Element: ElementMessage{
						ID: op.ElementID,
					},
				}
			}

			message, _ := json.Marshal(opMsg)
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Error broadcasting update:", err)
				mutex.Lock()
				delete(clients[noteID], client)
				mutex.Unlock()
				client.Close()
			} else {
				log.Printf("Broadcasted update to client for note %s: %s", noteID, message)
			}
		}
	}
}

type ElementMessage struct {
	ID       string        `json:"id"`
	Content  string        `json:"content"`
	Position crdt.Position `json:"position"`
}

type OperationMessage struct {
	Type    int            `json:"type"`
	Element ElementMessage `json:"element"`
}

const (
	insertOperation = iota
	deleteOperation
)
