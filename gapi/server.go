package gapi

import (
	"fmt"

	db "github.com/bbsemih/gobank/internal/db/sqlc"
	"github.com/bbsemih/gobank/pb"
	"github.com/bbsemih/gobank/pkg/token"
	"github.com/bbsemih/gobank/pkg/util"
)

type Server struct {
	pb.UnimplementedGoBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	return server, nil
}
