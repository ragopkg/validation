package validation

import (
	"fmt"
	"strings"
)

type validator struct {
	data    map[string]interface{}
	rules   map[string][]string
	message string
}

func New(data map[string]interface{}, rules map[string]interface{}) *validator {
	return &validator{
		data:  data,
		rules: explodeRules(rules),
	}
}

func explodeRules(rules map[string]interface{}) map[string][]string {
	r := map[string][]string{}

	for attribute, rule := range rules {
		switch rule.(type) {
		case string:
			strRule, _ := rule.(string)
			r[attribute] = strings.Split(strRule, "|")
		case []string:
			sliceRule, _ := rule.([]string)
			r[attribute] = sliceRule
		default:
			panic(fmt.Sprintf("validation: invalid rule type for attribute %s, string or []string expected.", attribute))
		}

		if !isValidRules(r[attribute]) {
			panic(fmt.Sprintf("validation: invalid rule for attribute %s", attribute))
		}
	}

	return r
}

func isValidRules(rules []string) bool {
	if len(rules) == 0 {
		return false
	}
	for _, rule := range rules {
		if rule == "" {
			return false
		}
	}

	return true
}

func (v *validator) Fails() bool {
	return !v.Passes()
}

func (v *validator) Passes() bool {
	for attribute, rules := range v.rules {
		for _, rule := range rules {
			if !v.validate(attribute, rule) {
				return false
			}
		}
	}

	return true
}

func (v *validator) GetMessage() string {
	return v.message
}

func (v *validator) validate(attribute, rule string) bool {
	rule, parameters := parseRule(rule)
	value := v.getValue(attribute)

	if rule != "required" && value == nil {
		return true
	}

	method, err := getRuleMethod(rule)
	if err != nil {
		panic(err)
	}

	// Call the method of rule.
	if !method(attribute, value, parameters) {
		var message string
		if rule == "max" || rule == "min" || rule == "size" || rule == "between" {
			message, _ = defaultRuleMessages2[rule][getType(value)]
		} else {
			message, _ = defaultRuleMessages[rule]
		}

		if rule == "size" {
			message = strings.Replace(message, ":size", parameters[0], -1)
		} else if rule == "max" {
			message = strings.Replace(message, ":max", parameters[0], -1)
		} else if rule == "min" {
			message = strings.Replace(message, ":min", parameters[0], -1)
		} else if rule == "between" {
			message = strings.Replace(message, ":min", parameters[0], -1)
			message = strings.Replace(message, ":max", parameters[1], -1)
		} else if rule == "in" {
			message = strings.Replace(message, ":values", strings.Join(parameters, ","), -1)
		}
		message = strings.Replace(message, ":attribute", attribute, -1)
		v.message = message

		return false
	}

	return true
}

func parseRule(rule string) (string, []string) {
	parameters := []string{}
	if strings.Index(rule, ":") != -1 {
		splitRule := strings.SplitN(rule, ":", 2)
		rule = splitRule[0]
		parameters = parseParameters(splitRule[0], splitRule[1])
	}

	return rule, parameters
}

func (v *validator) getValue(attribute string) interface{} {
	value, ok := v.data[attribute]
	if !ok {
		return nil
	}

	return value
}

func parseParameters(rule, parameter string) []string {
	parameters := []string{}
	if rule == "regex" {
		parameters = append(parameters, parameter)
	} else {
		parameters = strings.Split(parameter, ",")
	}

	for key, value := range parameters {
		parameters[key] = strings.TrimSpace(value)
	}

	return parameters
}
