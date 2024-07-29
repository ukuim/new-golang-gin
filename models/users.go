package models

type Users struct {
	Uuid string
}

func (user *Users) TableName() string {
	return "users"
}
