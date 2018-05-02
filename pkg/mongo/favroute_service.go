package mongo

import (
	"2C_vehicle_ms/pkg"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/night-codes/mgo-ai"
	//"reflect"
)

type FavrouteService struct {
	collection *mgo.Collection
	//hash       root.Hash
}

func NewFavrouteService(session *Session, dbName string, collectionName string) *FavrouteService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(favrouteModelIndex())
	return &FavrouteService{collection}
}

func (p *FavrouteService) CreateFavroute(v *root.Favroute) (*root.Favroute, error) {
	favroute := newFavrouteModel(v)

	// connect AutoIncrement to collection "vehicles"
	session, err := mgo.Dial("192.168.99.101:27017")
	if err != nil {
		log.Fatal(err)
	}
  ai.Connect(session.DB("2C_vehicle_db").C("favroute"))
	//favroute.FavrouteId = ai.Next("favroute")
	favroute.ID = ai.Next("favroute")
	favroute.FavroutesId = favroute.ID
	err = p.collection.Insert(&favroute)
	return favroute.toRootFavroute(), err
}

func (p *FavrouteService) DeleteById(id int)  error{
	err := p.collection.Remove(bson.M{"_id": id})
	return  err
}

func (p *FavrouteService) GetById(id int) (*root.Favroute, error) {
	model := favrouteModel{}
	err := p.collection.Find(bson.M{"_id": id}).One(&model)
	return model.toRootFavroute(), err
}

func (p *FavrouteService) GetAll() ([]*root.Favroute, error){
	model := []favrouteModel{}
	var salida []*root.Favroute
  err := p.collection.Find(nil).All(&model)
	aux := favrouteModel{}
	if len(model) > 0{
		for i := 1; i < len(model); i++{
			aux = model[i]
			salida = append(salida, aux.toRootFavroute())
		}
	}
	return salida, err
}

func (p *FavrouteService) GetByUserId(userid int) ([]*root.Favroute, error){
	model := []favrouteModel{}
	var salida []*root.Favroute
  err := p.collection.Find(bson.M{"user_id": userid}).All(&model)
	aux := favrouteModel{}
	log.Println("------------->")
	log.Println(userid, model)
	if len(model) > 0{
		for i := 0; i < len(model); i++{
			aux = model[i]
			salida = append(salida, aux.toRootFavroute())
		}
	}
	return salida, err
}

func (p *FavrouteService) UpdateById(id int, v *root.Favroute) (*root.Favroute, error){
	favroute := updateFavrouteModel(v)
	favroute.ID = uint64(id)
	favroute.FavrouteId = uint64(id)
	err := p.collection.UpdateId(id, &favroute)
	return favroute.toRootFavroute(), err
}
