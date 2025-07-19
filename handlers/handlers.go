package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"passwordGenerator/internal/forms"
	"passwordGenerator/internal/password"
)

type Handler struct {
	tmpl *template.Template
}

func NewHandler() (*Handler, error) {
	t, err := template.ParseFiles(filepath.Join("templates", "form.html"))
	if err != nil {
		return nil, err
	}
	return &Handler{tmpl: t}, nil
}

func (h *Handler) PasswordHandler(w http.ResponseWriter, r *http.Request) {
	data := forms.PasswordForm{
		Length:    8,
		Digits:    true,
		Lowercase: true,
		Uppercase: true,
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			data.Error = "Error with parsing form"
			err := h.tmpl.Execute(w, data)
			if err != nil {
				return
			}
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

		var generator password.Generator = password.New(length, digits, lowercase, uppercase, true)
		generatedPassword, err := generator.Generate()
		if err != nil {
			data.Error = err.Error()
		} else {
			data.Password = generatedPassword
		}
	}

	err := h.tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}
}
