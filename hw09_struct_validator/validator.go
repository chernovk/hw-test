package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrNotAStruct      = errors.New("not a structure received")
	ErrUnsupportedType = errors.New("type of field no supported")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	errorMessagse := ""
	for _, err := range v {
		errorMessagse += err.Field + ":" + err.Err.Error() + "\n"
	}
	return errorMessagse
}

type ValueErrors []error

func (v ValueErrors) Error() string {
	errorMessagse := ""
	for _, err := range v {
		errorMessagse += err.Error()
	}
	return errorMessagse
}

type Numeric interface {
	~int | ~int64 | ~float32 | ~float64
}

func convertToFloat64(value any) (float64, error) {
	switch v := value.(type) {
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		val, err := strconv.ParseFloat(value.(string), 64)
		if err != nil {
			return 0, ErrUnsupportedType
		}
		return val, nil
	default:
		return 0, ErrUnsupportedType
	}
}

func convertToString(value interface{}) string {
	switch v := value.(type) {
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", value)
	}
}

func inCheck(valueRecieved interface{}, valuesRequired []string) error {
	stringedValue := convertToString(valueRecieved)

	found := false
	for _, valueRequired := range valuesRequired {
		if stringedValue == valueRequired {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("value %v not in required %v", valueRecieved, valuesRequired)
	}
	return nil
}

func InCheck(valuesRecieved interface{}, valuesRequired []string) ValueErrors {
	var valueErrors ValueErrors

	rV := reflect.ValueOf(valuesRecieved)
	if rV.Kind() == reflect.Slice {
		for i := 0; i < rV.Len(); i++ {
			elem := rV.Index(i).Interface()
			if err := inCheck(elem, valuesRequired); err != nil {
				valueErrors = append(valueErrors, err)
			}
		}
	} else if err := inCheck(valuesRecieved, valuesRequired); err != nil {
		valueErrors = append(valueErrors, err)
	}
	return valueErrors
}

func compareCheck(valueRecieved interface{}, valueCompareFloat64 float64, caseName string) error {
	if _, ok := valueRecieved.(string); ok {
		return fmt.Errorf("value %v: %w", valueRecieved, ErrUnsupportedType)
	}

	valueRecievedFloat64, err := convertToFloat64(valueRecieved)
	if err != nil {
		return fmt.Errorf("value %v: %w", valueRecieved, err)
	}
	if valueRecievedFloat64 > valueCompareFloat64 && caseName == "max" {
		return fmt.Errorf("value %v greater then required %v", valueRecieved, valueCompareFloat64)
	} else if valueRecievedFloat64 < valueCompareFloat64 && caseName == "min" {
		return fmt.Errorf("value %v less then required %v", valueRecieved, valueCompareFloat64)
	}
	return nil
}

func CompareCheck(valuesRecieved interface{}, valueCompareFloat64 float64, caseName string) ValueErrors {
	var valueErrors ValueErrors

	rV := reflect.ValueOf(valuesRecieved)
	if rV.Kind() == reflect.Slice {
		for i := 0; i < rV.Len(); i++ {
			elem := rV.Index(i).Interface()
			if err := compareCheck(elem, valueCompareFloat64, caseName); err != nil {
				valueErrors = append(valueErrors, err)
			}
		}
	} else if err := compareCheck(valuesRecieved, valueCompareFloat64, caseName); err != nil {
		valueErrors = append(valueErrors, err)
	}
	return valueErrors
}

func lenCheck(valueRecieved interface{}, lengthRequired int) error {
	stringRecieved, ok := valueRecieved.(string)
	if !ok {
		return fmt.Errorf("value %v: %w", valueRecieved, ErrUnsupportedType)
	}
	if len([]rune(stringRecieved)) != lengthRequired {
		return fmt.Errorf("value %v is not of a length %v", valueRecieved, lengthRequired)
	}
	return nil
}

func LenCheck(valuesRecieved interface{}, lengthRequired int) ValueErrors {
	var valueErrors ValueErrors

	rV := reflect.ValueOf(valuesRecieved)
	if rV.Kind() == reflect.Slice {
		for _, valueRecieved := range valuesRecieved.([]string) {
			if err := lenCheck(valueRecieved, lengthRequired); err != nil {
				valueErrors = append(valueErrors, err)
			}
		}
	} else if err := lenCheck(valuesRecieved, lengthRequired); err != nil {
		valueErrors = append(valueErrors, err)
	}
	return valueErrors
}

func regexpCheck(valueRecieved interface{}, re *regexp.Regexp) error {
	stringRecieved, ok := valueRecieved.(string)
	if !ok {
		return fmt.Errorf("value %v: %w", valueRecieved, ErrUnsupportedType)
	}

	if !re.MatchString(stringRecieved) {
		return fmt.Errorf("value %v does not match the pattern %v", valueRecieved, re.String())
	}
	return nil
}

func RegexpCheck(valuesRecieved interface{}, re *regexp.Regexp) ValueErrors {
	var valueErrors ValueErrors

	rV := reflect.ValueOf(valuesRecieved)
	if rV.Kind() == reflect.Slice {
		for _, valueRecieved := range valuesRecieved.([]string) {
			if err := regexpCheck(valueRecieved, re); err != nil {
				valueErrors = append(valueErrors, err)
			}
		}
	} else if err := regexpCheck(valuesRecieved, re); err != nil {
		valueErrors = append(valueErrors, err)
	}
	return valueErrors
}

func processValue(fieldName, requirement string, value interface{}) ValidationError {
	reqParts := strings.Split(requirement, ":")
	reqName := reqParts[0]
	reqValues := reqParts[1]

	var fieldErrors ValueErrors

	switch reqName {
	case "in":
		reqValuesSplited := strings.Split(reqValues, ",")
		fieldErrors = InCheck(value, reqValuesSplited)
	case "max", "min":
		valueCompareFloat64, err := convertToFloat64(reqValues)
		if err != nil {
			fieldErrors = ValueErrors{fmt.Errorf("value %v: %w", reqValues, err)}
			break
		}
		fieldErrors = CompareCheck(value, valueCompareFloat64, reqName)
	case "len":
		intLengthRequired, err := strconv.Atoi(reqValues)
		if err != nil {
			fieldErrors = ValueErrors{fmt.Errorf("incorrect validate tag len: %v", reqValues)}
			break
		}
		fieldErrors = LenCheck(value, intLengthRequired)
	case "regexp":
		re, err := regexp.Compile(reqValues)
		if err != nil {
			fieldErrors = ValueErrors{fmt.Errorf("invalid regex pattern: %w", err)}
			break
		}
		fieldErrors = RegexpCheck(value, re)
	}
	if len(fieldErrors) != 0 {
		return ValidationError{
			Field: fieldName,
			Err:   fieldErrors,
		}
	}
	return ValidationError{}
}

func Validate(v interface{}) error {
	var validationErrors ValidationErrors

	rV := reflect.ValueOf(v)
	if rV.Kind() != reflect.Struct {
		return ErrNotAStruct
	}

	rT := rV.Type()
	for i := 0; i < rT.NumField(); i++ {
		currentField := rT.Field(i)

		if !unicode.IsUpper(rune(currentField.Name[0])) {
			continue
		}

		currentFieldValidateTag := currentField.Tag.Get("validate")
		if currentFieldValidateTag == "" {
			continue
		}
		currentFieldValue := rV.Field(i).Interface()

		if currentFieldValidateTag == "nested" && !reflect.ValueOf(currentFieldValue).IsZero() {
			var nestedValidationErrors ValidationErrors
			if nestedErr := Validate(currentFieldValue); errors.As(nestedErr, &nestedValidationErrors) {
				for _, nve := range nestedValidationErrors {
					validationErrors = append(validationErrors, ValidationError{
						Field: currentField.Name + ":" + nve.Field,
						Err:   nve.Err,
					})
				}
			}
			continue
		}

		requirements := strings.Split(currentFieldValidateTag, "|")
		for _, requirement := range requirements {
			if errs := processValue(currentField.Name, requirement, currentFieldValue); errs != (ValidationError{}) {
				validationErrors = append(validationErrors, errs)
			}
		}
	}

	if len(validationErrors) == 0 {
		return nil
	}

	return validationErrors
}
