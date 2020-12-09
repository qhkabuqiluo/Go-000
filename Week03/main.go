package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx0 := errgroup.WithContext(ctx)
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-sigs:
			log.Println("receive quit signal")
			cancel()
			return errors.New("receive quit signal")
		case <-ctx0.Done():
			return nil
		}
	})
	for i := 8010; i < 8020; i++ {
		addr := ":" + strconv.Itoa(i)
		g.Go(func() error {
			return Server(ctx0, addr)
		})
	}
	err := g.Wait()
	if errors.Is(err, context.Canceled) {
		log.Printf("%v", err)
	} else if err != nil {
		log.Printf("errorgroup error: %s\n", err)
	}
	log.Printf("ctx0 error: %s", ctx0.Err())
}

func Server(ctx context.Context, addr string) error {
	server := &http.Server{Addr: addr}
	go func() {
		select {
		case <-ctx.Done():
			ctx1, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			log.Printf("shutdown server , addr is %s", addr)
			server.Shutdown(ctx1)
		}
	}()
	log.Printf("%s server is starting", addr)
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			err = nil
		} else {
			log.Printf("%s  server started failed, err is %s", addr, err)
		}
	}
	return err
}
