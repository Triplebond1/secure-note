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

	// var numbers Stack[int]

	// numbers.Push(10)
	// numbers.Push(20)
	// numbers.Push(30)

	// fmt.Println(numbers.items)

	// fmt.Println(numbers.Pop())
	// fmt.Println(numbers.items)

	// var empty Stack[int]

	// fmt.Println(empty.Pop())
	// item, ok := empty.Pop()

	// fmt.Println(item, ok)

	// var names Stack[string]

	// names.Push("John")
	// names.Push("Mary")

	// name, ok := names.Pop()

	// fmt.Println(name, ok)

	// nam, ok := names.Pop()

	// fmt.Println(nam, ok)

	// na, ok := names.Pop()

	// fmt.Println(na, ok)

	// value, ok := numbers.Peek()

	// fmt.Println(value, ok)
	// fmt.Println(numbers.items)

	// var number Stack[[]int]

	// number.Push([]int{1, 2, 3})
	// number.Push([]int{4, 5, 6})

	// fmt.Println(number.items)

	//var users Stack[User]

	// var box Container[string] = &Box[string]{}

	// box.Add("Secure Notes")

	// fmt.Println(box.Get())

	// var box Container[int] = &Box[int]{}

	// box.Add(100)

	// fmt.Println(box.Get())

	// var intBox Container[int] = &Box[int]{}
	// var stringBox Container[string] = &Box[string]{}

	// intBox.Add(1000)
	// stringBox.Add("Secure Notes")

	// fmt.Println(intBox.Get())
	// fmt.Println(stringBox.Get())

	// intBox := &Box[int]{}
	// stringBox := &Box[string]{}

	// Store[int](intBox, 1000)
	// Store(stringBox, "Secure Notes")

	// result, err := divide(10, 0)

	// fmt.Println("Result:", result)
	// fmt.Println("Error:", err)
	// fmt.Printf("Error type: %T\n", err)

	// var err error = ValidationError{
	// 	Message: "username is required",
	// }

	// var err2 error = &ValidationError{
	// 	Message: "username is required",
	// }

	// //fmt.Println(err)
	// fmt.Println(err2)
	// fmt.Printf("Error type: %T\n", err2)

	// err := validateUsername("triplebond")

	// fmt.Printf("Error: %v\n", err)
	// fmt.Printf("Type: %T\n", err)

	// if err != nil {
	// 	fmt.Println("Validation failed:", err)
	// } else {
	// 	fmt.Println("Validation successful")
	// }

	// err3 := findNote(100)

	// if err3 != nil {
	// 	fmt.Println(err3)
	// }

	// err := findNote(200)

	// if errors.Is(err, ErrNoteNotFound) {
	// 	fmt.Println("The requested note does not exist")
	// } else {
	// 	fmt.Println("Something else happened")
	// }

	// err := validateUsername("")

	// var validationErr *ValidationError

	// if errors.As(err, &validationErr) {
	// 	fmt.Println("Field:", validationErr.Field)
	// 	fmt.Println("Message:", validationErr.Message)
	// }

	// err := authenticate("wrong-password")

	// var validationErr *ValidationError
	// var authErr *AuthenticationError

	// fmt.Println("Is validation error:", errors.As(err, &validationErr))
	// fmt.Println("Is authentication error:", errors.As(err, &authErr))

	// if authErr != nil {
	// 	fmt.Println("Authentication message:", authErr.Message)
	// }
	// handleGetNote(100)
	// handleGetNote(200)
	// handleGetNote(-1)
	// handleGetNote(500)

	// fmt.Println("Start")

	// defer fmt.Println("First deferred function")
	// defer fmt.Println("Second deferred function")

	// fmt.Println("End")

	// process()
	// fmt.Println("Main function continues")

	//err := getNotew()

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// process()
	// fmt.Println("Main function continues")

	// handler()
	// fmt.Println("Application continues")

	//text := "Hello"
	//text := "Hello 世界"

	// fmt.Println("String:", text)
	// fmt.Println("Length:", len(text))
	// fmt.Println("Character count:", utf8.RuneCountInString(text))

	// for i := 0; i < len(text); i++ {
	// 	fmt.Println(i, text[i])
	// }

	// for index, char := range text {
	// 	fmt.Printf("Index: %d | Character: %c | Rune value: %d\n",
	// 		index, char, char)
	// }

	// runes := []rune(text)

	// fmt.Println("String length:", len(text))
	// fmt.Println("Rune length:", len(runes))

	// for i, r := range runes {
	// 	fmt.Printf("Index: %d | Character: %c | Value: %d\n", i, r, r)
	// }

	// var letter rune = '世'
	// var number int32 = 19990

	// fmt.Println("Letter:", letter)
	// fmt.Println("Letter as character:", string(letter))
	// fmt.Println("Are they equal?", letter == rune(number))

	// fmt.Printf("Type of letter: %T\n", letter)

	// word := "Go"

	// var letter rune = 'G'

	// fmt.Println("String:", word)
	// fmt.Println("Rune:", letter)
	// fmt.Println("Rune as string:", string(letter))

	//text := "Go"
	// text := "Go 世界"

	// bytes := []byte(text)
	// runes := []rune(text)

	// fmt.Println("String:", text)
	// fmt.Println("Bytes:", bytes)
	// fmt.Println("Runes:", runes)

	// fmt.Printf("Bytes type: %T\n", bytes[0])
	// fmt.Printf("Runes type: %T\n", runes[0])

	// fmt.Println("Byte count:", len(bytes))
	// fmt.Println("Rune count:", len(runes))

	// text := "世"

	// bytes := []byte(text)
	// runes := []rune(text)

	// fmt.Printf("Bytes (decimal): %v\n", bytes)
	// fmt.Printf("Bytes (hex): %X\n", bytes)

	// fmt.Printf("Rune (decimal): %d\n", runes[0])
	// fmt.Printf("Rune (hex): %X\n", runes[0])

	// bytes := []byte{0xE4, 0xB8, 0x96}

	// r, size := utf8.DecodeRune(bytes)

	// fmt.Printf("Decoded rune: %c\n", r)
	// fmt.Printf("Rune value: U+%04X\n", r)
	// fmt.Printf("Bytes consumed: %d\n", size)

	// complete := []byte{0xE4, 0xB8, 0x96}
	// incomplete := []byte{0xE4, 0xB8}

	// r1, size1 := utf8.DecodeRune(complete)
	// r2, size2 := utf8.DecodeRune(incomplete)

	// fmt.Printf("Complete: %c | U+%04X | Size: %d\n", r1, r1, size1)
	// fmt.Printf("Incomplete: %c | U+%04X | Size: %d\n", r2, r2, size2)

	// data := []byte{0xE4, 0xB8}

	// for len(data) > 0 {
	// 	r, size := utf8.DecodeRune(data)

	// 	fmt.Printf("Rune: U+%04X | Size: %d\n", r, size)

	// 	data = data[size:]
	// }

	// valid := []byte("Hello 世界")
	// invalid := []byte{0xE4, 0xB8}

	// fmt.Println("Valid data:", utf8.Valid(valid))
	// fmt.Println("Invalid data:", utf8.Valid(invalid))
	// bytes := []byte{0xE4, 0xB8, 0x96}

	// b1 := uint32(bytes[0] & 0b00001111)
	// b2 := uint32(bytes[1] & 0b00111111)
	// b3 := uint32(bytes[2] & 0b00111111)

	// codePoint := (b1 << 12) | (b2 << 6) | b3

	// fmt.Printf("First part:  %04b\n", b1)
	// fmt.Printf("Second part: %06b\n", b2)
	// fmt.Printf("Third part:  %06b\n", b3)

	// fmt.Printf("Reconstructed: U+%04X\n", codePoint)
	// fmt.Printf("Decimal: %d\n", codePoint)

	// //bytes := []byte{0xE4, 0xB8, 0x96}

	// r, size := utf8.DecodeRune(bytes)

	// fmt.Printf("Decoded rune: %c\n", r)
	// fmt.Printf("Rune value: U+%04X\n", r)
	// fmt.Printf("Bytes consumed: %d\n", size)

	// number := uint32(5)

	// fmt.Printf("Original: %08b\n", number)
	// fmt.Printf("Shift 1:  %08b\n", number<<1)
	// fmt.Printf("Shift 2:  %08b\n", number<<2)
	// fmt.Printf("Shift 3:  %08b\n", number<<3)

	// a := uint32(13)
	// b := uint32(3)
	// fmt.Printf("Original a: %08b\n", a)
	// fmt.Printf("Original b: %08b\n", b)
	// fmt.Printf("AND: %04b\n", a&b)
	// fmt.Printf("OR:  %04b\n", a|b)

	// fmt.Printf("a:       %04b\n", a)
	// fmt.Printf("b:       %04b\n", b)
	// fmt.Printf("AND:     %04b\n", a&b)
	// fmt.Printf("OR:      %04b\n", a|b)
	// fmt.Printf("XOR:     %04b\n", a^b)
	// fmt.Printf("a XOR a: %04b\n", a^a)

	// result := a ^ b
	// restored := result ^ b
	// fmt.Println("Result:", result)
	// fmt.Println("Restored:", restored)

	// encrypted := a ^ b
	// decrypted := encrypted ^ b

	// fmt.Println("Original:", a)
	// fmt.Println("Key:", b)
	// fmt.Println("After XOR:", encrypted)
	// fmt.Println("After XOR again:", decrypted)

	original := "Hello"

	bytes := []byte(original)

	fmt.Println("Original string:", original)
	fmt.Println("Byte slice:", bytes)

	bytes[0] = 'Y'

	fmt.Println("After modification:")
	fmt.Println("Original string:", original)
	fmt.Println("Byte slice:", string(bytes))

}

// func getNotew() error {
// 	defer fmt.Println("Closing database connection")

// 	fmt.Println("Querying database")

// 	return errors.New("database query failed")
// }

// func process() {
// 	defer fmt.Println("Deferred cleanup")

// 	fmt.Println("Processing started")

// 	panic("something went terribly wrong")

// 	fmt.Println("Processing completed")
// }

// func service() {
// 	fmt.Println("Service started")
// 	panic("service crashed")
// }

// func handler() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Handler recovered:", r)
// 		}
// 	}()

// 	//service()

// 	//go service()

// 	go func() {
// 		defer func() {
// 			if r := recover(); r != nil {
// 				fmt.Println("Worker recovered:", r)
// 			}
// 		}()

// 		service()
// 	}()

// 	time.Sleep(time.Second)

// 	fmt.Println("Handler completed")
// }

// func process() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered from panic:", r)
// 		}
// 	}()

// 	defer fmt.Println("Deferred cleanup")

// 	fmt.Println("Processing started")

// 	panic("something went terribly wrong")

// 	fmt.Println("Processing completed")
// }

// var ErrNoteNotFound = errors.New("note not found")
// var ErrInvalidNoteID = errors.New("invalid note ID")
// var ErrDatabaseUnavailable = errors.New("database unavailable")

// Repository Layer
// func getNoteFromDB(id int) (string, error) {
// 	if id <= 0 {
// 		return "", ErrInvalidNoteID
// 	}

// 	// if id == 500 {
// 	// 	return "", ErrDatabaseUnavailable
// 	// }

// 	if id == 500 {
// 		return "", fmt.Errorf("repository: database query failed: %w", ErrDatabaseUnavailable)
// 	}

// 	if id != 100 {
// 		return "", ErrNoteNotFound
// 	}

// 	return "My first secure note", nil
// }

// // Service Layer
// func getNote(id int) (string, error) {
// 	note, err := getNoteFromDB(id)

// 	if err != nil {
// 		return "", fmt.Errorf("service: get note: %w", err)
// 	}

// 	return note, nil
// }

// // HTTP Handler Layer (simulated)
// // func handleGetNote(id int) {
// // 	note, err := getNote(id)

// // 	if err != nil {
// // 		fmt.Println("Handler error:", err)
// // 		return
// // 	}

// // 	fmt.Println("Note:", note)
// // }

// // func handleGetNote(id int) {
// // 	note, err := getNote(id)

// // 	if err != nil {
// // 		if errors.Is(err, ErrNoteNotFound) {
// // 			fmt.Println("HTTP Status: 404 Not Found")
// // 			return
// // 		}

// // 		fmt.Println("HTTP Status: 500 Internal Server Error")
// // 		return
// // 	}

// // 	fmt.Println("Note:", note)
// // }

// func handleGetNote(id int) {
// 	note, err := getNote(id)

// 	// if err != nil {
// 	// 	switch {
// 	// 	case errors.Is(err, ErrNoteNotFound):
// 	// 		fmt.Println("HTTP Status: 404 Not Found")

// 	// 	case errors.Is(err, ErrInvalidNoteID):
// 	// 		fmt.Println("HTTP Status: 400 Bad Request")

// 	// 	default:
// 	// 		fmt.Println("HTTP Status: 500 Internal Server Error")
// 	// 	}

// 	// 	return
// 	// }

// 	if err != nil {
// 		fmt.Println("Complete error:", err)

// 		fmt.Println(
// 			"Is database unavailable:",
// 			errors.Is(err, ErrDatabaseUnavailable),
// 		)

// 		return
// 	}
// 	fmt.Println("Note:", note)
// }

// type AuthenticationError struct {
// 	Message string
// }

// func (e *AuthenticationError) Error() string {
// 	return e.Message
// }

// func authenticate(password string) error {
// 	if password != "secret123" {
// 		return fmt.Errorf(
// 			"login failed: %w",
// 			&AuthenticationError{
// 				Message: "invalid credentials",
// 			},
// 		)
// 	}

// 	return nil
// }

// // func findNote(id int) error {
// // 	if id == 100 {
// // 		return nil
// // 	}

// // 	return errors.New("note not found")
// // }

// // func findNote(id int) error {
// // 	if id == 100 {
// // 		return nil
// // 	}

// //		return fmt.Errorf("findNote: failed to retrieve note: %w", ErrNoteNotFound)
// //	}

// func findNote(id int) error {
// 	if id == 100 {
// 		return nil
// 	}

// 	return fmt.Errorf("findNote: failed to retrieve note: %v", ErrNoteNotFound)
// }

// type ValidationError struct {
// 	Field   string
// 	Message string
// }

// func (e *ValidationError) Error() string {
// 	return e.Field + ": " + e.Message
// }

// // func validateUsername(username string) error {
// // 	if username == "" {
// // 		return &ValidationError{
// // 			Field:   "username",
// // 			Message: "is required",
// // 		}
// // 	}

// // 	return nil
// // }

// func validateUsername(username string) error {
// 	if username == "" {
// 		validationErr := &ValidationError{
// 			Field:   "username",
// 			Message: "is required",
// 		}

// 		return fmt.Errorf("registration failed: %w", validationErr)
// 	}

// 	return nil
// }

// type ValidationError struct {
// 	Message string
// }

// func (e *ValidationError) Error() string {
// 	return e.Message
// }

// func divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("cannot divide by zero")
// 	}

// 	return a / b, nil
// }

// func Store[T any](box Container[T], value T) {
// 	box.Add(value)
// 	fmt.Println(box.Get())
// }

// type Container[T any] interface {
// 	Add(item T)
// 	Get() T
// }

// type Box[T any] struct {
// 	item T
// }

// func (b *Box[T]) Add(item T) {
// 	b.item = item
// }

// func (b *Box[T]) Get() T {
// 	return b.item
// }

// type Stack[T any] struct {
// 	items []T
// }

// type Stack[T comparable] struct {
// 	items []T
// }

// type User struct {
// 	ID   int
// 	Name string
// 	Tags []string
// }

// func (s *Stack[T]) Contains(value T) bool {
// 	for _, item := range s.items {
// 		if item == value {
// 			return true
// 		}
// 	}

// 	return false
// }
// func (s *Stack[T]) Peek() (T, bool) {
// 	if len(s.items) == 0 {
// 		var zero T
// 		return zero, false
// 	}

// 	return s.items[len(s.items)-1], true
// }

// func (s *Stack[T]) Push(item T) {
// 	s.items = append(s.items, item)
// }

// func (s *Stack[T]) Pop() T {
// 	last := len(s.items) - 1
// 	item := s.items[last]
// 	s.items = s.items[:last]
// 	return item
// }

// func (s *Stack[T]) Pop() (T, bool) {
// 	if len(s.items) == 0 {
// 		var zero T
// 		return zero, false
// 	}

// 	last := len(s.items) - 1
// 	item := s.items[last]
// 	s.items = s.items[:last]

// 	return item, true
// }

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

// func Add[T interface{ ~int }](a, b T) T {
// 	return a + b
// }
// func Print[T any](value T) {
// 	fmt.Println(value)
// }

// type Addable interface {
// 	int | int64 | float64 | string
// }

// // func Add[T Addable](a T, b T) {
// // 	fmt.Println(a + b)
// // }

// func Equal[T comparable](a, b T) bool {
// 	return a == b
// }

// type Notification struct {
// 	Message string
// 	secret  string
// }

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
