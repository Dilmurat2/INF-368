package main

import "assingment4/internal/app"

func main() {
	app, err := app.AppInit()
	if err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		panic(err)
	}
}
