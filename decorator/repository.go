package main

import (
	"encoding/json"
	"fmt"
)

type userInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type userRepository struct {
	l     logger
	store map[int]userInfo
}

func newUserRepository(logger logger) *userRepository {
	return &userRepository{l: logger, store: make(map[int]userInfo)}
}

func (e *userRepository) allDisplay() {
	for _, v := range e.store {
		fmt.Println("保存されたデータ:", v)
	}
}

func (e *userRepository) register(id int, name, email string) {
	info := userInfo{Name: name, Email: email}
	e.store[id] = info
	b, _ := json.Marshal(info)
	e.l.write(b)
}
