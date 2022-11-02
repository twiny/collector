package api

import (
	"context"

	//
	"collector/pkg/collector/v1"
	"collector/pkg/config"
	"collector/service/badgerdb"
	"collector/service/localstore"

	//
	"github.com/twiny/wbot"
)

// API
type API struct {
	conf   *config.Config
	store  Store
	file   FileStore
	wbot   *wbot.WBot
	stream chan *collector.Details
	ctx    context.Context
	cancel context.CancelFunc
}

// NewAPI
func NewAPI(cf string) (*API, error) {
	cnf, err := config.ParseConfig(cf)
	if err != nil {
		return nil, err
	}

	allowed := []string{
		cnf.URLFilter,
	}

	opts := []wbot.Option{
		wbot.SetMaxDepth(cnf.MaxDepth),
		wbot.SetParallel(cnf.Workers),
		wbot.SetRateLimit(cnf.RateLimit.Rate, cnf.RateLimit.Interval),
	}

	if cnf.URLFilter != "" {
		opts = append(opts, wbot.SetFilter(allowed, nil))
	}

	if len(cnf.UserAgents) > 0 {
		opts = append(opts, wbot.SetUserAgents(cnf.UserAgents))
	}

	if len(cnf.Proxies) > 0 {
		opts = append(opts, wbot.SetProxies(cnf.Proxies))
	}

	wbot := wbot.NewWBot(opts...)

	// store
	store, err := badgerdb.NewBadgerDB(cnf.Storage.DB)
	if err != nil {
		return nil, err
	}

	// file store
	fstore, err := localstore.NewLocalStore(cnf.Storage.HTML)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	//
	return &API{
		conf:   cnf,
		store:  store,
		file:   fstore,
		wbot:   wbot,
		stream: make(chan *collector.Details, cnf.Workers),
		ctx:    ctx,
		cancel: cancel,
	}, nil
}

// Collect
