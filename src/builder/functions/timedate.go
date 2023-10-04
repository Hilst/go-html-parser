package functions

import "time"

// GENERATORS
func now() time.Time {
	return time.Now()
}

func timedate(f string, s string) time.Time {
	dt, _ := time.Parse(mapDateFormatter(f), s)
	return dt
}

// TRANSFORMERS
func dateFormat(f string, dt time.Time) string {
	return dt.Format(mapDateFormatter(f))
}
