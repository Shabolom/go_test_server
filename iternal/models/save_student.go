package models

// SaveStudent нужен для того чтоб распарсить получаемые данные от клиента
// `json:"name,omitempty"` говорит о том что принимает в себя джейсон с ключем name
// omitempty говорит о то что строка может быть пустой
// binding:"required является тегом который обозночает что поле обязательно должно быть заполненно или вернет ошибку
type SaveStudent struct {
	Name    string `json:"name,omitempty" binding:"required"`
	Surname string `json:"surname" binding:"required"`
	Gender  string `json:"gender" binding:"required"`
	Age     int    `json:"age" binding:"required"`
	Country string `json:"country" binding:"required"`
	Email   string `json:"email" binding:"required"`
}
