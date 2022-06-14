package main

import (
	"context"
	"github.com/BlackRRR/payment-emulator/internal/config"
	"github.com/BlackRRR/payment-emulator/internal/repository"
	"github.com/BlackRRR/payment-emulator/internal/server"
	"github.com/BlackRRR/payment-emulator/internal/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.TODO()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed to get config: %s", err.Error())
	}

	log.Println("init config success")

	db, err := repository.InitDataBase(ctx, cfg.RepositoryCfg)
	if err != nil {
		log.Fatalf("failed to init database: %s", err.Error())
	}

	log.Println("init database success")

	repositories, err := repository.InitRepositories(ctx, db)
	if err != nil {
		log.Fatalf("failed to init repositories: %s", err.Error())
	}

	log.Println("init repository success")

	svc := services.InitAllServices(repositories)

	log.Println("init services success")

	httpHandler := server.MakeHTTPHandler(server.InitServer(svc))

	go func() {

		log.Println("http server started on port:", cfg.ServicePort)
		serviceErr := http.ListenAndServe(":"+cfg.ServicePort, httpHandler)
		if serviceErr != nil {
			log.Fatalf("http handler was stoped by err: %s", serviceErr.Error())
		}
	}()

	sig := <-subscribeToSystemSignals()

	log.Printf("shutdown all process on '%s' system signal\n", sig.String())

}

func subscribeToSystemSignals() chan os.Signal {
	ch := make(chan os.Signal, 10)
	signal.Notify(ch,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGHUP,
	)
	return ch
}
