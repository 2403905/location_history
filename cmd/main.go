package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/2403905/location_history/internal/api"
	"github.com/2403905/location_history/internal/config"
	"github.com/2403905/location_history/internal/repository"
	"github.com/2403905/location_history/internal/service"
	"github.com/2403905/location_history/logger"
)

func main() {
	// Init config
	cf := config.InitConfig()
	// Init logger
	logger.InitPretty(zerolog.Level(cf.LogLevel))

	repo := repository.NewRepo()
	locationService := service.NewLocation(&repo)

	// Init handler
	handle := api.NewHandler(&locationService)

	srv := &http.Server{Addr: cf.ApiListener, Handler: handle}
	log.Info().Msgf("Start service on http://%s", cf.ApiListener)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Info().Msgf("Service - listen: %s", err)
		}
	}()
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		var once sync.Once
		for range signalChan {
			once.Do(func() {
				log.Warn().Msgf("Service received a shutdown signal...")
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				if err := srv.Shutdown(ctx); err != nil {
					log.Error().Err(err).Msg("Error")
				}
				log.Info().Msg("Service - Received an interrupt closing connection...")
				log.Warn().Msgf("Service stopped successfully")
				cleanupDone <- true
			})
		}
	}()
	log.Warn().Msgf("Service started successfully")
	<-cleanupDone
}
