package mongo

import (
	"2C_vehicle_ms/pkg"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/night-codes/mgo-ai"
	//"reflect"
)

type VehicleService struct {
	collection *mgo.Collection
	//hash       root.Hash
}

func NewVehicleService(session *Session, dbName string, collectionName string) *VehicleService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(vehicleModelIndex())
	return &VehicleService{collection}
}

func (p *VehicleService) CreateVehicle(v *root.Vehicle) (*root.Vehicle, error) {
	vehicle := newVehicleModel(v)

	// connect AutoIncrement to collection "vehicles"
	session, err := mgo.Dial("192.168.99.101:27017")
	if err != nil {
		log.Fatal(err)
	}
  ai.Connect(session.DB("2C_vehicle_db").C("vehicle"))
	//vehicle.VehicleId = ai.Next("vehicle")
	vehicle.ID = ai.Next("vehicle")
	vehicle.VehicleId = vehicle.ID
	err = p.collection.Insert(&vehicle)
	return vehicle.toRootVehicle(), err
}

func (p *VehicleService) DeleteById(id int)  error{
	err := p.collection.Remove(bson.M{"_id": id})
	return  err
}

func (p *VehicleService) GetById(id int) (*root.Vehicle, error) {
	model := vehicleModel{}
	err := p.collection.Find(bson.M{"_id": id}).One(&model)
	return model.toRootVehicle(), err
}

func (p *VehicleService) GetAll() ([]*root.Vehicle, error){
	model := []vehicleModel{}
	var salida []*root.Vehicle
  err := p.collection.Find(nil).All(&model)
	aux := vehicleModel{}
	if len(model) > 0{
		for i := 1; i < len(model); i++{
			aux = model[i]
			salida = append(salida, aux.toRootVehicle())
		}
	}
	return salida, err
}

func (p *VehicleService) GetByPlate(plate string) ([]*root.Vehicle, error) {
	model := []vehicleModel{}
	var salida []*root.Vehicle
  err := p.collection.Find(bson.M{"plate": plate}).All(&model)
	aux := vehicleModel{}
	log.Println("------------->")
	log.Println(plate, model)
	if len(model) > 0{
		for i := 0; i < len(model); i++{
			aux = model[i]
			salida = append(salida, aux.toRootVehicle())
		}
	}

	return salida, err

}

func (p *VehicleService) GetByUserId(userid int) ([]*root.Vehicle, error){
	model := []vehicleModel{}
	var salida []*root.Vehicle
  err := p.collection.Find(bson.M{"user_id": userid}).All(&model)
	aux := vehicleModel{}
	log.Println("------------->")
	log.Println(userid, model)
	if len(model) > 0{
		for i := 0; i < len(model); i++{
			aux = model[i]
			salida = append(salida, aux.toRootVehicle())
		}
	}
	return salida, err
}



func (p *VehicleService) UpdateById(id int, v *root.Vehicle) (*root.Vehicle, error){
	vehicle := updateVehicleModel(v)
	vehicle.ID = uint64(id)
	err := p.collection.UpdateId(id, &vehicle)
	return vehicle.toRootVehicle(), err
}
