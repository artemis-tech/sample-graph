// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

type NewUser struct {
	Name     string `json:"name"`
	Surename string `json:"surename"`
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Surename  string `json:"surename"`
	CreatedAt string `json:"createdAt"`
}
