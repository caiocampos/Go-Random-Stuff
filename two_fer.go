package twofer

import "strings"

/*
ShareWith function receives a string (name) as parameter, if name is empty then it is replaced by "you".
This function return the string "One for {{name}}, one for me."
*/
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	var b strings.Builder

	b.WriteString("One for ")
	b.WriteString(name)
	b.WriteString(", one for me.")
	return b.String()
}
