package main

import "shop/app"

func main() {
	application := app.New()
	application.InitRoutes()
	application.Start()
}
