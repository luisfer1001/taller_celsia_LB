package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/test/pkg/db"
	"github.com/test/pkg/formularios"
	"github.com/test/pkg/imagenes"
	"github.com/test/pkg/res"
)

func main() {

	err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	err = CrearteTables()
	if err != nil {
		log.Fatal(err)
	}

	err = HttpServer()
	if err != nil {
		log.Fatal(err)
	}

}

func CrearteTables() error {
	log.Println("CrearteTables")

	err := formularios.CreateTables()
	if err != nil {
		return err
	}
	err = imagenes.CreateTables()
	if err != nil {
		return err
	}
	return nil
}

func HttpServer() error {
	log.Println("HttpServer")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		res.JSON(w, r, http.StatusOK, res.Json{"message": "pong"})
	})

	formularios.Router(r)
	imagenes.Router(r)

	hadler := cors.AllowAll().Handler(r)
	addr := ":3000"
	server := &http.Server{
		Addr:    addr,
		Handler: hadler,
	}

	log.Println("Run http server en port:3000")
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

/*
// este es un ejemplo del manejo de funciones
func Suma(a int, b int) (int, error) {
	var c int
	var err error
	a = 4
	b = 3
	c = a + b
	err = nil
	return c, err
}

func Resta(a int, b int) (int, error) {
	var r int
	var err error
	r = a - b
	err = nil
	return r, err
}

func Operaciones(a int, b int) (int, int, error) {
	var suma int
	var resta int
	var err error
	suma = a + b
	resta = a - b
	return suma, resta, err
}
*/
