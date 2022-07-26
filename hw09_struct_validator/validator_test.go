package hw09structvalidator

import (
	"errors"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}

	Ints struct {
		IntField   int   `validate:"min:1|max:100"`
		Int8Field  int8  `validate:"min:1|max:100"`
		Int16Field int16 `validate:"min:1|max:100"`
		Int32Field int32 `validate:"min:1|max:100"`
		Int64Field int64 `validate:"min:1|max:100"`
	}
)

// negative fields

var emptyUser = User{}

var wrongUser = User{
	ID:     "012345678",
	Age:    120,
	Email:  "wrong.email.ru",
	Role:   "tester",
	Phones: []string{"34"},
}

var wrongResponse = Response{
	Code: 420,
}

var wrongInts = Ints{
	IntField:   0,
	Int8Field:  0,
	Int16Field: 0,
	Int32Field: 0,
	Int64Field: 100999,
}

// end negative fields

// positive fields

var user = User{
	ID:     "1",
	Age:    20,
	Email:  "test@yandex.ru",
	Role:   "stuff",
	Phones: []string{"79200000000", "79100000011"},
}

var app = App{
	Version: "200",
}

var token = Token{
	Header:    []byte("dfdsfs"),
	Payload:   []byte("dfdsfs"),
	Signature: []byte("dfdsfs"),
}

var response = Response{
	Code: 200,
}

// end positive fields

func TestValidate(t *testing.T) {
	pattern := "^\\w+@\\w+\\.\\w+$"
	emailPattern := regexp.MustCompile(pattern)
	roles := []string{"admin", "stuff"}

	tests := []struct {
		name        string
		in          interface{}
		expectedErr error
	}{
		{
			name: "empty user",
			in:   emptyUser,
			expectedErr: ValidationErrors{
				ValidationError{Field: "ID", Err: &ErrStrLen{0, 36}},
				ValidationError{Field: "Age", Err: &ErrIntMin{0, 18}},
				ValidationError{Field: "Email", Err: &ErrStrRegexp{"", emailPattern}},
				ValidationError{Field: "Role", Err: &ErrStrIn{"", roles}},
			},
		},
		{
			name: "incorrect user data",
			in:   wrongUser,
			expectedErr: ValidationErrors{
				ValidationError{Field: "ID", Err: &ErrStrLen{len(wrongUser.ID), 36}},
				ValidationError{Field: "Age", Err: &ErrIntMax{int64(wrongUser.Age), 50}},
				ValidationError{Field: "Email", Err: &ErrStrRegexp{wrongUser.Email, emailPattern}},
				ValidationError{Field: "Role", Err: &ErrStrIn{string(wrongUser.Role), roles}},
				ValidationError{Field: "Phones", Err: &ErrStrLen{len(wrongUser.Phones[0]), 11}},
			},
		},
		{
			name: "incorrect response",
			in:   wrongResponse,
			expectedErr: ValidationErrors{
				ValidationError{Field: "Code", Err: &ErrIntIn{int64(wrongResponse.Code), []int64{200, 404, 500}}},
			},
		},
		{
			name: "incorrect ints",
			in:   wrongInts,
			expectedErr: ValidationErrors{
				ValidationError{Field: "IntField", Err: &ErrIntMin{int64(wrongInts.IntField), 1}},
				ValidationError{Field: "Int8Field", Err: &ErrIntMin{int64(wrongInts.Int8Field), 1}},
				ValidationError{Field: "Int16Field", Err: &ErrIntMin{int64(wrongInts.Int16Field), 1}},
				ValidationError{Field: "Int32Field", Err: &ErrIntMin{int64(wrongInts.Int32Field), 1}},
				ValidationError{Field: "Int64Field", Err: &ErrIntMax{wrongInts.Int64Field, 100}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.in)
			require.Equal(t, tt.expectedErr, err)
			_ = tt
		})
	}
}

func TestErrIncorrectUse(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
	}{
		{
			name: "incorrect kind",
			in:   "text",
		},
		{
			name: "incorrect field type",
			in: struct {
				Height float64 `validate:"min:10"`
			}{},
		},
		{
			name: "unknown int rule",
			in: struct {
				Value int `validate:"dfd:dfdg"`
			}{},
		},
		{
			name: "unknown string rule",
			in: struct {
				Text string `validate:"test:true"`
			}{},
		},
		{
			name: "incorrect string len condition",
			in: struct {
				Value string `validate:"len:test"`
			}{},
		},
		{
			name: "incorrect string regexp condition",
			in: struct {
				Value string `validate:"regexp:+"`
			}{},
		},
		{
			name: "incorrect int min condition",
			in: struct {
				Value int `validate:"min:oops"`
			}{},
		},
		{
			name: "incorrect int no value condition",
			in: struct {
				Value int `validate:"min"`
			}{},
		},
		{
			name: "incorrect int in condition",
			in: struct {
				Value int `validate:"in:200,-,500"`
			}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.in)
			require.True(t, errors.Is(err, &ErrIncorrectUse{}))
		})
	}
}

func BenchmarkValidateErrors(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Validate(emptyUser)
			_ = Validate(wrongUser)
			_ = Validate(wrongResponse)
		}
	})
}

func BenchmarkValidateSuccess(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Validate(user)
			_ = Validate(app)
			_ = Validate(token)
			_ = Validate(response)
		}
	})
}
