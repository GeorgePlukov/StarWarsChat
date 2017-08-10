
package main

import (
  "log"
  "net/http"

  "github.com/gorilla/websocket"
)

var clients = make (map[*websocket.Conn]bool) // connected clients
var broadcast = make (chan Message)           // broadcast channel

var upgrader = websocket.Upgrader{}

// Message struct
type Message struct {
  Email string `json:"email"`
  Username string `json:"username"`
  Message string `json:"message"`
}


// Handle the incoming websocket
func handleConnections(w http.ResponseWriter, r *http.Request)  {

  // highjack initial get request for a websocket upgrade
  ws,err := upgrader.Upgrade(w,r,nil)
  if err != nil {
    log.Fatal(err)
  }

  // close the connection when funct
  defer ws.Close()

  // register new clients
  clients[ws] = true

  for {
    var msg Message
    // Read in a new message as JSON and map it to a message object
    err := ws.ReadJSON(&msg)
    if err !=nil {
      log.Printf("error: %v", err)
      delete(clients, ws)
      break
    }
    // Send the newly received message to the broadcast channel
    broadcast <- msg
  }

}

func handleMessages()  {
  for {
    // Get next message for broadcast channel
    msg := <-broadcast
    // send to every client connected
    for client := range clients {
      err:= client.WriteJSON(msg)
      if err != nil {
        log.Printf("error: %v", err)
        client.Close()
        delete(clients, client)
      }
    }
  }
}

func main()  {


  // fileserver
  fs := http.FileServer(http.Dir("StarWarsChat/web-backend/public"))
  http.Handle("/", fs)
  http.HandleFunc("/ws", handleConnections)


  // listen for incoming chat messages
  go handleMessages()

  log.Println("http server started on :8000")
  err:= http.ListenAndServe(":8000", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }

}
