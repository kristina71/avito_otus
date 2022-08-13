package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

type substr struct {
	prev      *substr
	index     int
	str       string
	char      rune
	protected bool
}

const Backslash = 92

func (s *substr) checkString() error {
	if s.prev == nil {
		if unicode.IsDigit(s.char) {
			return ErrInvalidString
		}

		if s.char != Backslash {
			s.str = string(s.char)
		}
		return nil
	}

	if s.prev.char == Backslash && !s.prev.protected {
		s.protected = true
	}

	if unicode.IsDigit(s.char) {
		if !s.prev.protected && unicode.IsDigit(s.prev.char) {
			return ErrInvalidString
		}

		if s.protected {
			s.str = string(s.char)
			return nil
		}
		s.prev.index = int(s.char - '0')
		return nil
	}

	if unicode.IsLetter(s.char) && s.protected {
		s.str = "\\" + string(s.char)
		return nil
	}

	if s.char == Backslash && !s.protected {
		return nil
	}

	s.str = string(s.char)
	return nil
}

func (s *substr) buildString() string {
	var b strings.Builder
	for i := 0; i < s.index; i++ {
		b.WriteString(s.str)
	}
	return b.String()
}

func Unpack(str string) (string, error) {
	runes := []rune(str)
	newSubstr := make([]*substr, 0)

	for i, r := range runes {
		var prev *substr
		if i > 0 {
			prev = newSubstr[i-1]
		}
		newSubstr = append(newSubstr, &substr{
			prev:  prev,
			index: 1,
			char:  r,
		})
	}

	var b strings.Builder
	for _, s := range newSubstr {
		if err := s.checkString(); err != nil {
			return "", err
		}
	}

	for _, t := range newSubstr {
		b.WriteString(t.buildString())
	}

	return b.String(), nil
}
