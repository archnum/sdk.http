/*
####### sdk.http (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package server

import (
	"context"
	"log"
	"time"

	"github.com/archnum/sdk.application/container"
	"github.com/archnum/sdk.base/gotracker"
	"github.com/archnum/sdk.base/kv"
	"github.com/archnum/sdk.base/logger"
	"github.com/archnum/sdk.base/logger/level"
	"github.com/archnum/sdk.base/util"

	_logger "github.com/archnum/sdk.application/component/logger"
	"github.com/archnum/sdk.http/component/handler"
	"github.com/archnum/sdk.http/server"
)

type (
	implComponent struct {
		*container.Component
		logger    *logger.Logger
		server    *server.Server
		goTracker *gotracker.GoTracker
	}
)

func New(c container.Container) *implComponent {
	return &implComponent{
		Component: container.NewComponent("http.server", c),
	}
}

//////////////////////
/// Implementation ///
//////////////////////

func (impl *implComponent) Build() error {
	c := impl.C()
	logger := _logger.Value(c)

	p := &server.Params{
		Config:  config(c),
		Handler: handler.Value(c),
		Logger:  logger.NewStdLogger(level.Error, "[http.server]", log.Llongfile),
	}

	server, err := server.New(p)
	if err != nil {
		return err
	}

	impl.logger = logger
	impl.server = server

	return nil
}

func (impl *implComponent) Start() error {
	gt := gotracker.New(gotracker.WithLogger(impl.logger))

	gt.Go( //@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		"http.server",
		func(_ context.Context) error {
			defer gt.Stop()

			server := impl.server

			impl.logger.Info( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
				util.If(server.TLS(), "Server HTTPS", "Server HTTP"),
				kv.String("addr", server.Addr()),
			)

			return server.Start()
		},
	)

	impl.goTracker = gt

	select {
	case <-gt.Done():
		return gt.Err()
	case <-time.After(100 * time.Millisecond):
		return nil
	}
}

func (impl *implComponent) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := impl.server.Stop(ctx)

	impl.goTracker.Stop()
	impl.goTracker.Wait()

	return err
}

/*
####### END ############################################################################################################
*/
