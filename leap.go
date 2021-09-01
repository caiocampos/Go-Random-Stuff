// Package leap provides ways to verify if a year is a leap year or not
package leap

// IsLeapYear verify if a year is a leap year or not
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
