package main

import (
	"carProject/api/handler"
	"carProject/config"
	"carProject/infrastructure/repository"
	"carProject/usecase/car"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	var err error
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	db, err := sql.Open("mysql", datasourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	carRepo := repository.NewCarMySQL(db)
	//carRepo := repository.NewInmem()
	carService := car.NewService(carRepo)

	r := mux.NewRouter()
	handler.MakeCarHandlers(r, carService)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     logger,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
