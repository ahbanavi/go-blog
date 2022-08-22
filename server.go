package main

import (
	"fmt"
	"log"

	"github.com/ahbanavi/go-blog/bootstrap"
	"github.com/ahbanavi/go-blog/config"
)

func main() {
	app := bootstrap.NewApp()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", config.GetD("APP_HOST", "localhost"), config.GetD("APP_PORT", "3000"))))
}
