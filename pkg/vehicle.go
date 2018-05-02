package root

type Vehicle struct {
  ID          uint64        `json:"_id"`
  //VehicleId   uint64        `json:"vehicleid"`
	Plate       string        `json:"plate"`
	UserId      uint64        `json:"user_id"`
	Kind        string        `json:"kind"`
  Model       uint64        `json:"model"`
  Color       string        `json:"color"`
  Capacity    uint64        `json:"capacity"`
  Image       string        `json:"image"`
  Brand       string        `json:"brand"`
}

type VehicleService interface {
  CreateVehicle(v *Vehicle) error
  GetById(vehicleid int) (*Vehicle,error)
  GetByPlate(plate string) (*Vehicle,error)
  DeleteById(vehicleid int) error
  UpdateById(vehicleid int, v *Vehicle) (*Vehicle,error)
  GetAll() ([]*Vehicle,error)
  GetByUserId(userid int) ([]*Vehicle,error)
}
