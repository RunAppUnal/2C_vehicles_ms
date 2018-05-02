package mongo

import (
  "2C_vehicle_ms/pkg"
  //"gopkg.in/mgo.v2/bson"
  "gopkg.in/mgo.v2"
  //"golang.org/x/crypto/bcrypt"
)

type favrouteModel struct {
  ID          uint64        `bson:"_id" json:"_id"`//bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
  FavroutesId uint64        `bson:"id" json:"id"`
	UserId      uint64        `bson:"user_id" json:"user_id"`
	Polyline1   string        `bson:"polyline1" json:"polyline1"`
  Polyline2   string        `bson:"polyline2" json:"polyline2"`
  Polyline3   string        `bson:"polyline3" json:"polyline3"`
  Polyline4   string        `bson:"polyline4" json:"polyline4"`
  Polyline5   string        `bson:"polyline5" json:"polyline5"`

}

func favrouteModelIndex() mgo.Index {
  return mgo.Index{
    Key:        []string{"_id"},
    Unique:     true,
    DropDups:   true,
    Background: true,
    Sparse:     true,
  }
}

func newFavrouteModel(f *root.Favroute) *favrouteModel {
  return &favrouteModel{
    FavroutesId:   f.FavroutesId,
  	UserId:        f.UserId,
  	Polyline1:     f.Polyline1,
    Polyline2:     f.Polyline2,
    Polyline3:     f.Polyline3,
    Polyline4:     f.Polyline4,
    Polyline5:     f.Polyline5}
}

func updateFavrouteModel(f *root.Favroute) *favrouteModel {
  return &favrouteModel{
    FavroutesId:   f.FavroutesId,
  	UserId:        f.UserId,
  	Polyline1:     f.Polyline1,
    Polyline2:     f.Polyline2,
    Polyline3:     f.Polyline3,
    Polyline4:     f.Polyline4,
    Polyline5:     f.Polyline5}
}

func(f *favrouteModel) toRootFavroute() *root.Favroute {
  return &root.Favroute{
    ID:            f.ID,
    FavroutesId:   f.FavroutesId,
  	UserId:        f.UserId,
  	Polyline1:     f.Polyline1,
    Polyline2:     f.Polyline2,
    Polyline3:     f.Polyline3,
    Polyline4:     f.Polyline4,
    Polyline5:     f.Polyline5}
}
