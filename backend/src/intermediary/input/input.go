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
	Name       string `json:"name" :"name"`
	Surname    string `json:"surname" :"surname"`
	Patronymic string `json:"patronymic" :"patronymic"`
	Phone      int    `json:"phone" :"phone"`
}

type CourseById struct {
	ID uint `json:"id" :"id"`
}

type UserDelete struct {
	ID uint `json:"id" :"id"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
