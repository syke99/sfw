package sfw

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/syke99/sfw/app/web"
	"log"
	"os"
	"path/filepath"
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

	mux := chi.NewRouter()

	_, err := web.NewWeb(mux, path)
	if err != nil {
		// TODO: handle err shutdown better
		log.Fatal(err)
	}
}
