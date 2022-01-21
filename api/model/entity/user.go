package entity
type user struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

//test data
var users = []user{
	{ID: "1",Name: "sample"},
}