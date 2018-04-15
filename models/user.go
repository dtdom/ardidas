package models

type User struct {
	ID    string `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
}

func (u *User) Store() {
	db.Create(&u)

}

func (u *User) Get() {
	db.First(u)
}

func GetAllUsers() []User {
	var users []User
	db.Find(&users)
	return users
}
