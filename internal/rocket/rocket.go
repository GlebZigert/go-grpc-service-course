//go:generate /home/administrator/go/bin/mockgen -destination=./rocket_mocks_test.go -package=rocket github.com/GlebZigert/go-grpc-service-course/internal/rocket Store

//go generate ./...
package rocket

import "context"

// Rocket - should contain things like
// ID of the rocket,
// the name of the rocket
// the type fo the rocket

type Rocket struct{

	ID string
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

func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {

	rkt,err := s.Store.GetRocketByID(id)
	if  err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

func (s Service) AddRocket(ctx context.Context,rkt Rocket) (Rocket, error){

	rkt,err := s.Store.InsertRocket(rkt)
	if  err != nil {
		return Rocket{}, err
	}

	return rkt, nil	
}

func (s Service) DeleteRocket(ctx context.Context,id string) error {

	err := s.Store.DeleteRocket(id)
	if err != nil {

		return err
	}

	return nil
}

