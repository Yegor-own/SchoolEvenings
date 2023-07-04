package input

type UserCreate struct {
	Name       string `json:"name" :"name"`
	Surname    string `json:"surname" :"surname"`
	Patronymic string `json:"patronymic" :"patronymic"`
	Email      string `json:"email" :"email"`
	Phone      int    `json:"phone" :"phone"`
	Password   string `json:"password" :"password"`
}

type UserUpdate struct {
	ID         uint   `json:"id" :"id"`
	Name       string `json:"name" :"name"`
	Surname    string `json:"surname" :"surname"`
	Patronymic string `json:"patronymic" :"patronymic"`
	Email      string `json:"email" :"email"`
	Phone      int    `json:"phone" :"phone"`
}

type UserById struct {
	ID uint `json:"id" :"id"`
}

type UserDelete struct {
	ID uint `json:"id" :"id"`
}
