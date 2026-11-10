package main

import (
	"fmt"
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

	// user := auth.UserPermissions{}
	// // user := auth.UserPermissions{
	// // 	Permissions: make(map[string]bool),
	// // }

	// user.Grant("write")

	// fmt.Println(user.Permissions)

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

	// email := auth.EmailNotification{}
	// sms := auth.SMSNotification{}

	// var notification *Notification = nil

	// var empty any = nil

	// // email.Send("Welcome to Secure Notes")
	// // sms.Send("Your verification code is 123456")
	// err := checkError()

	// fmt.Println(err == nil)

	// inspect(notification)
	// inspect(empty)

	// inspect("Secure Notes")
	// inspect(100)
	// inspect(true)
	// // inspect(auth.EmailNotification{})
	// inspect(&Notification{
	// 	Message: "Hello",
	// })
	// inspect(Notification{
	// 	Message: "Hello",
	// })

	// closeNotification(email)
	// closeNotification(sms)

	// sendMessage(email, "Hello from email")
	// sendMessage(sms, "Hello from SMS")

	// notify(email, "Welcome to Secure Notes")
	// notify(sms, "Your verification code is 123456")

	// 	inspectType("Secure Notes")
	// 	inspectType(100)
	// 	inspectType(true)
	// 	inspectType(Notification{
	// 		Message: "Hello",
	// 	})
	// 	inspectType(&Notification{
	// 		Message: "Hello",
	// 	})

	// 	notification := Notification{
	// 		Message: "Hello",
	// 		secret:  "hidden",
	// 	}

	// 	v := reflect.ValueOf(&notification).Elem()

	// 	fmt.Println("Message:", v.FieldByName("Message").CanSet())
	// 	fmt.Println("Secret:", v.FieldByName("secret").CanSet())
	// 	modifyMessage(&notification)

	// 	fmt.Println(notification.Message)
	// }

	// func notify(n auth.Notifier, message string) {
	// 	n.Send(message)
	// }

	// func sendMessage(s auth.MessageSender, message string) {
	// 	s.Send(message)
	// }

	// func closeNotification(n auth.MessageCloser) {
	// 	n.Close()
	// }

	// func inspect(value auth.Anything) {
	// 	fmt.Println(value)
	// }

	// func inspect(value any) {
	// 	fmt.Printf("Value: %v, Type: %T\n", value, value)
	// }

	// func inspect(value any) {
	// 	text, ok := value.(string)

	// 	if !ok {
	// 		fmt.Println("Value is not a string")
	// 		return
	// 	}

	//		fmt.Println("Text:", text)
	//	}
	// func inspect(value any) {
	// 	text, ok := value.(string)

	// 	fmt.Printf("text: %v, type: %T\n", text, text)
	// 	fmt.Printf("ok: %v, type: %T\n", ok, ok)
	// }

	// func inspect(value any) {
	// 	switch v := value.(type) {
	// 	case string:
	// 		fmt.Printf("%T\n", v)
	// 	case int:
	// 		fmt.Printf("%T\n", v)
	// 	case bool:
	// 		fmt.Printf("%T\n", v)
	// 	case Notification:
	// 		fmt.Println("Notification:", v.Message)
	// 	case nil:
	// 		fmt.Println("Nil interface")
	// 	case *Notification:
	// 		fmt.Println("Pointer Notification:", v == nil)
	// 	default:
	// 		fmt.Printf("%T\n", v)
	// 	}

	// user := User{
	// 	Name:     "John",
	// 	Password: "secret123",
	// }

	// // user1 := User{
	// // 	Name: "John",
	// // }

	// data, err := json.Marshal(user)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// inspectTags(User{})

	// user := User{
	// 	Name:     "",
	// 	Email:    "",
	// 	Password: "secret123",
	// 	Bio:      "",
	// }

	//validate(&user)

	// if err := validate(&user); err != nil {
	// 	fmt.Println("Validation failed:", err)
	// } else {
	// 	fmt.Println("Validation successful")
	// }

	// errors := validate(&user)

	// for _, err := range errors {
	// 	fmt.Println(err)
	// }
	// fmt.Println("String:")
	// fmt.Println(validate("John"))

	// fmt.Println("Integer:")
	// fmt.Println(validate(100))

	// fmt.Println("Nil:")
	// fmt.Println(validate(nil))

	// fmt.Println("Pointer:")
	// fmt.Println(validate(&user))

	// Print("Secure Notes")
	// Print(100)
	// Print(true)
	// Print(Notification{})
	// Add(10, 20)
	// Add("Hello", "World")

	// fmt.Println(Equal(10, 10))
	// fmt.Println(Equal("Go", "Go"))
	// fmt.Println(Equal(true, false))

	// a := [2]int{1, 2}
	// b := [2]int{1, 2}

	// fmt.Println(a == b) // true

	// a := []int{1, 2}
	// b := []int{1, 2}
	// c := []int{1, 3}

	// fmt.Println(slices.Equal(a, b))
	// fmt.Println(slices.Equal(a, c))

	// a := [][]int{{1, 2}, {3, 4}}
	// b := [][]int{{1, 2}, {3, 4}}

	// //fmt.Println(slices.Equal(a, b))
	// fmt.Println(slices.EqualFunc(a, b, func(x, y []int) bool {
	// 	return slices.Equal(x, y)
	// }))

	// scores := make([]int, 3, 5)

	// scores[0] = 10
	// scores[1] = 20
	// scores[2] = 30

	// other := scores

	// other = append(other, 40)

	// other[0] = 100

	// fmt.Println("Scores:", scores)
	// fmt.Println("Other:", other)
	// fmt.Println("Scores capacity:", cap(scores))
	// fmt.Println("Other capacity:", cap(other))

	// var a UserID = 10
	// var b UserID = 20

	// var userID UserID = 10
	// var productID ProductID = 20

	// fmt.Println(Add(userID, userID))
	// fmt.Println(Add(productID, productID))

	// fmt.Println(Add(ProductID(userID), productID))
	// fmt.Println(Add(a, b))
	// //fmt.Println(AddDifferent(userID, productID))
	// fmt.Println(AddDifferent(userID, 20.5))

	// fmt.Println(AddNumbers(10, 20))
	// fmt.Println(AddNumbers(10.5, 20.5))
	// fmt.Println(AddNumbers(10.5, 20))
	// fmt.Println(AddNumbers(10, 20.5))

	// var a int = 10
	// var b float64 = 20.5

	// //fmt.Println(AddNumbers(a, b))
	// fmt.Println(AddNumbers(float64(a), b))

	var numbers Stack[int]

	numbers.Push(10)
	numbers.Push(20)
	numbers.Push(30)

	fmt.Println(numbers.items)

	fmt.Println(numbers.Pop())
	fmt.Println(numbers.items)

	var empty Stack[int]

	fmt.Println(empty.Pop())
	item, ok := empty.Pop()

	fmt.Println(item, ok)

	var names Stack[string]

	names.Push("John")
	names.Push("Mary")

	name, ok := names.Pop()

	fmt.Println(name, ok)

	nam, ok := names.Pop()

	fmt.Println(nam, ok)

	na, ok := names.Pop()

	fmt.Println(na, ok)

	value, ok := numbers.Peek()

	fmt.Println(value, ok)
	fmt.Println(numbers.items)

	// var number Stack[[]int]

	// number.Push([]int{1, 2, 3})
	// number.Push([]int{4, 5, 6})

	// fmt.Println(number.items)

	var users Stack[User]

}

// type Stack[T any] struct {
// 	items []T
// }

type Stack[T comparable] struct {
	items []T
}

type User struct {
	ID   int
	Name string
	Tags []string
}

func (s *Stack[T]) Contains(value T) bool {
	for _, item := range s.items {
		if item == value {
			return true
		}
	}

	return false
}
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// func (s *Stack[T]) Pop() T {
// 	last := len(s.items) - 1
// 	item := s.items[last]
// 	s.items = s.items[:last]
// 	return item
// }

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	last := len(s.items) - 1
	item := s.items[last]
	s.items = s.items[:last]

	return item, true
}

// type Number interface {
// 	~int | ~int64 | ~float64
// }

// func AddNumbers[T Number](a, b T) T {
// 	return a + b
// }

// type UserID int
// type ProductID int

// func AddDifferent[A ~int, B ~float64](a A, b B) float64 {
// 	return float64(a) + float64(b)
// }

//	func AddDifferent[A interface{ ~int }, B interface{ ~int }](a A, b B) int {
//		return int(a) + int(b)
//	}

// func AddDifferent[A ~int, B ~int](a A, b B) int {
// 	return int(a) + int(b)
// }

func Add[T interface{ ~int }](a, b T) T {
	return a + b
}
func Print[T any](value T) {
	fmt.Println(value)
}

type Addable interface {
	int | int64 | float64 | string
}

// func Add[T Addable](a T, b T) {
// 	fmt.Println(a + b)
// }

func Equal[T comparable](a, b T) bool {
	return a == b
}

type Notification struct {
	Message string
	secret  string
}

// type MyError struct{}

// func (e *MyError) Error() string {
// 	return "something went wrong"
// }

// func checkError() error {
// 	var err *MyError = nil

// 	if err == nil {
// 		return nil
// 	}

// 	return err
// }

// func inspectType(value any) {
// 	t := reflect.TypeOf(value)
// 	v := reflect.ValueOf(value)

// 	fmt.Println("Type:", t)
// 	fmt.Println("Kind:", t.Kind())

// 	if t.Kind() == reflect.Ptr {
// 		fmt.Println("Element Type:", t.Elem())
// 		fmt.Println("Element Kind:", t.Elem().Kind())
// 	}

// 	fmt.Println("Value:", v)
// }

// func inspectType(value any) {
// 	v := reflect.ValueOf(value)
// 	// fmt.Println("Type1 :", v.Kind())
// 	if v.Kind() == reflect.Ptr {
// 		v = v.Elem()
// 	}

// 	if v.Kind() == reflect.Struct {
// 		for i := 0; i < v.NumField(); i++ {
// 			field := v.Field(i)
// 			fieldInfo := v.Type().Field(i)

// 			fmt.Println("Name:", fieldInfo.Name)
// 			fmt.Println("Value:", field)
// 			fmt.Println("CanSet:", field.CanSet())

// 		}
// 	}
// }

// func modifyMessage(value any) {
// 	v := reflect.ValueOf(value)

// 	if v.Kind() == reflect.Ptr {
// 		v = v.Elem()
// 	}

// 	field := v.FieldByName("Message")

// 	field.SetString("Modified through reflection")
// }

// func modifyMessage(value any) {
// 	v := reflect.ValueOf(value)

// 	fmt.Println("Addressable:", v.CanAddr())
// 	fmt.Println("Settable:", v.CanSet())

// 	if v.Kind() == reflect.Ptr {
// 		v = v.Elem()
// 	}

// 	fmt.Println("After Elem - Addressable:", v.CanAddr())
// 	fmt.Println("After Elem - Settable:", v.CanSet())

// 	field := v.FieldByName("Message")

// 	fmt.Println("Field Settable:", field.CanSet())

// 	field.SetString("Modified through reflection")
// }

// type User struct {
// 	Name     string
// 	password string
// }

// type User struct {
// 	Name     string `json:"name"`
// 	Password string `json:"-"` //- omit this field frm json output
// }

// type User struct {
// 	Name     string `json:"name,omitempty"`
// 	Password string `json:"password,omitempty"`
// }

// func inspectTags(value any) {
// 	t := reflect.TypeOf(value)

// 	if t.Kind() == reflect.Ptr {
// 		t = t.Elem()
// 	}

// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i)

// 		tag := field.Tag.Get("json")

// 		name, options, _ := strings.Cut(tag, ",")

// 		fmt.Println("JSON Name:", name)
// 		fmt.Println("JSON Options:", options)
// 		fmt.Println("Field:", field.Name)
// 		fmt.Println("JSON tag:", field.Tag.Get("json"))
// 	}
// }

// type User struct {
// 	Name     string `validate:"required"`
// 	Email    string `validate:"required"`
// 	Password string `validate:"required"`
// 	Bio      string `validate:"optional"`
// }

// var errors []error

// func validate(value any) []error {
// 	v := reflect.ValueOf(value)

// 	if !v.IsValid() {
// 		return []error{fmt.Errorf("invalid input")}
// 	}

// 	if v.Kind() == reflect.Ptr {
// 		if v.IsNil() {
// 			return []error{fmt.Errorf("nil pointer provided")}
// 		}

// 		v = v.Elem()
// 	}

// 	if v.Kind() != reflect.Struct {
// 		return []error{fmt.Errorf("expected a struct, got %s", v.Kind())}
// 	}

// 	t := v.Type()

// 	for i := 0; i < v.NumField(); i++ {
// 		field := t.Field(i)
// 		fieldValue := v.Field(i)

// 		rule := field.Tag.Get("validate")

// 		if rule == "required" && fieldValue.IsZero() {
// 			errors = append(errors, fmt.Errorf("%s is required", field.Name))
// 		}
// 	}
// 	return errors
// }
