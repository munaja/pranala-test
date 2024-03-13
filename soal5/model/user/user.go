package user

type User struct {
	Id   int    `json:"id" gorm:"autoincrement"`
	Name string `json:"name" gorm:"size:100"`
}

type CreateDto struct {
	Name string `json:"name"`
}

type UpdateDto struct {
	Id string `json:"id"`
}

type DeleteDto struct {
	Id string `json:"id"`
}
