package referenceorpointer

import (
	"fmt"
)

type User struct {
	name  string
	email string
}

type Admin struct {
	User
	level int8
}

func (user User) notify() {
	fmt.Println("------ Info ------")
	fmt.Println("User: ", user.name)
	fmt.Println("Email: ", user.email)
	fmt.Println()
}

func (user *User) changeEmail(email string) {
	if email != "" {
		user.email = email
		return
	}

	fmt.Println("Error: input email is empty")
}

func (user *User) changeName(name string) {
	if name != "" {
		user.name = name
		return
	}
	fmt.Println("Error: input name is empty")
}

func (user *User) showAddress() {
	fmt.Println("ShowAddress: ", &user)
}

func (user User) showCopyAddress() {
	fmt.Println("ShowCopyAddress: ", &user)
}

func main() {
	var bill User
	bill.email = "bill@email1.com"
	bill.name = "bill"

	fmt.Println("Print Bill Information")
	bill.notify()

	bill.changeName("bill2")
	bill.changeEmail("bill2@email1.com")
	fmt.Println("After changed to bill2, print Bill Information, change name to bill2, email to bill2@email1.com")
	bill.notify()

	gates := &User{"gates", "gates@email2.com"}
	fmt.Println("Print Gates Information")
	gates.notify()

	gates.changeName("Gate2")
	fmt.Println("After changed to Gates, print Bill Information, change name to Gate2")
	gates.notify()

	fmt.Println("var gates address: ", &gates)
	gates.showAddress()
	gates.showCopyAddress()

	return
}
