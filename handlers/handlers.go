package handlers

import (
	"html/template"
	"net/http"
	"passwordGenerator/internal/password"
	"strconv"
)

var tmpl = template.Must(template.New("form").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Password Generator</title>
</head>
<body>
	<center>
		<h1>Password Generator</h1>
		<form method="POST">
			<label>
				Length:
				<input type="number" name="length" min="1" max="40" value="{{.Length}}" required>
			</label>
			<br>
			<label>
				<input type="checkbox" name="digits" {{if .Digits}}checked{{end}}>
				Include digits
			</label>
			<br>
			<label>
				<input type="checkbox" name="lowercase" {{if .Lowercase}}checked{{end}}>
				Include lowercase
			</label>
			<br>
			<label>
				<input type="checkbox" name="uppercase" {{if .Uppercase}}checked{{end}}>
				Include uppercase
			</label>
			<br>
			<button type="submit">Generate</button>
		</form>
		{{if .Password}}
		<div class="result">Generated password: <b>{{.Password}}</b></div>
		{{end}}
		{{if .Error}}
		<div class="result" style="color: #a22;">Error: {{.Error}}</div>
		{{end}}
	</center>
</body>
</html>
`))

type Form struct {
	Length    int
	Digits    bool
	Lowercase bool
	Uppercase bool
	Password  string
	Error     string
}

func PasswordHandler(w http.ResponseWriter, r *http.Request) {

	// By default set
	data := Form{
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
