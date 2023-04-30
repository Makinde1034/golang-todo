package structs 

type Response struct {
	Name string `json:"code"`
	Body interface{} `json:"body"`
}
	
type PostRequest struct {
	Name string `json:"name"`
	Todo string `json:"todo"`
} 

type GetTodoRequest struct {
	Name string `json:"name"`
}