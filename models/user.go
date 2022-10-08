package models

import "golang.org/x/crypto/bcrypt"



type User struct {
	// json tag is used to specify the name of the field in the JSON response
	Id			uint	`json:"id"`
	FirstName	string	`json:"first_name"`
	LastName	string	`json:"last_name"`
	Email		string 	`json:"email" gorm:"unique"`
	Password	[]byte	`json:"-"`
}

func (user *User) SetPassword(password string){
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14) // so we store password as a hash
	user.Password = hashedPassword
}


func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}