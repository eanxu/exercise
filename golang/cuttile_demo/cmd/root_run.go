package cmd

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	route "cuttile_demo/api/router"
	"cuttile_demo/internal/config"
	"cuttile_demo/internal/logger"
	"cuttile_demo/internal/model"
	"sync"
	"time"
)

func run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tasks := []func() error {
		loggerInit,
		printStartgraph,
		postgreSQLConnection,
	}

	for _, t := range tasks {
		if err := t(); err != nil {
			return err
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		err := startHttpServer()
		if err != nil {
			logger.Logger.Error("http server start err",
				zap.Error(err))
		}
		wg.Done()
	}()
	wg.Wait()
	return nil

}

func loggerInit() error {
	lvl := int32(config.C.General.LogLevel)
	microName := config.C.General.Name
	var err error
	logger.Logger, err = logger.LoggerInit(lvl, microName)
	if err != nil {
		os.Exit(1)
	}
	return nil
}

func printStartgraph() error {
	// debug=0, info=1, warning=2, error=3, panic=4, fatal=5
	logLevel := map[int]string{
		0: "debug",
		1: "info",
		2: "warning",
		3: "error",
		4: "panic",
		5: "fatal",
	}
	logger.Logger.Info("starting rscb image server")
	logger.Logger.Info("set logger level",
		zap.String("logLevel", logLevel[config.C.General.LogLevel]))
	return nil
}

func postgreSQLConnection() error {
	logger.Logger.Info("connecting to postgresql",
		zap.String("dsn", config.C.PostgreSQL.DSN))
	err := model.ConnectToDB(config.C.PostgreSQL.DriverName, config.C.PostgreSQL.DSN)
	if err != nil {
		logger.Logger.Error(err.Error())
		os.Exit(1)
	}
	return nil
}

func startHttpServer() error {
	if config.C.General.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()

	mwCORS := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 2400 * time.Hour,
	})
	engine.Use(mwCORS)
	route.SetupRouter(engine)

	if err := engine.Run(config.C.Service.Http); err != nil {
		return err
	}
	return nil
}
