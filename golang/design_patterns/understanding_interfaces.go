package designpatterns

import "database/sql"

// This file will contain the code for understanding how interfaces can be used in golang, we will be taking the example of transaction

type User struct {
	Id string
	Name string
	Email string
	Phone string
}

type UserOperations interface {
	Create(User) error
	Update(User) error
	Delete(id string) error
	Read(id string) User
	ReadAll() []User
}

type UserRepository struct {
	db *sql.DB
}

type UserController struct {
	UserOperations UserOperations
}

// Implementing the interface

func (uc *UserRepository) Create(user User) (error) {
	// business logic
	return nil
}

func (uc *UserRepository) Update(user User) error {
	// business logic
	return nil
}

func (uc *UserRepository) Delete(id string) error {
	// business logic
	return nil
}

func (uc *UserRepository) Read(id string) User {
	// business logic
	return User{}
}

func (uc *UserRepository) ReadAll() []User {
	return nil
}

