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

type favrouteRouter struct {
	favrouteService root.FavrouteService
}

func NewFavrouteRouter(v root.FavrouteService, router *mux.Router) *mux.Router {
	favrouteRouter := favrouteRouter{v}

	router.HandleFunc("/fav_routes/", favrouteRouter.createFavrouteHandler).Methods("POST")
	router.HandleFunc("/fav_routes/", favrouteRouter.getAllFavrouteHandler).Methods("GET")
  router.HandleFunc("/fav_routes/{id}/", favrouteRouter.getFavrouteByIdHandler).Methods("GET")
	router.HandleFunc("/fav_routes/my_favRoutes/{user_id}/", favrouteRouter.getFavrouteByUserIdHandler).Methods("GET")
	router.HandleFunc("/fav_routes/{id}/", favrouteRouter.deleteFavrouteHandler).Methods("DELETE")
	router.HandleFunc("/fav_routes/{id}/", favrouteRouter.updateFavrouteHandler).Methods("PUT")
	return router
}

func (vr *favrouteRouter) createFavrouteHandler(w http.ResponseWriter, r *http.Request) {
	favroute, err := decodeFavroute(r)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
  salida, err := vr.favrouteService.CreateFavroute(&favroute)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, http.StatusOK, salida)
}

func (vr *favrouteRouter) deleteFavrouteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	favrouteId, err := strconv.Atoi(vars["id"])

	err = vr.favrouteService.DeleteById(favrouteId)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, err)
}

func (vr *favrouteRouter) getFavrouteByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	favrouteId, err := strconv.Atoi(vars["id"])
	favroute, err := vr.favrouteService.GetById(favrouteId)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, favroute)
}

func (vr *favrouteRouter) updateFavrouteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	favrouteId, err := strconv.Atoi(vars["id"])
	favroute, err := decodeFavroute(r)

	favrouteUp, err := vr.favrouteService.UpdateById(favrouteId, &favroute)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, favrouteUp)
}

func (vr *favrouteRouter) getAllFavrouteHandler(w http.ResponseWriter, r *http.Request) {
	favroutes, err := vr.favrouteService.GetAll()
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, favroutes)

}

func (vr *favrouteRouter) getFavrouteByUserIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	favrouteId, err := strconv.Atoi(vars["user_id"])
	favroute, err := vr.favrouteService.GetByUserId(favrouteId)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, favroute)
}


func decodeFavroute(r *http.Request) (root.Favroute, error) {
  var v root.Favroute
  if r.Body == nil {
    return v, errors.New("no request body")
  }
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&v)
  return v, err
}
