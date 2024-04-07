package persistence

import (
	"github.com/WBarge/GoContactManager/models"
)

type PersonRepo interface {
	GetAll() []models.Person
	GetPerson(id uint)
	Insert(person models.Person)
	Update(person models.Person)
	Delete(person models.Person)
}

type PhoneRepo interface {
	GetPhone(id uint)
	Insert(phone models.Phone)
	Update(phone models.Phone)
	Delete(phone models.Phone)
}

type PersonalPhoneRepo interface {
	DoesPersonHaveNumberAsync(personId uint, countryCode string, areaCode string, number string, phoneType string)
	AddPhoneNumberToPerson(personId uint, countryCode string, areaCode string, number string, phoneType string)
}
