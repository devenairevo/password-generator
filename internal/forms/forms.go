package forms

type PasswordForm struct {
	Length    int
	Digits    bool
	Lowercase bool
	Uppercase bool
	Password  string
	Error     string
}
