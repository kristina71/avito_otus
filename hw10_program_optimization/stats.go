package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const bufferSize = 512

//easyjson:json
type UserEmail struct {
	Email string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	if domain == "" {
		return nil, fmt.Errorf("domain must not be empty")
	}

	result := DomainStat{}
	domainSuffix := "." + domain

	scanner := bufio.NewScanner(r)
	buffer := make([]byte, bufferSize)
	scanner.Buffer(buffer, 3*bufferSize)
	userEmail := &UserEmail{}
	for scanner.Scan() {
		*userEmail = UserEmail{}
		if err := userEmail.UnmarshalJSON(scanner.Bytes()); err != nil {
			return nil, err
		}

		email := userEmail.Email
		if strings.HasSuffix(email, domainSuffix) {
			idx := strings.IndexRune(email, '@')
			if idx < 0 {
				return nil, fmt.Errorf("invalid email format: %s is not valid (should contain \"@\" symbol)", userEmail.Email)
			}
			result[strings.ToLower(email[idx+1:])]++
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading error: %w", err)
	}
	return result, nil
}
