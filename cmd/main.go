package main

import (
	"github.com/daniil-sargsyan/go-basic/pkg/app"
	"github.com/daniil-sargsyan/go-basic/pkg/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func main() {
	var configPath string
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}
	cnf, err := config.New(configPath)
	if err != nil {
		log.Fatal(err)
	}
	application := app.New(cnf)

	go func() {
		log.Println(application.Run())
	}()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case <-quitCh:
			err = application.GracefulStop()
			if err != nil {
				log.Fatalln(err)
			}
			os.Exit(1)
		}
	}
}
