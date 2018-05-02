package server

import (
	"encoding/json"
	"errors"
	"2C_vehicle_ms/pkg"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type vehicleRouter struct {
	vehicleService root.VehicleService
}

func NewVehicleRouter(v root.VehicleService, router *mux.Router) *mux.Router {
	vehicleRouter := vehicleRouter{v}

	router.HandleFunc("", vehicleRouter.createVehicleHandler).Methods("POST")
	router.HandleFunc("", vehicleRouter.getAllVehicleHandler).Methods("GET")
  router.HandleFunc("/{id}", vehicleRouter.getVehicleByIdHandler).Methods("GET")
	router.HandleFunc("/find_vehicle/{plate}", vehicleRouter.getVehicleByPlateHandler).Methods("GET")
	router.HandleFunc("/my_vehicles/{user_id}", vehicleRouter.getVehicleByUserIdHandler).Methods("GET")
	router.HandleFunc("/{id}", vehicleRouter.deleteVehicleHandler).Methods("DELETE")
	router.HandleFunc("/{id}", vehicleRouter.updateVehicleHandler).Methods("PUT")
	return router
}

func (vr *vehicleRouter) createVehicleHandler(w http.ResponseWriter, r *http.Request) {
	vehicle, err := decodeVehicle(r)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
  err = vr.vehicleService.CreateVehicle(&vehicle)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, http.StatusOK, err)
}

func (vr *vehicleRouter) deleteVehicleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	vehicleId, err := strconv.Atoi(vars["id"])

	err = vr.vehicleService.DeleteById(vehicleId)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, err)
}

func (vr *vehicleRouter) getVehicleByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vehicleId, err := strconv.Atoi(vars["id"])
	vehicle, err := vr.vehicleService.GetById(vehicleId)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, vehicle)
}

func (vr *vehicleRouter) updateVehicleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vehicleId, err := strconv.Atoi(vars["id"])
	vehicle, err := decodeVehicle(r)

	vehicleUp, err := vr.vehicleService.UpdateById(vehicleId, &vehicle)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, vehicleUp)
}

func (vr *vehicleRouter) getAllVehicleHandler(w http.ResponseWriter, r *http.Request) {
	vehicles, err := vr.vehicleService.GetAll()
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, vehicles)
}

func (vr *vehicleRouter) getVehicleByPlateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vehicle, err := vr.vehicleService.GetByPlate(vars["plate"])
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, vehicle)
}

func (vr *vehicleRouter) getVehicleByUserIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vehicleId, err := strconv.Atoi(vars["user_id"])
	vehicle, err := vr.vehicleService.GetByUserId(vehicleId)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, vehicle)
}


func decodeVehicle(r *http.Request) (root.Vehicle, error) {
  var v root.Vehicle
  if r.Body == nil {
    return v, errors.New("no request body")
  }
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&v)
  return v, err
}
