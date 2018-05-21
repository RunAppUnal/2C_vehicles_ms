package main

import (
	"2C_vehicle_ms/pkg/mongo"
	"2C_vehicle_ms/pkg/server"
	"log"
)

func main() {
	ms, err := mongo.NewSession("vehicles-db:7005")
	if err != nil {
		log.Println(err)
		log.Fatalln("unable to connect to mongodb")
	}
	defer ms.Close()

	//h := crypto.Hash{}
	u := mongo.NewVehicleService(ms.Copy(), "2C_vehicle_db", "vehicle")
	//s := server.NewServerVehicle(u)


	uf := mongo.NewFavrouteService(ms.Copy(), "2C_vehicle_db", "favroute")
	//sf := server.NewServerFavroute(uf)

	s := server.NewServer(u, uf)

	s.Start()
	//log.Println("--------++++++++++")
	//sf.Start()

}
