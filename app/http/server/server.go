package server

import (
	"context"
	"fmt"
	"hello-world-api/app/config"
	"hello-world-api/app/http/router"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Run(cfg *config.Config) *http.Server {

	r := router.Init()

	srv := &http.Server{
		Addr: "0.0.0.0:" + strconv.Itoa(cfg.AppConfig.Port),

		// good practice to set timeouts to avoid Slowloris attacks
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 10,

		// pass our instance of gorilla/mux in
		Handler: r,
	}

	// run our server in a goroutine so that it doesn't block
	go func() {

		err := srv.ListenAndServe()
		if err != nil {
			log.Println(err)
			panic("Service shutting down unexpectedly...")
		}
	}()

	fmt.Println("Service started.........")
	fmt.Println(fmt.Sprintf("Listening on %v ........", srv.Addr))

	return srv
}

func Stop(ctx context.Context, srv *http.Server) {

	fmt.Println("Service shutting down...")

	_ = srv.Shutdown(ctx)

}
