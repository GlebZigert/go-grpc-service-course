//go:generate /home/administrator/go/bin/mockgen -package=rocket github.com/GlebZigert/go-grpc-service-course/internal/rocket Store

package rocket

// Rocket - should contain things like
// ID of the rocket,
// the name of the rocket
// the type fo the rocket

type Rocket struct{

	ID int
	Name string
	Type string
}

//Service - our rocket service, used for updating our rocket 
//inventory

type Service struct{

	Store Store
}

//New -return new rocket service
func New(store Store) Service{

	return Service{

		Store: store,
	}
}

//Store - defines the interface we need to satisfy for our
//service to work correctly

type Store interface{

	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

func (s Service) GetRocketByID(id string) (Rocket, error) {

	rkt,err := s.Store.GetRocketByID(id)
	if  err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

func (s Service) AddRocket(rkt Rocket) (Rocket, error){

	rkt,err := s.Store.InsertRocket(rkt)
	if  err != nil {
		return Rocket{}, err
	}

	return rkt, nil	
}

func (s Service) DeleteRocket(id string) error {

	err := s.Store.DeleteRocket(id)
	if err != nil {

		return err
	}

	return nil
}

