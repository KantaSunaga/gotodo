package models

type ToDo struct{
	Title string `json:"title"`
	Body  string `json:"string"`
	Done  bool   `json:"done"`
}

func (td *ToDo) Create() (result bool) {
	result = true
	return result
}