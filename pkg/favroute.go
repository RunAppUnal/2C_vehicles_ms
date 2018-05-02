package root

type Favroute struct {
  ID          uint64        `json:"_id"`
  FavroutesId uint64        `json:"id"`
	UserId      uint64        `json:"user_id"`
	Polyline1   string        `json:"polyline1"`
  Polyline2   string        `json:"polyline2"`
  Polyline3   string        `json:"polyline3"`
  Polyline4   string        `json:"polyline4"`
  Polyline5   string        `json:"polyline5"`
}

type FavrouteService interface {
  CreateFavroute(f *Favroute) (*Favroute,error)
  GetById(favrouteid int) (*Favroute,error)
  DeleteById(favrouteid int) error
  UpdateById(favrouteid int, v *Favroute) (*Favroute,error)
  GetAll() ([]*Favroute,error)
  GetByUserId(favrouteid int) ([]*Favroute,error)
}
