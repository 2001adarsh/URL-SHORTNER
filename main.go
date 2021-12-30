package main

import (
	"github.com/2001adarsh/url-shortner/app"
)

func main() {
	application := app.App{}
	application.Initialization()
	application.Run()
}
