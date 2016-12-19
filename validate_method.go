package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	REGEXP_ALPHA     = "^[a-zA-Z]+$"
	REGEXP_ALPHA_NUM = "^[a-zA-Z0-9]+$"
	REGEXP_NUM       = "^[0-9]+$"
	REGEXP_EMAIL     = "^[a-zA-Z0-9]+([_\\-.][a-zA-Z0-9]+)*@[a-zA-Z0-9]+([-.][a-zA-Z0-9]+)*\\.[a-zA-Z0-9]+([-.][a-zA-Z0-9]+)*$"
)

func requireParameterCount(count int, parameters []string, rule string) {
	if len(parameters) < count {
		panic(fmt.Sprintf("validation: rule %s requires at least %d parameters.", rule, count))
	}
}

func getSize(rule string, value interface{}) (float64, error) {
	switch value.(type) {
	case int:
		return float64(value.(int)), nil
	case float64:
		return value.(float64), nil
	case string:
		return float64(len(value.(string))), nil
	default:
		return 0.0, fmt.Errorf("validation: rule %s should only be used by the value of (float or string).", rule)
	}
}

func getType(value interface{}) string {
	switch value.(type) {
	case int:
		return "float"
	case float64:
		return "float"
	case string:
		return "string"
	default:
		return "string"
	}
}

func stringTofloat64(rule, s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(fmt.Sprintf("validation: invalid parameter for rule %s, a float string is required.", rule))
	}

	return f
}

func validateAlpha(attribute string, value interface{}, parameters []string) bool {
	var strValue string

	switch value.(type) {
	case string:
		strValue, _ = value.(string)
	default:
		return false
	}

	if matched, _ := regexp.MatchString(REGEXP_ALPHA, strValue); !matched {
		return false
	}

	return true
}

func validateAlphaNum(attribute string, value interface{}, parameters []string) bool {
	var strValue string

	switch value.(type) {
	case string:
		strValue, _ = value.(string)
	default:
		return false
	}

	if matched, _ := regexp.MatchString(REGEXP_ALPHA_NUM, strValue); !matched {
		return false
	}

	return true
}

func validateBetween(attribute string, value interface{}, parameters []string) bool {
	requireParameterCount(2, parameters, "between")
	size, err := getSize("between", value)
	if err != nil {
		return false
	}

	return size >= stringTofloat64("between", parameters[0]) && size <= stringTofloat64("between", parameters[1])
}

func validateBool(attribute string, value interface{}, parameters []string) bool {
	if value == nil {
		return false
	}

	switch value.(type) {
	case bool:
		return true
	default:
		return false
	}
}

func validateEmail(attribute string, value interface{}, parameters []string) bool {
	var strValue string

	switch value.(type) {
	case string:
		strValue, _ = value.(string)
	default:
		return false
	}

	if matched, _ := regexp.MatchString(REGEXP_EMAIL, strValue); !matched {
		return false
	}

	return true
}

func validateFloat(attribute string, value interface{}, parameters []string) bool {
	if value == nil {
		return false
	}

	switch value.(type) {
	case float64:
		return true
	default:
		return false
	}
}

func validateIn(attribute string, value interface{}, parameters []string) bool {
	requireParameterCount(1, parameters, "in")

	var strValue string
	switch value.(type) {
	case int:
		strValue = strconv.Itoa(value.(int))
	case float64:
		strValue = strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case string:
		strValue = value.(string)
	default:
		return false
	}

	for _, v := range parameters {
		if strValue == v {
			return true
		}
	}

	return false
}

func validateMax(attribute string, value interface{}, parameters []string) bool {
	requireParameterCount(1, parameters, "max")

	size, err := getSize("max", value)
	if err != nil {
		return false
	}

	return size <= stringTofloat64("max", parameters[0])
}

func validateMin(attribute string, value interface{}, parameters []string) bool {
	requireParameterCount(1, parameters, "min")

	size, err := getSize("min", value)
	if err != nil {
		return false
	}

	return size >= stringTofloat64("min", parameters[0])
}

func validateNum(attribute string, value interface{}, parameters []string) bool {
	strValue, ok := value.(string)
	if !ok {
		return false
	}
	if matched, _ := regexp.MatchString(REGEXP_NUM, strValue); !matched {
		return false
	}

	return true
}

func validateRegex(attribute string, value interface{}, parameters []string) bool {
	requireParameterCount(1, parameters, "regex")

	strValue, ok := value.(string)
	if !ok {
		return false
	}
	if matched, _ := regexp.MatchString(parameters[0], strValue); !matched {
		return false
	}

	return true
}

func validateRequired(attribute string, value interface{}, parameters []string) bool {
	if value == nil {
		return false
	}

	switch value.(type) {
	case string:
		strValue, ok := value.(string)
		if !ok {
			return false
		}
		if strings.TrimSpace(strValue) == "" {
			return false
		}
	}

	return true
}

func validateSize(attribute string, value interface{}, parameters []string) bool {
	requireParameterCount(1, parameters, "size")

	size, err := getSize("size", value)
	if err != nil {
		return false
	}

	return size == stringTofloat64("size", parameters[0])
}

func validateString(attribute string, value interface{}, parameters []string) bool {
	if value == nil {
		return false
	}

	switch value.(type) {
	case string:
		return true
	default:
		return false
	}
}
