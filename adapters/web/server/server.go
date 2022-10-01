package server

import (
	"github.com/filipedtristao/hexagonal-architecture/application"
	"github.com/filipedtristao/hexagonal-architecture/adapters/web/handler"
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"os"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func NewWebserver(service application.ProductServiceInterface) *Webserver {
	return &Webserver{service}
}

func (w *Webserver) Start() {
	router := mux.NewRouter()
	negroni := negroni.New(negroni.NewLogger())

	handler.MakeProductHandler(router, negroni, w.Service)
	http.Handle("/", router)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr: ":9000",
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr, "ERROR ", log.LstdFlags),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}