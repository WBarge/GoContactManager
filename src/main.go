package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/WBarge/GoContactManager/application"
)

func main() {

	configuration := application.NewConfig()

	//todo inject core peices and the persistance store peices
	webServer := application.NewWebServer(*configuration)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := webServer.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
