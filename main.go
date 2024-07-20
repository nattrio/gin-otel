package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/nattrio/gin-otel/app/post"
	"github.com/nattrio/gin-otel/config"
	"github.com/nattrio/gin-otel/db"
	"github.com/nattrio/gin-otel/logger"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"github.com/uptrace/uptrace-go/uptrace"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

const serviceName = "gin-otel"

func main() {
	cfg := config.Value()
	ctx := context.Background()

	logger := otelzap.New(logger.NewZap())
	defer logger.Sync()

	undo := otelzap.ReplaceGlobals(logger)
	defer undo()

	uptrace.ConfigureOpentelemetry()
	defer uptrace.Shutdown(ctx)

	postgres := db.NewPGXPool(cfg.PostgresURL())
	defer postgres.Close()

	postRepository := post.NewPostRepo(postgres)
	postUsecase := post.NewPostUsecase(postRepository)
	postHandler := post.NewPostHandler(postUsecase)

	router := gin.Default()

	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type", "TransactionID"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))

	router.Use(otelgin.Middleware(serviceName))

	router.POST("/posts", postHandler.CreatePost)
	router.GET("/posts", postHandler.GetPosts)
	router.GET("/posts/:id", postHandler.GetPost)
	router.PATCH("/posts/:id", postHandler.UpdatePost)
	router.DELETE("/posts/:id", postHandler.DeletePost)

	srv := http.Server{
		Addr:              ":" + cfg.App.Port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Graceful shutdown
	ch := make(chan struct{})

	stgCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		<-stgCtx.Done()
		logger.Info("Shutting down...")
		if err := srv.Shutdown(ctx); err != nil {
			logger.Info("HTTP server Shutdown: " + err.Error())
		}
		close(ch)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatal("HTTP server ListenAndServe: " + err.Error())
	}

	<-ch
	logger.Info("Shutdown gracefully")
}
