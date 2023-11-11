package templates

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
func decimalFormat(local string, fixed int, el float64) string {
	decimal := number.Decimal(el, number.Scale(fixed))
	return printer(local).Sprintf("%v", decimal)
}

func percentFormat(local string, fixed int, el float64) string {
	percentage := number.Percent(el, number.Scale(fixed))
	return printer(local).Sprintf("%v", percentage)
}

func currencyFormat(local string, el float64) string {
	lang := language.MustParse(local)
	cur, _ := currency.FromTag(lang)
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(el, number.Scale(scale))
	return printer(local).Sprintf("%v%v", currency.Symbol(cur), dec)
}

func numberFormat(local string, i int) string {
	return printer(local).Sprintf("%d", i)
}

func printer(local string) *message.Printer {
	return message.NewPrinter(message.MatchLanguage(local))
}
