package com_qa_test

import "fmt"

type User struct {
	Uid    int
	UName  string
	UHobby map[int]string
}

func GetUser(uid int, uname string, uhobby map[int]string) *User {
	if uid > 10 && uname == "adam" {
		fmt.Println("pass")
	}
	return &User{
		Uid:    uid,
		UName:  uname,
		UHobby: uhobby,
	}
}
