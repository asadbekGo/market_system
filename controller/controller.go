package controller

import (
	"github.com/asadbekGo/market_system/config"
	"github.com/asadbekGo/market_system/storage"
)

type Conn struct {
	cfg     *config.Config
	storage storage.StorageI
}

func NewController(cfg *config.Config, strg storage.StorageI) *Conn {
	return &Conn{cfg: cfg, storage: strg}
}
