package dess

import "testing"

//
// external tests
//

func TestDecimalString(t *testing.T) {
	decimal := NewDecimal(4567, 2)
	if decimal.String() != "45.67" {
		t.Error(decimal.String())
	}

	decimal = NewDecimal(0, 45)
	if decimal.String() != "0.0" {
		t.Error(decimal.String())
	}

	decimal = NewDecimal(4, 4)
	if decimal.String() != "0.0004" {
		t.Error(decimal.String())
	}

	decimal = NewDecimal(-4, 4)
	if decimal.String() != "-0.0004" {
		t.Error(decimal.String())
	}
}

func TestDecimalAdd(t *testing.T) {
	d1 := NewDecimal(4567, 2)
	d2 := NewDecimal(2345, 2)
	decimal := d1.Add(d2)
	if decimal.String() != "69.12" {
		t.Error(decimal.String())
	}

	d1 = NewDecimal(4567, 3)
	d2 = NewDecimal(2345, 2)
	decimal = d1.Add(d2)
	if decimal.String() != "28.017" {
		t.Error(decimal.String())
	}
	decimal = d2.Add(d1)
	if decimal.String() != "28.017" {
		t.Error(decimal.String())
	}

	d1 = NewDecimal(-4567, 3)
	d2 = NewDecimal(2345, 2)
	decimal = d1.Add(d2)
	if decimal.String() != "18.883" {
		t.Error(decimal.String())
	}
	decimal = d2.Add(d1)
	if decimal.String() != "18.883" {
		t.Error(decimal.String())
	}
}

func TestDecimalEquals(t *testing.T) {
	d1 := NewDecimal(4567, 2)
	d2 := NewDecimal(2345, 2)
	if d1.Equals(d2) {
		t.Fail()
	}

	d3 := NewDecimal(4567, 2)
	if !d1.Equals(d3) {
		t.Fail()
	}
}

func TestDecimalSubtract(t *testing.T) {
	d1 := NewDecimal(4567, 2)
	d2 := NewDecimal(2345, 2)
	decimal := d1.Sub(d2)
	if decimal.String() != "22.22" {
		t.Error(decimal.String())
	}

	d1 = NewDecimal(4567, 3)
	d2 = NewDecimal(2345, 2)
	decimal = d1.Sub(d2)
	if decimal.String() != "-18.883" {
		t.Error(decimal.String())
	}
	decimal = d2.Sub(d1)
	if decimal.String() != "18.883" {
		t.Error(decimal.String())
	}
}

//
// internal tests
//

func TestPow(t *testing.T) {
	if pow(2, 3) != 8 {
		t.Fail()
	}

	if pow(-2, 3) != -8 {
		t.Fail()
	}

	if pow(2, 0) != 1 {
		t.Fail()
	}

	if pow(12, 12) != 8916100448256 {
		t.Fail()
	}

	if pow(-12, 12) != 8916100448256 {
		t.Fail()
	}
}
