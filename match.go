package validator

import (
	"regexp"
)

const (
	EMAIL_PATTERN = "([a-zA-Z0-9])+(@)([a-zA-Z0-9])+((.)[a-zA-Z0-9])+"
)

type Match struct {
	V *Validator
}

func (m *Match) validate(v *Validator, pattern string, msg ...string) {
	if str, ok := v.Object.(string); ok {
		matched, err := regexp.MatchString(pattern, str)
		if err != nil || matched == false {
			v.Add("pattern mismatch", msg...)
		} else {
			v.Add("matching cannot be applied on this object")
		}
	} else {
		v.Add("cannot be applied on this object")
	}
}
