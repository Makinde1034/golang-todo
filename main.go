package main

import (

	controller "get-going/controllers"
	  "get-going/models"
	  "log"
	  "net/http"
    _ "github.com/go-sql-driver/mysql"
)

type user struct {
  name string
  age int
  verified bool
}

func (u user) getNameAndAge() (string,int) {  
  return u.name , u.age
}


func main() {
  mux := controller.StartServer()
  db := models.Connect()
  defer db.Close()
  log.Fatal(http.ListenAndServe("localhost:4000",mux))

}