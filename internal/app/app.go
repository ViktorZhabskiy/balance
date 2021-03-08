package app

import (
	v1 "balance/internal/api/v1"
	"balance/internal/api/v1/handlers"
	"balance/internal/config"
	"balance/internal/repository"
	"balance/internal/service"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {

	if err := config.Init(); err != nil {
		log.Fatalf("Failed read config. Reason: %s", err)
	}

	InitLogger()

	// open connect to db
	db, err := NewPsqlDb()
	if err != nil {
		logrus.Fatalf("Error connect to database. Reason: %s", err)
	}

	balanceRepo := repository.NewBalance(db)
	balanceSrv := service.NewBalance(balanceRepo)

	endpoints := handlers.NewHandlers(balanceSrv)
	r := v1.NewV1Routes(endpoints)

	// init http server
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("server.apiV1.port")),
		Handler: r,
	}

	// start http server
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			logrus.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	logrus.Infof("Http server started on port %s", viper.GetString("server.apiV1.port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	logrus.Info("shutting down...")
	err = httpServer.Shutdown(ctx)
	if err != nil {
		logrus.Fatalf("server Shutdown Failed:%+s", err)
	}

}

func NewPsqlDb() (*sqlx.DB, error) {
	psqlDataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		viper.GetString("db.psql.host"),
		viper.GetString("db.psql.port"),
		viper.GetString("db.psql.username"),
		viper.GetString("db.psql.dbname"),
		viper.GetString("db.psql.sslmode"),
		viper.GetString("db.psql.password"),
	)

	db, err := sqlx.Connect("postgres", psqlDataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetOutput(os.Stdout)

	logrus.SetLevel(logrus.Level(viper.GetInt("logger.level")))
}
