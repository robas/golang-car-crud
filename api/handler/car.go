package handler

import (
	"carProject/api/presenter"
	"carProject/entity"
	"carProject/usecase/car"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func createCar(service car.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding car"
		var input struct {
			Brand        string `json:"brand"`
			Model        string `json:"model"`
			DoorQuantity int    `json:"doorQuantity"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.CreateCar(input.Brand, input.Model, input.DoorQuantity)

		fmt.Println(id)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Car{
			ID:           id,
			Brand:        input.Brand,
			Model:        input.Model,
			DoorQuantity: input.DoorQuantity,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getCar(service car.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error getting car"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.GetCar(id)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Car{
			ID:           data.ID,
			Brand:        data.Brand,
			Model:        data.Model,
			DoorQuantity: data.DoorQuantity,
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func listCars(service car.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading cars"
		var data []*entity.Car
		var err error
		brand := r.URL.Query().Get("brand")

		switch {
		case brand == "":
			data, err = service.ListCars()
		default:
			data, err = service.SearchCars(brand)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		var toJ []*presenter.Car
		for _, c := range data {
			toJ = append(toJ, &presenter.Car{
				ID:           c.ID,
				Brand:        c.Brand,
				Model:        c.Model,
				DoorQuantity: c.DoorQuantity,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func MakeCarHandlers(r *mux.Router, service car.UseCase) {
	r.Handle("/v1/car", createCar(service)).
		Methods("POST", "OPTIONS").Name("createCar")

	r.Handle("/v1/car/{id}", getCar(service)).
		Methods("GET", "OPTIONS").Name("getCar")

	r.Handle("/v1/car", listCars(service)).
		Methods("GET", "OPTIONS").Name("listCars")
}
