package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/harshit3011/students-api/internal/config"
	"github.com/harshit3011/students-api/internal/http/handlers/student"
)

 func main(){
	cfg:= config.MustLoad()
	router:= http.NewServeMux()

	router.HandleFunc("POST /api/students",student.Create)

	server:=http.Server{
		Addr: cfg.HTTPServer.Addr,
		Handler: router,
	}
	slog.Info("server started", slog.String("address", cfg.Addr))
	done:= make(chan os.Signal,1)
	signal.Notify(done,os.Interrupt,syscall.SIGINT, syscall.SIGTERM)
	go func ()  {
		err:= server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
        log.Fatal("Failed to start the server: ", err.Error())
    }
	}()

	<-done
	slog.Info("Shutting down the server")

	ctx, cancel:= context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err:= server.Shutdown(ctx); err!=nil{
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfully")	
 }