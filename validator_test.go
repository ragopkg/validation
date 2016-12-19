package validation

import (
	"testing"
)

func TestRules(t *testing.T) {
	tests := map[string]struct {
		data  map[string]interface{}
		rules map[string]interface{}
		pass  bool
	}{
		// alpha rule
		// ======================================================
		"alpha-true1": {
			map[string]interface{}{"foo": "apz"},
			map[string]interface{}{"foo": "alpha"},
			true,
		},
		"alpha-true2": {
			map[string]interface{}{"foo": "aPz"},
			map[string]interface{}{"foo": "alpha"},
			true,
		},

		"alpha-false1": {
			map[string]interface{}{"foo": false},
			map[string]interface{}{"foo": "alpha"},
			false,
		},
		"alpha-false2": {
			map[string]interface{}{"foo": 1},
			map[string]interface{}{"foo": "alpha"},
			false,
		},
		"alpha-false3": {
			map[string]interface{}{"foo": 1.0},
			map[string]interface{}{"foo": "alpha"},
			false,
		},
		"alpha-false4": {
			map[string]interface{}{"foo": "az09"},
			map[string]interface{}{"foo": "alpha"},
			false,
		},
		"alpha-false5": {
			map[string]interface{}{"foo": "09"},
			map[string]interface{}{"foo": "alpha"},
			false,
		},

		// alpha_num rule
		// ======================================================
		"alpha_num-true1": {
			map[string]interface{}{"foo": "apz"},
			map[string]interface{}{"foo": "alpha_num"},
			true,
		},
		"alpha_num-true2": {
			map[string]interface{}{"foo": "aPz"},
			map[string]interface{}{"foo": "alpha_num"},
			true,
		},
		"alpha_num-true3": {
			map[string]interface{}{"foo": "aPz09"},
			map[string]interface{}{"foo": "alpha_num"},
			true,
		},
		"alpha_num-true4": {
			map[string]interface{}{"foo": "09"},
			map[string]interface{}{"foo": "alpha_num"},
			true,
		},

		"alpha_num-false1": {
			map[string]interface{}{"foo": false},
			map[string]interface{}{"foo": "alpha_num"},
			false,
		},
		"alpha_num-false2": {
			map[string]interface{}{"foo": 1},
			map[string]interface{}{"foo": "alpha_num"},
			false,
		},
		"alpha_num-false3": {
			map[string]interface{}{"foo": 1.0},
			map[string]interface{}{"foo": "alpha_num"},
			false,
		},
		"alpha_num-false4": {
			map[string]interface{}{"foo": "az09."},
			map[string]interface{}{"foo": "alpha_num"},
			false,
		},
		"alpha_num-false5": {
			map[string]interface{}{"foo": "@$"},
			map[string]interface{}{"foo": "alpha_num"},
			false,
		},

		// between rule
		// ======================================================
		"between-true1": {
			map[string]interface{}{"foo": 1},
			map[string]interface{}{"foo": "between:1,3"},
			true,
		},
		"between-true2": {
			map[string]interface{}{"foo": 3},
			map[string]interface{}{"foo": "between:1,3"},
			true,
		},
		"between-true3": {
			map[string]interface{}{"foo": 1.0},
			map[string]interface{}{"foo": "between:1,3"},
			true,
		},
		"between-true4": {
			map[string]interface{}{"foo": 3.0},
			map[string]interface{}{"foo": "between:1,3"},
			true,
		},
		"between-true5": {
			map[string]interface{}{"foo": "a"},
			map[string]interface{}{"foo": "between:1,3"},
			true,
		},
		"between-true6": {
			map[string]interface{}{"foo": "aPz"},
			map[string]interface{}{"foo": "between:1,3"},
			true,
		},

		"between-false1": {
			map[string]interface{}{"foo": 3},
			map[string]interface{}{"foo": "between:1,2"},
			false,
		},
		"between-false2": {
			map[string]interface{}{"foo": "aPz"},
			map[string]interface{}{"foo": "between:1,2"},
			false,
		},

		// bool rule
		// ======================================================
		"bool-true1": {
			map[string]interface{}{"foo": false},
			map[string]interface{}{"foo": "bool"},
			true,
		},
		"bool-true2": {
			map[string]interface{}{"foo": true},
			map[string]interface{}{"foo": "bool"},
			true,
		},

		"bool-false1": {
			map[string]interface{}{"foo": 1},
			map[string]interface{}{"foo": "bool"},
			false,
		},
		"bool-false2": {
			map[string]interface{}{"foo": 1.0},
			map[string]interface{}{"foo": "bool"},
			false,
		},
		"bool-false3": {
			map[string]interface{}{"foo": "a"},
			map[string]interface{}{"foo": "bool"},
			false,
		},

		// email rule
		// ======================================================
		"email-true1": {
			map[string]interface{}{"foo": "abc@x.com"},
			map[string]interface{}{"foo": "email"},
			true,
		},
		"email-true2": {
			map[string]interface{}{"foo": "a-b_c.z@x.com.cn"},
			map[string]interface{}{"foo": "email"},
			true,
		},

		"email-false1": {
			map[string]interface{}{"foo": 1},
			map[string]interface{}{"foo": "email"},
			false,
		},
		"email-false2": {
			map[string]interface{}{"foo": 1.0},
			map[string]interface{}{"foo": "email"},
			false,
		},
		"email-false3": {
			map[string]interface{}{"foo": "ax.com"},
			map[string]interface{}{"foo": "email"},
			false,
		},
		"email-false4": {
			map[string]interface{}{"foo": "abc@x"},
			map[string]interface{}{"foo": "email"},
			false,
		},

		// float rule
		// ======================================================
		"float-true1": {
			map[string]interface{}{"foo": 0.0},
			map[string]interface{}{"foo": "float"},
			true,
		},
		"float-true2": {
			map[string]interface{}{"foo": -1.0},
			map[string]interface{}{"foo": "float"},
			true,
		},

		"float-false1": {
			map[string]interface{}{"foo": true},
			map[string]interface{}{"foo": "float"},
			false,
		},
		"float-false2": {
			map[string]interface{}{"foo": 1},
			map[string]interface{}{"foo": "float"},
			false,
		},
		"float-false3": {
			map[string]interface{}{"foo": "a"},
			map[string]interface{}{"foo": "float"},
			false,
		},

		// in rule
		// ======================================================
		"in-true1": {
			map[string]interface{}{"foo": 1},
			map[string]interface{}{"foo": "in:1,3"},
			true,
		},
		"in-true2": {
			map[string]interface{}{"foo": 3},
			map[string]interface{}{"foo": "in:1,3"},
			true,
		},
		"in-true3": {
			map[string]interface{}{"foo": 1.0},
			map[string]interface{}{"foo": "in:1,3"},
			true,
		},
		"in-true4": {
			map[string]interface{}{"foo": 3.0},
			map[string]interface{}{"foo": "in:1,3"},
			true,
		},
		"in-true5": {
			map[string]interface{}{"foo": 3.1},
			map[string]interface{}{"foo": "in:1,3.1"},
			true,
		},
		"in-true6": {
			map[string]interface{}{"foo": "a"},
			map[string]interface{}{"foo": "in:a,b"},
			true,
		},

		"in-false1": {
			map[string]interface{}{"foo": 3},
			map[string]interface{}{"foo": "in:1,2"},
			false,
		},
		"in-false2": {
			map[string]interface{}{"foo": "a"},
			map[string]interface{}{"foo": "in:1,2"},
			false,
		},

		// max rule
		// ======================================================
		"max-true1": {
			map[string]interface{}{"foo": 3},
			map[string]interface{}{"foo": "max:3"},
			true,
		},
		"max-true2": {
			map[string]interface{}{"foo": 3.0},
			map[string]interface{}{"foo": "max:3.0"},
			true,
		},
		"max-true3": {
			map[string]interface{}{"foo": 3.1},
			map[string]interface{}{"foo": "max:3.1"},
			true,
		},
		"max-true4": {
			map[string]interface{}{"foo": "abc"},
			map[string]interface{}{"foo": "max:3"},
			true,
		},

		"max-false1": {
			map[string]interface{}{"foo": 3},
			map[string]interface{}{"foo": "max:1"},
			false,
		},
		"max-false2": {
			map[string]interface{}{"foo": 3.0},
			map[string]interface{}{"foo": "max:1"},
			false,
		},
		"max-false3": {
			map[string]interface{}{"foo": 3.0},
			map[string]interface{}{"foo": "max:1.0"},
			false,
		},
		"max-false4": {
			map[string]interface{}{"foo": "abc"},
			map[string]interface{}{"foo": "max:1"},
			false,
		},

		// min rule
		// ======================================================
		"min-true1": {
			map[string]interface{}{"foo": 3},
			map[string]interface{}{"foo": "min:3"},
			true,
		},
		"min-true2": {
			map[string]interface{}{"foo": 3.0},
			map[string]interface{}{"foo": "min:3.0"},
			true,
		},
		"min-true3": {
			map[string]interface{}{"foo": 3.1},
			map[string]interface{}{"foo": "min:3.1"},
			true,
		},
		"min-true4": {
			map[string]interface{}{"foo": "abc"},
			map[string]interface{}{"foo": "min:3"},
			true,
		},

		"min-false1": {
			map[string]interface{}{"foo": 1},
			map[string]interface{}{"foo": "min:3"},
			false,
		},
		"min-false2": {
			map[string]interface{}{"foo": 1.0},
			map[string]interface{}{"foo": "min:3"},
			false,
		},
		"min-false3": {
			map[string]interface{}{"foo": 1.0},
			map[string]interface{}{"foo": "min:3.0"},
			false,
		},
		"min-false4": {
			map[string]interface{}{"foo": "a"},
			map[string]interface{}{"foo": "min:3"},
			false,
		},

		// num rule
		// ======================================================
		"num-true1": {
			map[string]interface{}{"foo": "09"},
			map[string]interface{}{"foo": "num"},
			true,
		},
		"num-true2": {
			map[string]interface{}{"foo": "0123456789"},
			map[string]interface{}{"foo": "num"},
			true,
		},

		"num-false1": {
			map[string]interface{}{"foo": false},
			map[string]interface{}{"foo": "num"},
			false,
		},
		"num-false2": {
			map[string]interface{}{"foo": 1},
			map[string]interface{}{"foo": "num"},
			false,
		},
		"num-false3": {
			map[string]interface{}{"foo": 1.0},
			map[string]interface{}{"foo": "num"},
			false,
		},
		"num-false4": {
			map[string]interface{}{"foo": "az09."},
			map[string]interface{}{"foo": "num"},
			false,
		},
		"num-false5": {
			map[string]interface{}{"foo": "@$"},
			map[string]interface{}{"foo": "num"},
			false,
		},

		// regex rule
		// ======================================================
		"regex-true1": {
			map[string]interface{}{"foo": "123"},
			map[string]interface{}{"foo": "regex:[0-9]+"},
			true,
		},
		"regex-true2": {
			map[string]interface{}{"foo": "a1c"},
			map[string]interface{}{"foo": "regex:[0-9]+"},
			true,
		},
		"regex-false1": {
			map[string]interface{}{"foo": true},
			map[string]interface{}{"foo": "regex:[0-9]+"},
			false,
		},
		"regex-false2": {
			map[string]interface{}{"foo": 1},
			map[string]interface{}{"foo": "regex:[0-9]+"},
			false,
		},
		"regex-false3": {
			map[string]interface{}{"foo": 3.0},
			map[string]interface{}{"foo": "regex:[0-9]+"},
			false,
		},

		// required rule
		// ======================================================
		"required-true1": {
			map[string]interface{}{"foo": false},
			map[string]interface{}{"foo": "required"},
			true,
		},
		"required-true2": {
			map[string]interface{}{"foo": 1},
			map[string]interface{}{"foo": "required"},
			true,
		},
		"required-true3": {
			map[string]interface{}{"foo": 1.0},
			map[string]interface{}{"foo": "required"},
			true,
		},
		"required-true4": {
			map[string]interface{}{"foo": "bar"},
			map[string]interface{}{"foo": "required"},
			true,
		},

		"required-false1": {
			map[string]interface{}{"foo": nil},
			map[string]interface{}{"foo": "required"},
			false,
		},
		"required-false2": {
			map[string]interface{}{"bar": "foo"},
			map[string]interface{}{"foo": "required"},
			false,
		},

		// size rule
		// ======================================================
		"size-true1": {
			map[string]interface{}{"foo": 3},
			map[string]interface{}{"foo": "size:3"},
			true,
		},
		"size-true2": {
			map[string]interface{}{"foo": 3.0},
			map[string]interface{}{"foo": "size:3.0"},
			true,
		},
		"size-true3": {
			map[string]interface{}{"foo": 3.1},
			map[string]interface{}{"foo": "size:3.1"},
			true,
		},
		"size-true4": {
			map[string]interface{}{"foo": "abc"},
			map[string]interface{}{"foo": "size:3"},
			true,
		},
		"size-true5": {
			map[string]interface{}{"foo": "abc"},
			map[string]interface{}{"foo": "size:3.0"},
			true,
		},

		"size-false1": {
			map[string]interface{}{"foo": 1},
			map[string]interface{}{"foo": "size:3"},
			false,
		},
		"size-false2": {
			map[string]interface{}{"foo": 1.0},
			map[string]interface{}{"foo": "size:3"},
			false,
		},
		"size-false3": {
			map[string]interface{}{"foo": 1.0},
			map[string]interface{}{"foo": "size:3.0"},
			false,
		},
		"size-false4": {
			map[string]interface{}{"foo": "a"},
			map[string]interface{}{"foo": "size:3"},
			false,
		},
		"size-false5": {
			map[string]interface{}{"foo": "a"},
			map[string]interface{}{"foo": "size:3.0"},
			false,
		},

		// string rule
		// ======================================================
		"string-true1": {
			map[string]interface{}{"foo": "aPz"},
			map[string]interface{}{"foo": "string"},
			true,
		},
		"string-true2": {
			map[string]interface{}{"foo": ""},
			map[string]interface{}{"foo": "string"},
			true,
		},

		"string-false1": {
			map[string]interface{}{"foo": true},
			map[string]interface{}{"foo": "string"},
			false,
		},
		"string-false2": {
			map[string]interface{}{"foo": 1},
			map[string]interface{}{"foo": "string"},
			false,
		},
		"string-false3": {
			map[string]interface{}{"foo": 1.0},
			map[string]interface{}{"foo": "string"},
			false,
		},

		// ======================================================
		// multi rules
		// ======================================================
		"multi-true1": {
			map[string]interface{}{"foo": "aPz"},
			map[string]interface{}{"foo": "required|string|max:5"},
			true,
		},
		"multi-true2": {
			map[string]interface{}{"foo": "aPz"},
			map[string]interface{}{"foo": "string|max:5"},
			true,
		},

		"multi-false1": {
			map[string]interface{}{"foo": "aPz"},
			map[string]interface{}{"foo": "required|string|max:2"},
			false,
		},
		"multi-false2": {
			map[string]interface{}{"foo": "aPz"},
			map[string]interface{}{"foo": "string|max:2"},
			false,
		},
	}

	for i, tt := range tests {
		validator := New(tt.data, tt.rules)
		if validator.Passes() != tt.pass {
			t.Errorf("Test %s failed", i)
		}
	}
}
