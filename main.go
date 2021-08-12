package main

import (
	controller "github.com/felipefrm/go-sistema-vendas/controller"
	view "github.com/felipefrm/go-sistema-vendas/view"
)

func main() {
	mdc := controller.MainDaoController{View: view.MainView{}}
	mdc.OptionsMenu()
}
