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

		// required rule
		// ======================================================
		"required-bool": {
			map[string]interface{}{"foo": false},
			map[string]interface{}{"foo": "required"},
			true,
		},
		"required-int": {
			map[string]interface{}{"foo": 1},
			map[string]interface{}{"foo": "required"},
			true,
		},
		"required-float": {
			map[string]interface{}{"foo": 1.0},
			map[string]interface{}{"foo": "required"},
			true,
		},
		"required-string": {
			map[string]interface{}{"foo": "bar"},
			map[string]interface{}{"foo": "required"},
			true,
		},

		"required-nil": {
			map[string]interface{}{"foo": nil},
			map[string]interface{}{"foo": "required"},
			false,
		},
		"required-no": {
			map[string]interface{}{"bar": "foo"},
			map[string]interface{}{"foo": "required"},
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
