package validation

import (
	"errors"
)

type ruleMethod func(string, interface{}, []string) bool

var ruleMethodMap = map[string]ruleMethod{
	"alpha":     validateAlpha,
	"alpha_num": validateAlphaNum,
	"between":   validateBetween,
	"bool":      validateBool,
	"email":     validateEmail,
	"float":     validateFloat,
	"in":        validateIn,
	"max":       validateMax,
	"min":       validateMin,
	"num":       validateNum,
	"regex":     validateRegex,
	"required":  validateRequired,
	"size":      validateSize,
	"string":    validateString,
}

var defaultRuleMessages = map[string]string{
	"alpha":     "The :attribute may only contain letters.",
	"alpha_num": "The :attribute may only contain letters and numbers.",
	"bool":      "The :attribute field must be true or false.",
	"email":     "The :attribute must be a valid email address.",
	"float":     "The :attribute must be a float.",
	"in":        "The :attribute field must one of (:values).",
	"num":       "The :attribute may only contain numbers.",
	"regex":     "The :attribute format is invalid.",
	"required":  "The :attribute field is required.",
	"string":    "The :attribute must be a string.",
}

var defaultRuleMessages2 = map[string]map[string]string{
	"between": {
		"float":  "The :attribute must be between :min and :max.",
		"string": "The :attribute must be between :min and :max characters.",
	},
	"max": {
		"float":  "The :attribute may not be greater than :max.",
		"string": "The :attribute may not be greater than :max characters.",
	},
	"min": {
		"float":  "The :attribute must be at least :min.",
		"string": "The :attribute must be at least :min characters.",
	},
	"size": {
		"float":  "The :attribute must be :size.",
		"string": "The :attribute must be :size characters.",
	},
}

func getRuleMethod(rule string) (ruleMethod, error) {
	method, ok := ruleMethodMap[rule]
	if !ok {
		return nil, errors.New("validation: rule " + rule + " not supported.")
	}

	return method, nil
}
