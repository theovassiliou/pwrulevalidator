package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	a := App{}
	a.Initialize()
	a.Run(":8010")
	fmt.Println("HELLO THEO")
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
}
