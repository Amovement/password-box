package main

import (
	"log"

	"github.com/Amovement/password-box/pkg/router"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	oauthServer := router.SetupRouterAndGetServer()
	g.Go(func() error {
		return oauthServer.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
