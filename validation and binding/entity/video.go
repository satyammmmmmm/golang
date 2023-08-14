package entity

type Person struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=30"`
	Email     string `json:"email" validate:"required,email"`
}

type Video struct {
	Title  string `json:"title" binding:"min=2,max=10" validate:"is-cool"`
	Desc   string `json:"desc" binding:"max=20"`
	URL    string `json:"url" binding:"required,url"`
	Author string `json:"author" binding:"required"`
}
