package routes

import (
	"context"
	"errors"
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"homework/infra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx infra.ServiceContext) {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(ctx.Middleware.Logger())
	r.POST("/login", ctx.Ctl.AuthHandler.Login)
	r.POST("/otp", ctx.Ctl.PasswordResetHandler.Create)
	r.PUT("/otp/:id", ctx.Ctl.PasswordResetHandler.Update)
	r.PUT("/users/:id", ctx.Ctl.UserHandler.Update)
	r.POST("/register", ctx.Ctl.UserHandler.Registration)

	r.Use(ctx.Middleware.Jwt.AuthJWT())

	gracefulShutdown(ctx, r.Handler())
}

func gracefulShutdown(ctx infra.ServiceContext, handler http.Handler) {
	srv := &http.Server{
		Addr:    ctx.Cfg.ServerPort,
		Handler: handler,
	}

	if ctx.Cfg.ShutdownTimeout == 0 {
		launchServer(srv, ctx.Cfg.ServerPort)
		return
	}

	go func() {
		launchServer(srv, ctx.Cfg.ServerPort)
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	appContext, cancel := context.WithTimeout(context.Background(), time.Duration(ctx.Cfg.ShutdownTimeout)*time.Second)
	defer cancel()
	if err := srv.Shutdown(appContext); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching appContext.Done(). timeout of ShutdownTimeout seconds.
	select {
	case <-appContext.Done():
		log.Println(fmt.Sprintf("timeout of %d seconds.", ctx.Cfg.ShutdownTimeout))
	}
	log.Println("Server exiting")
}

func launchServer(server *http.Server, port string) {
	// service connections
	log.Println("Listening and serving HTTP on", port)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen: %s\n", err)
	}
}
