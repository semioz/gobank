package api

import (
	db "github.com/bbsemih/gobank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	r := gin.Default()

	r.POST("/accounts", server.createAccount)
	r.GET("/accounts/:id", server.getAccount) //TODO
	r.GET("/accounts", server.listAccounts)   //TODO

	server.router = r
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}