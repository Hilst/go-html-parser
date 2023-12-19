package templates

import (
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"

	opt "github.com/moznion/go-optional"
)

// TYPE ENFORCER
func integer(el opt.Option[any]) opt.Option[int] {
	if el, err := el.Take(); err == nil {
		return opt.Some[int](int(el.(float64)))
	}
	return opt.None[int]()
}

func float(el opt.Option[any]) opt.Option[float64] {
	if el, err := el.Take(); err == nil {
		return opt.Some[float64](el.(float64))
	}
	return opt.None[float64]()
}

// FORMATTERS
func decimalFormat(local string, fixed int, el opt.Option[float64]) opt.Option[string] {
	if el, err := el.Take(); err == nil {
		decimal := number.Decimal(el, number.Scale(fixed))
		return opt.Some[string](printer(local).Sprintf("%v", decimal))
	}
	return opt.None[string]()
}

func percentFormat(local string, fixed int, el opt.Option[float64]) opt.Option[string] {
	if el, err := el.Take(); err == nil {
		percentage := number.Percent(el, number.Scale(fixed))
		return opt.Some[string](printer(local).Sprintf("%v", percentage))
	}
	return opt.None[string]()
}

func currencyFormat(local string, el opt.Option[float64]) opt.Option[string] {
	if el, err := el.Take(); err == nil {
		lang := language.MustParse(local)
		cur, _ := currency.FromTag(lang)
		scale, _ := currency.Cash.Rounding(cur)
		dec := number.Decimal(el, number.Scale(scale))
		return opt.Some[string](printer(local).Sprintf("%v%v", currency.Symbol(cur), dec))
	}
	return opt.None[string]()
}

func numberFormat(local string, i opt.Option[int]) opt.Option[string] {
	if i, err := i.Take(); err == nil {
		return opt.Some[string](printer(local).Sprintf("%d", i))
	}
	return opt.None[string]()
}

func printer(local string) *message.Printer {
	return message.NewPrinter(message.MatchLanguage(local))
}
