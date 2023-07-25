package app

import (
	"fmt"
	"github.com/Nikkoz/mp.gateway/internal/configs"
	httpDelivery "github.com/Nikkoz/mp.gateway/internal/delivery/http"
	grpcRepository "github.com/Nikkoz/mp.gateway/internal/repository/store/grpc"
	storeRepository "github.com/Nikkoz/mp.gateway/internal/repository/store/storage"
	storeUseCase "github.com/Nikkoz/mp.gateway/internal/useCase/factories/store"
	grpcClient "github.com/Nikkoz/mp.gateway/pkg/grpc"
	"github.com/Nikkoz/mp.gateway/pkg/types/context"
	"github.com/Nikkoz/mp.gateway/pkg/types/logger"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var config *configs.Config

func init() {
	envInit()
	configInit()
}

func Run() {
	ctx := context.Empty()
	defer ctx.Cancel()

	logger.New(config.App.Environment.IsProduction(), config.Log.Level.String())

	conn, conClose := connectionDB()
	defer conClose()

	migrate(conn)

	grpcConn, err := grpcClient.New(ctx, config.Grpc.Host, config.Grpc.Port, config.App.Name, config.App.Version).GetConnection()
	if err != nil {
		_ = logger.ErrorWithContext(ctx, errors.Wrap(err, "failed to create grpc client"))

		return
	}

	var (
		grpc         = grpcRepository.New(grpcConn, grpcRepository.Options{})
		storage      = storeRepository.New(conn, storeRepository.Options{})
		ucStore      = storeUseCase.New(storage, grpc, storeUseCase.Options{})
		listenerHttp = httpDelivery.New(ucStore, httpDelivery.Options{})
	)

	listenerHttp.Run(*config)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err := <-listenerHttp.Notify():
		logger.Error(fmt.Errorf("app - Run http server: %v", err))
	case done := <-ctx.Done():
		logger.Info(fmt.Sprintf("app - Run - ctx.Done: %v", done))
	}
}

func envInit() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("load environment failed: %v\n", err)
	}
}

func configInit() {
	cfg, err := configs.New()
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	config = cfg
}
