package main

import (
	"log"
	"github.com/ctycho/go_auth_chat/db"
	"github.com/ctycho/go_auth_chat/internal/user"
	"github.com/ctycho/go_auth_chat/router"
	"github.com/ctycho/go_auth_chat/internal/ws"
)
func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Could not connect to the Database %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("localhost:8000")
}