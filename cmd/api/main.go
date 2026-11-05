package main

import (
	"fmt"

	"github.com/Triplebond1/secure-notes/internal/auth"
)

func main() {
	// fmt.Println(auth.GenerateToken())

	// token := auth.NewToken()

	// fmt.Println(token.GetSecret())

	// token.ChangeSecret("new-secret")

	// fmt.Println(token.GetSecret())

	// session := auth.Session{
	// 	Scopes: []string{"read"},
	// }

	// session.AddScope("write")

	// fmt.Println(session.Scopes)

	// session := auth.Session{
	// 	Scopes: make([]string, 1, 5),
	// }

	// session.Scopes[0] = "read"

	// session.AddScope("write")

	// fmt.Println(session.Scopes)
	// fmt.Println("Length:", len(session.Scopes))
	// fmt.Println("Capacity:", cap(session.Scopes))
	// expanded := session.Scopes[:2]

	// fmt.Println("Expanded:", expanded)

	// user := auth.UserPermissions{
	// 	Permissions: map[string]bool{
	// 		"read": true,
	// 	},
	// }

	user := auth.UserPermissions{}
	// user := auth.UserPermissions{
	// 	Permissions: make(map[string]bool),
	// }

	user.Grant("write")

	fmt.Println(user.Permissions)

	// balance := 100

	// account := auth.Account{
	// 	Balance: &balance,
	// }

	// account.Deposit(50)

	// fmt.Println(*account.Balance)

	// balance := 100

	// account := auth.Account{
	// 	Balance: &balance,
	// }

	// account.ReplaceBalance(500)

	// fmt.Println(*account.Balance)

	// balance := new(int)

	// fmt.Println(*balance)

	// *balance = 100

	// fmt.Println(*balance)

	email := auth.EmailNotification{}
	sms := auth.SMSNotification{}

	email.Send("Welcome to Secure Notes")
	sms.Send("Your verification code is 123456")

}
