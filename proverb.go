package proverb

import "strings"

// Proverb function generates a proverb
func Proverb(rhyme []string) []string {
	return proverbVA(rhyme)
}

// proverbVA function generates a proverb (Version A)
func proverbVA(rhyme []string) []string {
	length := len(rhyme)
	res := make([]string, 0, length)
	if length == 0 {
		return res
	}
	for i := 1; i < length; i++ {
		res = append(res, forWant(rhyme[i-1], rhyme[i]))
	}
	res = append(res, forAll(rhyme[0]))
	return res
}

func forWant(want, lost string) string {
	b := strings.Builder{}
	b.WriteString("For want of a ")
	b.WriteString(want)
	b.WriteString(" the ")
	b.WriteString(lost)
	b.WriteString(" was lost.")
	return b.String()
}

func forAll(want string) string {
	b := strings.Builder{}
	b.WriteString("And all for the want of a ")
	b.WriteString(want)
	b.WriteString(".")
	return b.String()
}

// proverbVB function generates a proverb (Version B)
func proverbVB(rhyme []string) []string {
	length := len(rhyme)
	res := make([]string, 0, length)
	if length == 0 {
		return res
	}
	b := strings.Builder{}
	for i := 1; i < length; i++ {
		b.Reset()
		b.WriteString("For want of a ")
		b.WriteString(rhyme[i-1])
		b.WriteString(" the ")
		b.WriteString(rhyme[i])
		b.WriteString(" was lost.")
		res = append(res, b.String())
	}
	b.Reset()
	b.WriteString("And all for the want of a ")
	b.WriteString(rhyme[0])
	b.WriteString(".")
	res = append(res, b.String())
	return res
}
