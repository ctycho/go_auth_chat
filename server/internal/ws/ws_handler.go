package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req *CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[req.ID] = &Room{
		ID: req.ID,
		Name: req.Name,
		Clients: make(map[string]*Client),
	}

	c.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomID := c.Param("roomId")
	userID := c.Query("userId")
	username := c.Query("username")

	cl := &Client{
		Conn: conn,
		Message: make(chan *Message, 10),
		RoomID: roomID,
		ID: userID,
		Username: username,
	}

	m := &Message{
		Content: "A new user has joined",
		RoomID: roomID,
		Username: username,
	}

	// Register a new client through the hub
	h.hub.Register <- cl
	// Broadcast the message
	h.hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(h.hub)
}

type RoomsRes struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]RoomsRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomsRes{
			ID: r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, rooms)
}

type ClientsRes struct {
	ID string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(c *gin.Context) {
	var clients []ClientsRes
	roomID := c.Param("roomId")

	if _, ok := h.hub.Rooms[roomID]; !ok {
		clients := make([]ClientsRes, 0)
		c.JSON(http.StatusOK, clients)
	}

	for _, cl := range h.hub.Rooms[roomID].Clients {
		clients = append(clients, ClientsRes{
			ID: cl.ID,
			Username: cl.Username,
		})
	}

	c.JSON(http.StatusOK, clients)
}