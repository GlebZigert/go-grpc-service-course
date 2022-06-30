package main

import "fmt"
import "log"

func Run() error {

	fmt.Println("Rocket Service Starting...")
	return nil
}

func main(){
if err:=Run();err != nil {
	log.Fatal(err)
}
}