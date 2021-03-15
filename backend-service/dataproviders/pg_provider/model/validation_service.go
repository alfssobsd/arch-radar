package model

import (
	"unicode/utf8"
)

const (
	ErrEmptyValue = "empty"
	ErrMaxLength  = "len"
	ErrWrongValue = "value"
)

func (m Service) Validate() (errors map[string]string, valid bool) {
	errors = map[string]string{}
	if utf8.RuneCountInString(m.Title) > 128 {
		errors[Columns.Service.Title] = ErrMaxLength
	}

	return errors, len(errors) == 0
}
