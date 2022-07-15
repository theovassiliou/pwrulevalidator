package main

import "github.com/gorilla/mux"

func main() {
	a := App{}
	a.Initialize()
	a.Run(":8010")
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
}
