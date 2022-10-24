package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type User struct {
	name string
}

func (user User) notify() {
	fmt.Println("name: ", user.name)
}

type Admin struct {
	User
	level string
}

func (admin Admin) notify() {
	fmt.Println("name: ", admin.name, " level: ", admin.level)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	var u User
	u.name = "user"

	fmt.Println("notify by user")
	sendNotification(u)

	admin := Admin{User{"admin"}, "1"}
	fmt.Println("notify by admin")
	sendNotification(admin)
	fmt.Println("notify by admin.user")
	sendNotification(admin.User)

	return
}
