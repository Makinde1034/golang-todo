package controller

import (
	"net/http"
)

func StartServer() *http.ServeMux{
	mux := http.NewServeMux() 
	mux.HandleFunc("/ping",ping())
	mux.HandleFunc("/create",create())
	mux.HandleFunc("/getAll",getAllTodos())
	mux.HandleFunc("/getSingle",getSingleTodo())
	mux.HandleFunc("/deleted",DeleteTodo())
  
	return mux
}