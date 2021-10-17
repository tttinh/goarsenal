package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tttinh/goarsenal/app/wager"
	"github.com/tttinh/goarsenal/infra/config"
	"github.com/tttinh/goarsenal/infra/log"
	"github.com/tttinh/goarsenal/infra/persistence"
	httptransport "github.com/tttinh/goarsenal/infra/transport/http"
	"github.com/tttinh/goarsenal/repository"
)

func main() {
	// Loading configuration.
	cfg := config.NewConfig()

	// Init log
	logger := log.NewLogger(cfg.Server.Mode)

	// Connecting DB.
	db := persistence.NewDB(cfg)

	// Setup Gin.
	gin.SetMode(cfg.Server.Mode)
	//r := gin.Default()
	r := gin.New()
	httpLogger := logger.With("component", "http")
	r.Use(httptransport.Logger(httpLogger))
	r.Use(httptransport.Recovery(httpLogger))

	// Create application logic services.
	groupRepository := repository.NewGroupRepository(db)
	var groupService wager.Service
	groupService = wager.NewService(groupRepository)
	groupService = wager.NewLoggingService(logger.With("component", "wager"), groupService)
	wager.SetRoutes(r, groupService)

	// Start server.
	run(logger, cfg, r)
}

func run(logger log.Logger, cfg config.Config, r *gin.Engine) {
	readTimeout := time.Duration(cfg.Server.ReadTimeout) * time.Second
	writeTimeout := time.Duration(cfg.Server.WriteTimeout) * time.Second
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           cfg.Server.Port,
		Handler:        r,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	errs := make(chan error, 2)
	go func() {
		logger.Info("starting server on port ", cfg.Server.Port)
		errs <- server.ListenAndServe()
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Info("server stopped: ", <-errs)
}
