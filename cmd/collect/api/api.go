package api

import (
	"collector/pkg/config"
	"context"

	"github.com/twiny/wbot"
)

// API
type API struct {
	conf   *config.Config
	store  Store
	file   FileStore
	wbot   *wbot.WBot
	stream chan *Data
	ctx    context.Context
	cancel context.CancelFunc
}

// NewAPI
func NewAPI(cf string) (*API, error) {

	//
	return nil, nil
}

// Collect
