package main

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"time"
)

type User struct {
	email string
	pass  string
	date  time.Time
}

var Users []User

func Register(email, pass string) bool {
	re, _ := regexp.Compile("^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(.[a-zA-Z0-9_-]+)+$")
	if !re.MatchString(email) {
		return false
	}
	if len(pass) <= 10 || len(pass) >= 30 {
		return false
	}
	up, down, symbol, number := 0, 0, 0, 0
	for _, letter := range pass {
		if 'A' <= letter && 'Z' >= letter {
			up = 1
		} else if 'a' <= letter && 'z' >= letter {
			down = 1
		} else if '0' <= letter && '9' >= letter {
			number = 1
		} else {
			symbol = 1
		}
	}
	if up+down+symbol+number < 3 {
		return false
	}
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(pass))
	password := Md5Inst.Sum([]byte("test"))
	Users = append(Users, User{email: email, pass: string(password), date:time.Now()})
	return true
}

func Login(email, pass string) bool {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(pass))
	password := Md5Inst.Sum([]byte("test"))
	for i := 0; i < len(Users); i++ {
		if email == Users[i].email && string(password) == Users[i].pass {
			return true
		}
	}
	return false
}

func ChangePass(email, oldpass, newpass string) bool {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(oldpass))
	password := Md5Inst.Sum([]byte("test"))
	for i := 0; i < len(Users); i++ {
		if email == Users[i].email && string(password) == Users[i].pass {
			Md5Inst.Write([]byte(newpass))
			password := Md5Inst.Sum([]byte("test"))
			Users[i].pass = string(password)
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(Register("1234@11.com","131213qwrqweQEWWE3"))
	fmt.Println(Login("1234@11.com","131213qwrqweQEWWE3"))
	fmt.Println(Login("1234@111.com","13123"))
	fmt.Println(ChangePass("1234@11.com", "131213qwrqweQEWWE3","123456QWER2qwer"))
}
