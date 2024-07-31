package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"
)

type UserRole string

type (
	Address struct {
		City    string
		Country string `validate:"in:USA,Canada"`
	}

	User struct {
		ID      string `json:"id" validate:"len:36"`
		Name    string
		Age     int             `validate:"min:18|max:50"`
		Email   string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role    UserRole        `validate:"in:admin,stuff"`
		Phones  []string        `validate:"len:11"`
		meta    json.RawMessage //nolint:unused
		Address Address         `validate:"nested"`
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
		Code    int      `validate:"in:200,404,500"`
		Body    string   `json:"omitempty"`
		Strings []string `validate:"len:11"`
	}
)

func TestValidateUser(t *testing.T) {
	tests := []struct {
		in          User
		expectedErr error
	}{
		{
			in: User{
				ID:     "12345678-1234-1234-1234-123456789012",
				Name:   "John",
				Age:    25,
				Email:  "johndoe@example.com",
				Role:   "admin",
				Phones: []string{"12345678901"},
				Address: Address{
					City:    "Boston",
					Country: "USA",
				},
			},
			expectedErr: nil,
		},
		{
			in: User{
				ID:     "123",
				Name:   "John",
				Age:    25,
				Email:  "johndoe@example.com",
				Role:   "admin",
				Phones: []string{"12345678901"},
				Address: Address{
					City:    "Boston",
					Country: "USA",
				},
			},
			expectedErr: ValidationErrors{{Field: "ID", Err: fmt.Errorf("value 123 is not of a length 36")}},
		},
		// Добавьте остальные тесты для структуры User сюда
	}

	runTests(t, tests)
}

func TestValidateApp(t *testing.T) {
	tests := []struct {
		in          App
		expectedErr error
	}{
		{
			in:          App{Version: "1.0.0"},
			expectedErr: nil,
		},
		{
			in:          App{Version: "1.0.00"},
			expectedErr: ValidationErrors{{Field: "Version", Err: fmt.Errorf("value 1.0.00 is not of a length 5")}},
		},
	}

	runTests(t, tests)
}

func TestValidateResponse(t *testing.T) {
	tests := []struct {
		in          Response
		expectedErr error
	}{
		{
			in:          Response{Code: 200, Body: "hey", Strings: []string{"12345678901"}},
			expectedErr: nil,
		},
		{
			in:          Response{Code: 201, Body: "hey", Strings: []string{"12345678901"}},
			expectedErr: ValidationErrors{{Field: "Code", Err: fmt.Errorf("value 201 not in required [200 404 500]")}},
		},
	}

	runTests(t, tests)
}

func runTests[T any](t *testing.T, tests []struct {
	in          T
	expectedErr error
},
) {
	t.Helper()
	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()
			err := Validate(tt.in)
			switch {
			case tt.expectedErr != nil && err == nil:
				t.Errorf("ожидалась ошибка: %v, но получена nil", tt.expectedErr)
			case tt.expectedErr == nil && err != nil:
				t.Errorf("ожидалось отсутствие ошибок, но получена ошибка: %v", err)
			case tt.expectedErr != nil && err != nil && tt.expectedErr.Error() != err.Error():
				t.Errorf("ожидалась ошибка: %v, но получена: %v", tt.expectedErr, err)
			}
		})
	}
}
