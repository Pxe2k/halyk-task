package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Pxe2k/halyk-task/internal/delivery/show"
	"github.com/Pxe2k/halyk-task/internal/repository"
	show2 "github.com/Pxe2k/halyk-task/internal/usecase/show"

	"github.com/Pxe2k/halyk-task/pkg"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	db, err := pkg.NewMySQLDB(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	srv := http.Server{
		Addr: fmt.Sprint(":", os.Getenv("PORT")),
		Handler: show.New(
			show2.New(
				db,
				repository.NewShowRepo(db),
				repository.NewSeatRepo(db),
			),
		),
	}

	log.Println("server started...")

	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
