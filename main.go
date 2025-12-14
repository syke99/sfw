package sfw

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/syke99/sfw/app/web"
)

var (
	path     = ""
	pathFlag = flag.String("path", "", "path to look for files")
)

func init() {
	flag.Parse()
}

func checkPath(path string, calls int) error {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("path does not exist")
	}
	if err != nil {
		fmt.Println("Error occurred:", err)
		return fmt.Errorf("error checking path")
	}

	if !info.IsDir() {
		return fmt.Errorf("path is not a directory")
	}

	if base := filepath.Base(path); base != "spiderweb" {
		if calls == 0 {
			return checkPath(path, 1)
		}

		return fmt.Errorf("path does not end in or have a direct child directory named spiderweb")
	}

	return nil
}

func main() {
	if *pathFlag == "" {
		wd, err := os.Getwd()
		if err != nil {
			// TODO: handle fatal error here better
			log.Fatal(err)
		}

		path = wd
	}

	if err := checkPath(*pathFlag, 0); err != nil {
		// TODO: handle fatal error here better
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	// TODO: configure port to run on
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	spiderWeb, err := web.NewWeb(mux, path)
	if err != nil {
		// TODO: handle err shutdown better
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		err = spiderWeb.Cast(ctx)
		if err != nil {
			os.Exit(1)
			return
		}
	}()

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// TODO: handle server shutdown err better
			log.Fatalf("server error: %v", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	log.Println("Shutting down...")

	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown error: %v", err)
	}
}
