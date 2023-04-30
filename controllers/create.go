package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"get-going/models"
	"get-going/structs"
)

func create() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodPost {
			data := structs.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			if err := models.CreateTodo(data.Name,data.Todo); err != nil{
				w.Write([]byte("An error occured"))
			}
			w.Write([]byte("Create todo"))
			json.NewEncoder(w).Encode(data)
		}
	} 
}

func getAllTodos() http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		data,err := models.ReadAll()
		if err != nil {
			w.Write([]byte(err.Error()))
		}
	
		json.NewEncoder(w).Encode(data)
	}
}

func getSingleTodo() http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		data := structs.GetTodoRequest{}
		json.NewDecoder(r.Body).Decode(&data)

		todo,err := models.ReadByName(data.Name)
		fmt.Println(todo)
		if err != nil {
			w.Write([]byte("An error occured"))
		}

		json.NewEncoder(w).Encode(todo)
	}
}

func DeleteTodo() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {         
		if r.Method == http.MethodDelete{
			data := structs.GetTodoRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			if err := models.DeleteTodo(data.Name); err !=nil{
				w.Write([]byte("An error occured"))
			}
			res := struct{
				Msg string `json:"msg"`
			}{"Successfully deleted todo"}

			json.NewEncoder(w).Encode(res)
		}
		
	}
}