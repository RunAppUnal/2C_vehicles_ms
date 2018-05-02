package mongo

import (
  "2C_vehicle_ms/pkg"
  //"gopkg.in/mgo.v2/bson"
  "gopkg.in/mgo.v2"
  //"golang.org/x/crypto/bcrypt"
)

type vehicleModel struct {
  ID          uint64         `bson:"_id" json:"_id"`//bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
  VehicleId          uint64         `bson:"id" json:"id"`
  //VehicleId   uint64        `bson:"vehicleid" json:"vehicleid"`
	Plate       string        `bson:"plate" json:"plate"`
	UserId      uint64        `bson:"user_id" json:"user_id"`
	Kind        string        `bson:"kind" json:"kind"`
  Model       uint64        `bson:"model" json:"model"`
  Color       string        `bson:"color" json:"color"`
  Capacity    uint64        `bson:"capacity" json:"capacity"`
  Image       string        `bson:"image" json:"image"`
  Brand        string       `bson:"brand" json:"brand"`
}

func vehicleModelIndex() mgo.Index {
  return mgo.Index{
    Key:        []string{"plate"},
    Unique:     true,
    DropDups:   true,
    Background: true,
    Sparse:     true,
  }
}

func newVehicleModel(v *root.Vehicle) *vehicleModel {
  return &vehicleModel{
    VehicleId:   v.VehicleId,
    Plate:       v.Plate,
  	UserId:      v.UserId,
  	Kind:        v.Kind,
    Model:       v.Model,
    Color:       v.Color,
    Capacity:    v.Capacity,
    Image:       v.Image,
    Brand:       v.Brand}
}

func updateVehicleModel(v *root.Vehicle) *vehicleModel {
  return &vehicleModel{
    ID:          v.ID,
    VehicleId:   v.VehicleId,
    Plate:       v.Plate,
  	UserId:      v.UserId,
  	Kind:        v.Kind,
    Model:       v.Model,
    Color:       v.Color,
    Capacity:    v.Capacity,
    Image:       v.Image,
    Brand:       v.Brand}
}

func(v *vehicleModel) toRootVehicle() *root.Vehicle {
  return &root.Vehicle{
    ID:          v.ID,
    VehicleId:   v.VehicleId,
    Plate:       v.Plate,
  	UserId:      v.UserId,
  	Kind:        v.Kind,
    Model:       v.Model,
    Color:       v.Color,
    Capacity:    v.Capacity,
    Image:       v.Image,
    Brand:       v.Brand}
}
