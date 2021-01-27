package main

import (
	"candy/observability"
	"candy/pubsub"
	"candy/server"
)

func main() {
	logger := observability.NewLogger(observability.Info)

	pubSubRemote := pubsub.NewRemote(&logger)

	gameServer := server.NewGame(&logger, pubSubRemote)

	go func() {
		server.WaitReady("localhost", 8081)
		pubSubRemote.Start("localhost", 8081)
	}()

	err := gameServer.Start(8082)
	if err != nil {
		panic(err)
	}
}
