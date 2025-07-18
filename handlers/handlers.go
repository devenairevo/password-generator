package handlers

import (
	"html/template"
	"net/http"
	"passwordGenerator/internal/forms"
	"passwordGenerator/internal/password"
	"path/filepath"
	"strconv"
)

var tmpl = template.Must(template.ParseFiles(filepath.Join("templates", "form.html")))

func PasswordHandler(w http.ResponseWriter, r *http.Request) {

	data := forms.PasswordForm{
		Length:    8,
		Digits:    true,
		Lowercase: true,
		Uppercase: true,
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			return
		}
		length, _ := strconv.Atoi(r.FormValue("length"))
		digits := r.FormValue("digits") == "on"
		lowercase := r.FormValue("lowercase") == "on"
		uppercase := r.FormValue("uppercase") == "on"

		data.Length = length
		data.Digits = digits
		data.Lowercase = lowercase
		data.Uppercase = uppercase

		var generator password.Generator = password.New(length, digits, lowercase, uppercase)
		generatedPassword, err := generator.Generate()
		if err != nil {
			data.Error = err.Error()
		} else {
			data.Password = generatedPassword
		}
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		return
	}
}
