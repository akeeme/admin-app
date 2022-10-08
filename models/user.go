package models

type User struct {
	// json tag is used to specify the name of the field in the JSON response
	Id			uint	`json:"id"`
	FirstName	string	`json:"first_name"`
	LastName	string	`json:"last_name"`
	Email		string 	`json:"email" gorm:"unique"`
	Password	[]byte	`json:"-"`
}