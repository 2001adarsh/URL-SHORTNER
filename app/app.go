package app

import (
	"context"
	"github.com/2001adarsh/url-shortner/app/handler"
	"github.com/2001adarsh/url-shortner/app/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	Router *mux.Router
}

var (
	db    storage.Database
	logOp = log.New(os.Stdout, "URL SHORTNER", log.LstdFlags)
)

func (app *App) Initialization() {
	app.initializeDB()
	app.initializeRoutes()
}

func (app *App) Run() {
	customServer := &http.Server{
		Handler:      app.Router,
		Addr:         ":9000",
		ErrorLog:     logOp,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	logOp.Println("[INFO] SERVER STARTUP")
	go func() {
		logOp.Fatal(customServer.ListenAndServe())
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	//Block until a signal is received.
	sig := <-c
	logOp.Println("Got Signal:", sig, ". Hence Gracefully closing down.")

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	cntx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	customServer.Shutdown(cntx)
}

func (app *App) initializeRoutes() {
	app.Router = mux.NewRouter()
	shortUrlHandler := handler.NewShortURL(logOp, db)

	createRouter := app.Router.Path("/create").Subrouter()
	createRouter.Methods(http.MethodPost, http.MethodPut).HandlerFunc(shortUrlHandler.CreateHandler)

	urlRouter := app.Router.PathPrefix("/{url}").Subrouter()
	urlRouter.Path("").Methods(http.MethodGet).HandlerFunc(shortUrlHandler.MapHandler)
	urlRouter.Path("/view").Methods(http.MethodGet).HandlerFunc(shortUrlHandler.ViewHandler)
	urlRouter.Use(shortUrlHandler.ValidateMiddleware)
}

func (app *App) initializeDB() {
	db = storage.NewRedisDB(logOp, "localhost:6379", 10*time.Hour)
}
