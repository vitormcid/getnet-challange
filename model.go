package main

import (
	"database/sql"
	"errors"
)

type user struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *user) getUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *user) updateUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *user) deleteUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *user) createUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	return nil, errors.New("Not implemented")
}
