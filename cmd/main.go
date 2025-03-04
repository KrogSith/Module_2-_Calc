package main

import "calculator/pkg/application"

func main() {
	app := application.New()
	//app.Run()
	app.RunServer()
}
