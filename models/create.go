package models

import (
	"fmt"
	"get-going/structs"
)

func CreateTodo(name string, todo string) error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES(?, ?)", name, todo)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ReadAll() ([]structs.PostRequest,error){                   
	rows, err := con.Query("SELECT * FROM TODO")

	if err != nil {
		return nil, err
	}
	todos := []structs.PostRequest{}
	
	for rows.Next(){
		data := structs.PostRequest{}
		rows.Scan(&data.Name,&data.Todo)
		todos = append(todos, data)
	}

	return todos,nil
}

func ReadByName(name string) ([]structs.PostRequest, error){      
	rows, err := con.Query("SELECT * FROM TODO WHERE name=?",name)

	if err != nil {
		return nil, err
	}

	todos := []structs.PostRequest{}

	for rows.Next() {
		data := structs.PostRequest{}
		rows.Scan(&data.Name,&data.Todo)
		todos = append(todos, data)
	}

	return todos, nil
}

func DeleteTodo(name string) error {
	
	deleteQ, err := con.Query("DELETE FROM TODO WHERE name=?",name)
	defer deleteQ.Close()

	if(err != nil){
		return err
	}

	return nil

	
}