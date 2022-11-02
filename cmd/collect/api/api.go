package api

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

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
	wg     *sync.WaitGroup
	conf   *config.Config
	store  Store
	file   FileStore
	wbot   *wbot.WBot
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
		wg:     &sync.WaitGroup{},
		conf:   cnf,
		store:  store,
		file:   fstore,
		wbot:   wbot,
		ctx:    ctx,
		cancel: cancel,
	}, nil
}

// Collect
func (a *API) Collect() {
	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		if err := a.wbot.Crawl(a.conf.StartURL); err != nil {
			// TODO: log this err
			log.Println(err)
		}
	}()

	//
	a.wg.Add(a.conf.Workers)
	for i := 0; i < a.conf.Workers; i++ {
		go func() {
			defer a.wg.Done()

			for resp := range a.wbot.Stream() {
				//
				if resp.Status != http.StatusOK {
					// TODO: log err
					log.Println("bad status")
					continue
				}

				link := resp.URL.String()
				hostname := resp.URL.Hostname()

				details := &collector.Details{
					ID:         hashURL(link),
					Website:    hostname,
					URL:        link,
					HTMLFile:   hostname + "_" + hashURL(link) + ".html",
					FirstVisit: time.Now(),
					LastVisit:  time.Now(),
				}

				// save on db
				if err := a.store.StoreDetails(context.TODO(), details); err != nil {
					// TODO: handle err
					// TODO: delete details record from db if error saving file
					log.Println(err)
					continue
				}

				// save html file
				if err := a.file.SaveHTML(context.TODO(), details.HTMLFile, resp.Body); err != nil {
					// TODO: handle err
					log.Println(err)
					continue
				}
			}
		}()
	}

}

// Refresh
func (a *API) Refresh() {
	// TODO: logic to refresh records on db
}

// Close
func (a *API) Close() {
	a.cancel()

	a.wg.Wait()
	a.wbot.Close()

	a.file.Close()
	a.store.Close()
}
