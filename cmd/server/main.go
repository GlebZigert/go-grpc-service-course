package main

import "fmt"
import "log"

func Run() error {

	fmt.Println("Rocket Service Starting...")
/*
	rocketStore, err:=db.New()
	if err != nil {
		return err	
	}

	_ = rocket.New(rocketStore)
	*/
	return nil
}

func main(){
if err:=Run();err != nil {
	log.Fatal(err)
}
}