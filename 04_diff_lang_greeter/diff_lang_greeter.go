package main

import "strings"

var langs  =map[string]string{
	"spanish":"Hola",
	"arabic" : "Marhaba",
	"english":"Hello",
	"french":"Bonjour",
}
func greeter(name,lang string) string  {
	salutation:=""
	if strings.EqualFold(langs[lang],""){
		salutation="Hello"
	}else {
		salutation=langs[lang]
	}
	if name==""{
		return salutation+", World"
	}
	return salutation+", "+name
}
