package entity

type User struct {
	ID    string `json:"id" form:"-"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}
