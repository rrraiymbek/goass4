package entity

import (
	"errors"
	"strings"
)

type Contact struct {
	ID         string
	FirstName  string
	LastName   string
	MiddleName string
	Phone      string
}

type Group struct {
	ID       string
	Name     string
	Contacts []Contact
}

func NewContact(id, firstName, lastName, middleName, phone string) (*Contact, error) {
	if len(phone) != len(strings.Trim(phone, "0123456789")) {
		return nil, errors.New("phone number should contain only digits")
	}

	return &Contact{
		ID:         id,
		FirstName:  firstName,
		LastName:   lastName,
		MiddleName: middleName,
		Phone:      phone,
	}, nil
}

func (c *Contact) FullName() string {
	return c.FirstName + " " + c.MiddleName + " " + c.LastName
}

func NewGroup(id, name string, contacts []Contact) (*Group, error) {
	if len(name) > 250 {
		return nil, errors.New("group name should not exceed 250 characters")
	}

	return &Group{
		ID:       id,
		Name:     name,
		Contacts: contacts,
	}, nil
}
