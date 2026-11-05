package auth

func GenerateToken() string {
	if !validatePassword("mypassword123") {
		return ""
	}

	return createToken()
}

func createToken() string {
	return "My first token"
}

type Token struct {
	Value  string
	secret string
}

func NewToken() Token {
	return Token{
		Value:  "access-token",
		secret: "private-secret",
	}
}

func (t *Token) GetSecret() string {
	return t.secret
}

func (t *Token) ChangeSecret(newSecret string) {
	t.secret = newSecret
}

type Session struct {
	Scopes []string
}

func (s Session) AddScope(scope string) {
	s.Scopes = append(s.Scopes, scope)
}

type UserPermissions struct {
	Permissions map[string]bool
}

// func (u UserPermissions) Grant(permission string) {
// 	u.Permissions[permission] = true

// }

// func (u UserPermissions) Grant(permission string) {
// 	if u.Permissions == nil {
// 		u.Permissions = make(map[string]bool)
// 	}

// 	u.Permissions[permission] = true
// }

func (u *UserPermissions) Grant(permission string) {
	if u.Permissions == nil {
		u.Permissions = make(map[string]bool)
	}

	u.Permissions[permission] = true
}

type Account struct {
	Balance *int
}

func (a Account) Deposit(amount int) {
	*a.Balance += amount
}

func (a Account) ReplaceBalance(newBalance int) {
	a.Balance = &newBalance
}
