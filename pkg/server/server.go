package server

import (
	"2C_vehicle_ms/pkg"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer(v root.VehicleService) *Server {
	s := Server{router: mux.NewRouter()}
	NewVehicleRouter(v, s.newSubrouter("/vehicles"))
	return &s
}

func (s *Server) Start() {
	log.Println("Listening on port 6005")
	if err := http.ListenAndServe(":6005", handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) newSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
