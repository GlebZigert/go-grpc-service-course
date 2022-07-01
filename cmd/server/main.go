package main



import (
	
	"github.com/GlebZigert/go-grpc-services-course/internal/db"
	"github.com/GlebZigert/go-grpc-services-course/internal/rocket"	
	   "fmt"
	   "log"
)

func Run() error {

	fmt.Println("Rocket Service Starting...")

	rocketStore, err:=db.New()
	if err != nil {
		return err	
	}

	err = rocketStore.Migrate()

	if err != nil {

		log.Println("Failed to run migrations")
		return err
	}



	fmt.Println("db connection success")
	_ = rocket.New(rocketStore)
	
	return nil
}



func main(){
if err:=Run();err != nil {
	log.Fatal(err)
}
}