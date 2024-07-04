package validator

import "testing"

func TestValidator_Valid_Default(t *testing.T) {
	v := New()

	if !v.Valid() {
		t.Errorf("got %v want %v", v.Valid(), false)
	}
}

func TestValidator_Valid_Invalid(t *testing.T) {
	v := New()

	v.AddError("name", "must not be empty")

	if v.Valid() {
		t.Errorf("got %v want %v", v.Valid(), true)
	}
}

func TestValidator_AddError(t *testing.T) {
	v := New()

	v.AddError("name", "must not be empty")

	if len(v.Errors) != 1 {
		t.Errorf("got %v want %v", len(v.Errors), 1)
	}
}

func TestValidator_NotEmpty(t *testing.T) {
	var cases = []struct {
		value    string
		expected bool
	}{{"test", true}, {"", false}, {" ", false}}

	for _, c := range cases {
		v := New()

		NotEmpty(v, "test", c.value)
		if v.Valid() != c.expected {
			t.Errorf("failed case %v: got %v want %v", c.value, v.Valid(), c.expected)
		}
	}
}

func TestValidator_GreaterThenEquals(t *testing.T) {
	min := 10

	var cases = []struct {
		value    int
		expected bool
	}{
		{1, false},
		{10, true},
		{11, true},
	}

	for _, c := range cases {
		v := New()

		GreaterThenEquals(v, "test", c.value, min)

		if v.Valid() != c.expected {
			t.Errorf("failed case %v: got %v want %v", c.value, v.Valid(), c.expected)
		}
	}
}

func TestValidator_LowerThenEquals(t *testing.T) {
	max := 10

	var cases = []struct {
		value    int
		expected bool
	}{
		{1, true},
		{10, true},
		{11, false},
	}

	for _, c := range cases {
		v := New()

		LowerThenEquals(v, "test", c.value, max)

		if v.Valid() != c.expected {
			t.Errorf("failed case %v: got %v want %v", c.value, v.Valid(), c.expected)
		}
	}
}
