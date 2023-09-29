package functions

import (
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

// TYPE ENFORCER
func integer(el any) int {
	return int(el.(float64))
}

func float(el any) float64 {
	return el.(float64)
}

// FORMATTERS
func decimalformat(local string, fixed int, el float64) string {
	decimal := number.Decimal(el, number.Scale(fixed))
	return _printer(local).Sprintf("%v", decimal)
}

func percentformat(local string, fixed int, el float64) string {
	percentage := number.Percent(el, number.Scale(fixed))
	return _printer(local).Sprintf("%v", percentage)
}

func currencyformat(local string, el float64) string {
	lang := language.MustParse(local)
	cur, _ := currency.FromTag(lang)
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(el, number.Scale(scale))
	return _printer(local).Sprintf("%v%v", currency.Symbol(cur), dec)
}

func numberformat(local string, i int) string {
	return _printer(local).Sprintf("%d", i)
}

func _printer(local string) *message.Printer {
	return message.NewPrinter(message.MatchLanguage(local))
}
