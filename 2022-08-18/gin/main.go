package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.RedirectTrailingSlash = false
	r.Use(func() func(gCtx *gin.Context) {

		file, err := os.OpenFile("/tmp/tmo.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if nil != err {
			fmt.Println(err)
		}
		fmt.Println(file.Fd())
		// defer func() {
		// 	fmt.Println("1")
		// 	fmt.Println(file.Close())
		// }()

		return func(gCtx *gin.Context) {

			fmt.Println("2")
			fmt.Println(file.Write([]byte("123\n")))
			gCtx.Next()
			fmt.Println("3")

		}

	}())
	r.GET("/test/name", func(gCtx *gin.Context) {
		gCtx.JSON(http.StatusOK, gin.H{
			"label":  "/test/name",
			"name":   gCtx.Param("name"),
			"active": gCtx.Param("active"),
		})
	})
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
