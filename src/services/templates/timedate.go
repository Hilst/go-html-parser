package templates

import (
	"time"

	opt "github.com/moznion/go-optional"
	DateFormatExchange "github.com/newm4n/go-dfe"
)

// GENERATORS
func now() opt.Option[time.Time] {
	return opt.Some[time.Time](time.Now())
}

func timedate(f string, s opt.Option[string]) opt.Option[time.Time] {
	if s, err := s.Take(); err == nil {
		translation := DateFormatExchange.NewPatternTranslation()
		dt, _ := time.Parse(translation.JavaToGoFormat(f), s)
		return opt.Some[time.Time](dt)
	}
	return opt.None[time.Time]()
}

// TRANSFORMERS
func dateFormat(f string, dt opt.Option[time.Time]) opt.Option[string] {
	if dt, err := dt.Take(); err == nil {
		translation := DateFormatExchange.NewPatternTranslation()
		return opt.Some[string](dt.Format(translation.JavaToGoFormat(f)))
	}
	return opt.None[string]()
}
